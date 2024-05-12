package main

import (
	"fmt"
	"strconv"
)

func ConvertToDigitalFormat(hour, minute, second int) string {
	return fmt.Sprintf("%s:%s:%s", addZero(hour), addZero(minute), addZero(second))
}

func addZero(num int) string {
	if num < 10 {
		return "0" + strconv.Itoa(num)
	}
	return strconv.Itoa(num)
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	hours := seconds / 3600
	seconds = seconds - hours*3600
	minutes := seconds / 60
	seconds = seconds - minutes*60
	return hours, minutes, seconds
}
