package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUsagePatternsOfSpaceships() {
	spaceship, usages := readInputPatterns()
	fmt.Println(spaceship, usages)
}

func readInputPatterns() ([][]int, []string) {
	scanner := bufio.NewScanner(os.Stdin)

	var input [][]int
	var spaceshipNames []string

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		nameAndSequence := strings.Split(scanner.Text(), " ")
		spaceshipNames = append(spaceshipNames, nameAndSequence[0])
		sequenceString := nameAndSequence[1:]
		input = append(input, convertStringArrToIntArr(sequenceString))
	}

	return input, spaceshipNames
}
