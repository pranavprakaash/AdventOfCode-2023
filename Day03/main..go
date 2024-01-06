package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	calculate_sum("input.txt")
}

func calculate_sum(input_file string) {
	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println(err)
	}

	data_in_string := string(data)
	x := strings.Split(data_in_string, "\n")

	for _, val := range x {
		for _, char := range val {
			if fmt.Sprintf("%c", char) != "." {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
	}
}
