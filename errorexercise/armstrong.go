package errorexercise

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func isArmstrong(num int) bool {
	strNum := strconv.Itoa(num)
	n := len(strNum)
	sum := 0
	for _, digit := range strNum {
		d, _ := strconv.Atoi(string(digit))
		sum += int(math.Pow(float64(d), float64(n)))
	}
	return sum == num
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	return text, err
}

func IsArmstrong() {
	input, _ := readInput()

	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input, -1)
	sum := 0
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		sum += num
	}

	if isArmstrong(sum) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
