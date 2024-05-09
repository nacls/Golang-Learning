package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateTax() {
	// Read input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	// Tokenize input
	words := strings.Fields(line)

	// Check input size
	if len(words) != 1 {
		return
	}

	// Turn string to int
	income, err := strconv.Atoi(words[0])
	if err != nil {
		return
	}

	if income < 1000 || income > 100000000 {
		return
	}

	fmt.Println("Input in correct format")
}
