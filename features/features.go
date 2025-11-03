package main

import (
	"fmt"
	"math"
)

func stringToInt(str string) int {
	numOfDigits := len(str)
	integer := 0
	for i, char := range str {
		currDigit := int(char - '0')
		nthDigit := numOfDigits - i
		power := nthDigit - 1
		integer += currDigit * (int(math.Pow10(power)))
	}

	return integer
}

func intToString(integer int) string {
	var numOfDigits int
	for power := 1; ; power++ {
		multiplier := int(math.Pow10(power))
		if integer < multiplier {
			numOfDigits = power
			break
		}
	}

	digits := make([]int, numOfDigits)
	for power := (numOfDigits - 1); power >= 0; power-- {
		divider := int(math.Pow10(power))
		quotient := integer / divider
		index := (numOfDigits - 1) - power
		digits[index] = quotient
		integer -= quotient * divider
	}

	outputString := ""
	for _, digit := range digits {
		digitStr := string('0' + digit)
		outputString += digitStr
	}

	return outputString
}

func main() {
	strings := []string{"8", "45", "9901928390", "00000000001", "762838"}
	integers := []int{8, 45, 9901928390, 1, 762838}

	// Testing stringToInt()
	fmt.Println("####### Testing stringToInt() #######")
	for i, str := range strings {
		fmt.Printf("Expecting: %d, Actual: %d\n", integers[i], stringToInt(str))
	}
	fmt.Println("")

	// Testing intToString()
	fmt.Println("####### Testing intToString() #######")
	for i, integer := range integers {
		intToString(integer)
		expectedValue := strings[i]
		if integer == 1 {
			expectedValue = "1"
		}
		fmt.Printf("Expecting: %s, Actual: %s\n", expectedValue, intToString(integer))
	}
	fmt.Println("")
}
