package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printCountryCodes() {
	scanner := bufio.NewScanner(os.Stdin)
	var countryCodes = make(map[string]string)

	scanner.Scan()
	numOfCountries, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < numOfCountries; i++ {
		scanner.Scan()
		countryAndCode := strings.Split(scanner.Text(), " ")
		countryCodes[countryAndCode[1]] = countryAndCode[0]
	}

	scanner.Scan()
	numOfNumbers, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < numOfNumbers; i++ {
		scanner.Scan()
		number := scanner.Text()
		found := false
		for key, value := range countryCodes {
			if strings.HasPrefix(number, key) {
				fmt.Println(value)
				found = true
			}
		}
		if !found {
			fmt.Println("Invalid Number")
		}
	}
}
