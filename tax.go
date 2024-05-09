package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func tax() {
	income, err := readIncome()
	if err != nil {
		fmt.Println(0)
		return
	}

	tax := calculateTax(income)
	fmt.Println(tax)
}

func readIncome() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	// convert string to int
	income, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}
	return income, nil
}

func calculateTax(income int) int {
	const (
		taxRate1      = 5
		taxRate2      = 10
		taxRate3      = 15
		taxRate4      = 20
		taxStep1      = 10000
		taxStep2      = 40000
		taxStep3      = 50000
		taxLowerLimit = 1000
	)

	tax := 0

	if income < taxLowerLimit {
		return tax
	}

	// Calculate tax
	if income > taxStep1 {
		tax += (taxRate1 * taxStep1) / 100
		income -= taxStep1
	} else {
		tax += (taxRate1 * income) / 100
		return tax
	}

	if income > taxStep2 {
		tax += (taxRate2 * taxStep2) / 100
		income -= taxStep2
	} else {
		tax += (taxRate2 * income) / 100
		return tax
	}

	if income > taxStep3 {
		tax += (taxRate3 * taxStep3) / 100
		income -= taxStep3
	} else {
		tax += (taxRate3 * income) / 100
		return tax
	}

	tax += (taxRate4 * income) / 100
	return tax
}
