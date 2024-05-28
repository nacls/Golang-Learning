package introduction

import (
	"math"
	"strconv"
)

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	root := math.Sqrt(float64(x))
	return root == float64(int(root))
}

func IsPalindrome(x int) bool {
	// Convert integer to string
	strNum := strconv.Itoa(Abs(x))

	// Check if the string is palindrome
	for i, j := 0, len(strNum)-1; i < j; i, j = i+1, j-1 {
		if strNum[i] != strNum[j] {
			return false
		}
	}
	return true
}

func Abs(num int) int {
	if num < 0 {
		num = num * -1
	}
	return num
}

func Cube(num int) int {
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	var output = make([]int, 0)
	for i := 0; i < len(input); i++ {
		if f(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func Map(input []int, m MapperFunc) []int {
	var output []int
	for i := 0; i < len(input); i++ {
		output = append(output, m(input[i]))
	}
	return output
}
