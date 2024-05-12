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

	for i := 0; i < len(scores); i++ {
		average := calculateScore(scores[i])
		currentTeacher := teacherNames[i]
		switch {
		case average >= 80:
			fmt.Printf("%v Excellent\n", currentTeacher)
		case average >= 60:
			fmt.Printf("%v Very Good\n", currentTeacher)
		case average >= 40:
			fmt.Printf("%v Good\n", currentTeacher)
		default:
			fmt.Printf("%v Fair\n", currentTeacher)
		}
	}
}

func calculateScore(scores []int) int {
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	return sum / len(scores)
}

func convertStringArrToIntArr(stringArray []string) []int {
	var intArray []int
	for _, str := range stringArray {
		num, err := strconv.Atoi(str)
		if err != nil {
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
