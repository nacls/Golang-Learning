package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func runLibrary() {
	var books = make(map[int]string)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	instructions, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < instructions; i++ {
		scanner.Scan()
		instructions := strings.Split(scanner.Text(), " ")
		if instructions[0] == "ADD" {
			isbn, _ := strconv.Atoi(instructions[1])
			books[isbn] = instructions[2]
		} else if instructions[0] == "REMOVE" {
			isbn, _ := strconv.Atoi(instructions[1])
			delete(books, isbn)
		}
	}

	for _, kv := range sortMapByValueThenKey(books) {
		fmt.Println(kv)
	}
}

func sortMapByValueThenKey(m map[int]string) []int {
	// Create a slice to store key-value pairs
	var kvSlice []struct {
		Key   int
		Value string
	}

	// Populate the slice with key-value pairs from the map
	for k, v := range m {
		kvSlice = append(kvSlice, struct {
			Key   int
			Value string
		}{k, v})
	}

	// Sort the slice based on values and keys
	sort.Slice(kvSlice, func(i, j int) bool {
		if kvSlice[i].Value != kvSlice[j].Value {
			return kvSlice[i].Value < kvSlice[j].Value
		}
		return kvSlice[i].Key < kvSlice[j].Key
	})

	// Extract keys from the sorted slice
	var sortedKeys []int
	for _, kv := range kvSlice {
		sortedKeys = append(sortedKeys, kv.Key)
	}

	return sortedKeys
}
