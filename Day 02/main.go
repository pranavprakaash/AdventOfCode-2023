package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Outcome struct {
	r, g, b int
}

type Game struct {
	gameID   int
	outcomes []Outcome
}

// outcomes = [ 1 green, 8 red, 7 blue  8 blue, 8 red, 3 green  1 blue, 2 red  4 red, 7 blue  3 green, 3 blue, 3 red]
// colorGroup = [ 1 green, 8 red, 7 blue ]
// colors = [1 green]

func ReadAndProcess(inputFile string) []Game {

	data, err := os.ReadFile(inputFile)

	if err != nil {
		panic(err)
	}

	dataString := strings.Split(string(data), "\n")

	var gameSlice []Game

	for _, val := range dataString {

		// Skipping empty lines
		if len(val) > 0 {

			subVal := strings.Split(val, ":")

			gameID, err := strconv.Atoi(strings.Split(subVal[0], " ")[1])

			if err != nil {
				panic(err)
			}

			outcomes := strings.Split(subVal[1], ";")
			var outcomesSlice []Outcome

			for _, colorGroup := range outcomes {
				var red, green, blue int = 0, 0, 0

				colors := strings.Split(colorGroup, ",")

				for _, color := range colors {

					// fmt.Printf("%s - %s\t", strings.Split(color, " ")[1], strings.Split(color, " ")[2])

					if strings.Split(color, " ")[2] == "red" {
						count, err := strconv.Atoi(strings.Split(color, " ")[1])

						if err != nil {
							panic(err)
						}

						red += count
					} else if strings.Split(color, " ")[2] == "green" {
						count, err := strconv.Atoi(strings.Split(color, " ")[1])

						if err != nil {
							panic(err)
						}

						green += count
					} else if strings.Split(color, " ")[2] == "blue" {
						count, err := strconv.Atoi(strings.Split(color, " ")[1])

						if err != nil {
							panic(err)
						}

						blue += count
					}

				}

				outcomesSlice = append(outcomesSlice, Outcome{red, green, blue})
			}
			gameSlice = append(gameSlice, Game{gameID: gameID, outcomes: outcomesSlice})
		} else {
			continue
		}
	}
	return gameSlice
}

func Validate(outcome Game, valid Outcome) bool {
	for _, val := range outcome.outcomes {
		if val.r > valid.r {
			return false
		}
		if val.g > valid.g {
			return false
		}
		if val.b > valid.b {
			return false
		}
	}
	return true
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
	x := ReadAndProcess("input.txt")

	var sum int = 0

	for _, val := range x {
		fmt.Println(Validate(val, Outcome{r: 12, g: 13, b: 14}))
		fmt.Println(val)

		if Validate(val, Outcome{r: 12, g: 13, b: 14}) {
			sum += val.gameID
		}

	}

	fmt.Println(sum)

}
