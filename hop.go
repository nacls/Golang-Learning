package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hop() {
	// Read input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	// Tokenize input
	words := strings.Fields(line)

	// Check input size
	if len(words) != 2 {
		return
	}

	// Turn string to int64 (first p)
	p, err := strconv.Atoi(words[0])
	if err != nil {
		return
	}
	// Turn string to int64 (second q)
	q, err := strconv.Atoi(words[1])
	if err != nil {
		return
	}

	// Check input conditions
	if p > 100 || p < 2 || q > 100000 || q < 2 {
		return
	}

	// Modified Hope Logic
	for i := 1; i <= q; i++ {
		if i%p == 0 {
			var m = i / p
			for j := 1; j <= m; j++ {
				fmt.Print("Hop")
				if j != m {
					fmt.Print(" ")
				}
			}
		} else {
			fmt.Print(i)
		}

		if i != q {
			fmt.Println()
		}
	}
}
