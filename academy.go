package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateTeachersScores() {
	scores, teacherNames := readInput()

	fmt.Printf("printing input: %v, %v", scores, teacherNames)
}

func convertStringArrToIntArr(stringArray []string) []int {
	var intArray []int
	for _, str := range stringArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting %s to integer: %v\n", str, err)
			return nil
		}
		intArray = append(intArray, num)
	}
	return intArray
}

func readInput() ([][]int, []string) {
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
