package main

import (
	"Goland-Learning/concurrency/gameserver"
	"fmt"
)

func main() {

	fmt.Println("test 1:")
	test1()
	fmt.Println("------------")

	fmt.Println("test 2:")
	test2()
	fmt.Println("------------")

	fmt.Println("test 3:")
	test3()
	fmt.Println("------------")
}

func test1() {
	g, err := gameserver.NewGame([]int{})
	fmt.Println(err == nil)
	fmt.Println(g != nil)
}

func test2() {
	g, err := gameserver.NewGame([]int{1, 2, 3})
	fmt.Println(err == nil)

	err = g.ConnectPlayer("Cyn")
	fmt.Println(err == nil)
}

func test3() {
	g, err := gameserver.NewGame([]int{1, 2, 3})
	fmt.Println(err == nil)

	err = g.ConnectPlayer("Cyn")
	fmt.Println(err == nil)

	p, err := g.GetPlayer("CyN")
	fmt.Println(err == nil)
	fmt.Println(p != nil)
}
