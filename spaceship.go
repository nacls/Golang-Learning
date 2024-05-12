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

	// number of teachers
	scanner.Scan()
	line := scanner.Text()

	// convert string to int
	numOfTeachers, _ := strconv.Atoi(line)

	var input [][]int
	var teacherNames []string

	i := 0
	for i < numOfTeachers*2 {
		scanner.Scan()
		line := scanner.Text()

		if i%2 == 0 { // is writing teacher's name
			teacherNames = append(teacherNames, line)
		} else { // is a line of student scores
			inputScoreLine := strings.Fields(line)
			input = append(input, convertStringArrToIntArr(inputScoreLine))
		}
		i++
	}

	return input, teacherNames
}
