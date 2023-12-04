package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readInput reads the file and returns a slice of strings.

// Args:
//		fileName (string) - The path of the file to be read.

// Returns:
//		([]string) - A slice of strings.

// Errors:
//		- If the specified file is not found, the function panics.
//		- If there is any error reading the specified file, the function panics.

func readInput(fileName string) []string {

	// Reading the input file.
	data, err := os.ReadFile(fileName)

	// Checking for errors.
	if err != nil {
		panic(err)
	}

	// Converting the slice of bytes to string, and then splitting it with the newline delimiter.
	stringSlice := strings.Split(string(data), "\n")

	return stringSlice
}

// calculateCalibrationValue takes in a string as input and returns a two-digit integer which is made up of the first and last numerical values in the input string.

// Args:
// 		inputString (string) - ex: 1threenine241gnrdfqn5

// Returns:
// 		(int) - ex: 15

// Errors:
// The function will panic if error occurs while converting the string to integer.

func calculateCalibrationValue(inputString string) int {

	var calibrationDigitsSliceAll []string
	for _, val := range inputString {

		// val will have the ascii values of each character.
		// checking if the ascii value falls in the range of 48-57. (denotes numbers 0-9)
		if val >= 48 && val <= 57 {
			calibrationDigitsSliceAll = append(calibrationDigitsSliceAll, string(val))
		}
	}

	calibrationDigitsSlice := []string{calibrationDigitsSliceAll[0], calibrationDigitsSliceAll[len(calibrationDigitsSliceAll)-1]}

	var calibrationDigitsString string

	for _, val := range calibrationDigitsSlice {
		calibrationDigitsString += val
	}

	calibrationDigitsInt, err := strconv.Atoi(calibrationDigitsString)

	if err != nil {
		panic(err)
	}

	return calibrationDigitsInt

}

func main() {
	inputSlice := readInput("input.txt")

	var sum int = 0

	for _, val := range inputSlice {
		sum += calculateCalibrationValue(val)
	}

	fmt.Println("The sum of calibration values is: ", sum)
}
