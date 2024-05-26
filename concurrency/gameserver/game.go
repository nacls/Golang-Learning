package gameserver

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Player struct {
	name    string
	channel chan string
	zone    *Map
}

type Map struct {
	id      int
	players map[string]*Player
	channel chan string
	sync.Mutex
}

type Game struct {
	maps    map[int]*Map
	players map[string]*Player
	sync.Mutex
}

func NewGame(mapIds []int) (*Game, error) {
	if len(mapIds) < 1 {
		return &Game{}, nil
	}

	g := &Game{
		maps:    make(map[int]*Map),
		players: make(map[string]*Player),
	}

	for _, id := range mapIds {
		if id <= 0 {
			return nil, errors.New("invalid map ID")
		}
		g.maps[id] = &Map{
			id:      id,
			players: make(map[string]*Player),
			channel: make(chan string),
		}
		go g.maps[id].FanOutMessages()
	}

	return g, nil
}

func (g *Game) ConnectPlayer(name string) error {
	g.Lock()
	defer g.Unlock()

	name = strings.ToLower(name)
	if _, exists := g.players[name]; exists {
		return errors.New("player already exists")
	}

	player := &Player{
		name:    name,
		channel: make(chan string, 100),
	}

	g.players[name] = player
	return nil
}

func (g *Game) SwitchPlayerMap(name string, mapId int) error {
	g.Lock()
	defer g.Unlock()

	name = strings.ToLower(name)
	player, playerExists := g.players[name]
	if !playerExists {
		return errors.New("player doesn't exist")
	}

	if _, mapExists := g.maps[mapId]; !mapExists {
		return errors.New("map doesn't exist")
	}

	if player.zone != nil && player.zone.id == mapId {
		return errors.New("player already in the specified map")
	}

	if player.zone != nil {
		oldMap := g.maps[player.zone.id]
		oldMap.Lock()
		delete(oldMap.players, name)
		oldMap.Unlock()
	}

	newMap := g.maps[mapId]
	newMap.Lock()
	newMap.players[name] = player
	newMap.Unlock()

	player.zone = newMap
	return nil
}

func (g *Game) GetPlayer(name string) (*Player, error) {
	g.Lock()
	defer g.Unlock()

	name = strings.ToLower(name)
	player, exists := g.players[name]
	if !exists {
		return nil, errors.New("player doesn't exist")
	}

	return player, nil
}

func (g *Game) GetMap(mapId int) (*Map, error) {
	g.Lock()
	defer g.Unlock()

	mapObj, exists := g.maps[mapId]
	if !exists {
		return nil, errors.New("map doesn't exist")
	}

	return mapObj, nil
}

func (m *Map) FanOutMessages() {
	for msg := range m.channel {
		m.Lock()
		for _, player := range m.players {
			if !strings.HasPrefix(msg, fmt.Sprintf("%s says:", strings.Title(player.name))) {
				select {
				case player.channel <- msg:
					fmt.Println(msg)
				default:
					// Drop the message if the player's channel is full
				}
			}
		}
		m.Unlock()
	}
}

func (p *Player) GetChannel() <-chan string {
	return p.channel
}

func (p *Player) SendMessage(msg string) error {
	msg = fmt.Sprintf("%s says: %s", strings.Title(p.name), msg)
	mapObj := p.zone
	if mapObj == nil {
		return errors.New("player is not in any map")
	}

	mapObj.channel <- msg
	return nil
}

func (p *Player) GetName() string {
	return p.name
}
