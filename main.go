package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	numDice := 1
	seed := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(seed)

	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error assigning arg to number of dice: %v\n", err)
		} else {
			numDice = i
		}
	}

	var dice [][]string

	for i := 0; i < numDice; i++ {
		die := getSide(randomNumber.Intn(5) + 1)
		dice = append(dice, die)
	}

	// die := []string{
	// 	" ------   ",
	// 	"|      |  ",
	// 	"|      |  ",
	// 	"|      |  ",
	// 	" ------   ",
	// }

	for i := 0; i < 5; i++ {
		for j, n := 0, numDice; j < n; j++ {

			fmt.Print(dice[j][i])

		}
		fmt.Println()
	}
}

func getSide(side int) []string {

	switch side {
	case 1:
		return []string{
			"┌───────┐  ",
			"│       │  ",
			"│   ∙   │  ",
			"│       │  ",
			"└───────┘  ",
		}

	case 2:
		return []string{
			"┌───────┐  ",
			"│       │  ",
			"│ ∙   ∙ │  ",
			"│       │  ",
			"└───────┘  ",
		}

	case 3:
		return []string{
			"┌───────┐  ",
			"│ ∙     │  ",
			"│   ∙   │  ",
			"│     ∙ │  ",
			"└───────┘  ",
		}

	case 4:
		return []string{
			"┌───────┐  ",
			"│ ∙   ∙ │  ",
			"│       │  ",
			"│ ∙   ∙ │  ",
			"└───────┘  ",
		}

	case 5:
		return []string{
			"┌───────┐  ",
			"│ ∙   ∙ │  ",
			"│   ∙   │  ",
			"│ ∙   ∙ │  ",
			"└───────┘  ",
		}

	case 6:
		return []string{
			"┌───────┐  ",
			"│ ∙   ∙ │  ",
			"│ ∙   ∙ │  ",
			"│ ∙   ∙ │  ",
			"└───────┘  ",
		}
	}
	return nil
}
