package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUsagePatternsOfSpaceships() {
	usages, spaceship := readInputPatterns()

	for i := 0; i < len(spaceship); i++ {
		patternNum := calculatePatternNum(usages[i])
		fmt.Printf("%s %v\n", spaceship[i], patternNum)
	}
}

func calculatePatternNum(usageForSpaceship []int) int {
	if len(usageForSpaceship) < 3 {
		return 0
	}
	numOfSequences := 0

	for i := 0; i < len(usageForSpaceship); i++ {
		for j := len(usageForSpaceship); j >= 0; j-- {
			if i <= j {
				if doesPatternExist(usageForSpaceship[i:j]) {
					numOfSequences++
				}
			}
		}
	}

	return numOfSequences
}

func doesPatternExist(sequence []int) bool {
	if len(sequence) < 3 {
		return false
	}
	difference := -1
	for i := 0; i < len(sequence); i++ {
		if i == 1 {
			difference = sequence[i] - sequence[i-1]
		} else if i > 1 {
			if sequence[i]-sequence[i-1] != difference {
				return false
			}
		}
	}
	return true
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
