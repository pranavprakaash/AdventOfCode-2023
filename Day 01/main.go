package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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

	// Creating a new slice with only the first and last digits.
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

type CalibrationDigitsWithIndex struct {
	value string
	index int
}

func calculateUpdatedCalibrationValue(inputString string) int {
	var calibrationDigitsAll []CalibrationDigitsWithIndex
	crossCheckList := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for _, val := range crossCheckList {
		if strings.Contains(inputString, val) {
			firstIndex := strings.Index(inputString, val)
			calibrationDigitsAll = append(calibrationDigitsAll, CalibrationDigitsWithIndex{value: val, index: firstIndex})

			lastIndex := strings.LastIndex(inputString, val)
			calibrationDigitsAll = append(calibrationDigitsAll, CalibrationDigitsWithIndex{value: val, index: lastIndex})

		}
	}

	sort.Slice(calibrationDigitsAll[:], func(i, j int) bool { return calibrationDigitsAll[i].index < calibrationDigitsAll[j].index })

	// Creating a new slice with only the first and last digits.
	calibrationDigits := []string{calibrationDigitsAll[0].value, calibrationDigitsAll[len(calibrationDigitsAll)-1].value}

	var calibrationDigitsFinal []string

	for _, val := range calibrationDigits {
		calibrationDigitsFinal = append(calibrationDigitsFinal, switchStrings(val))
	}

	var calibrationDigitsString string

	for _, val := range calibrationDigitsFinal {
		calibrationDigitsString += val
	}

	calibrationDigitsInt, err := strconv.Atoi(calibrationDigitsString)

	if err != nil {
		panic(err)
	}

	return calibrationDigitsInt
}

func switchStrings(inputDigit string) string {
	switch inputDigit {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return inputDigit
	}
}

// To measure time taken.
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")() // <-- The trailing () is the deferred call
	inputSlice := readInput("input.txt")

	var sum int = 0

	for _, val := range inputSlice {
		sum += calculateCalibrationValue(val)
	}

	fmt.Println("The sum of calibration values is: ", sum)

	sum = 0

	for _, val := range inputSlice {
		sum += calculateUpdatedCalibrationValue(val)
	}

	fmt.Println("The sum of calibration values using the updated function is: ", sum)
}
