package main

import (
	"aoc-2022/common"
	"fmt"
	"os"
)

func main() {

	lines := common.ReadFile(os.Args[1])

	// Rock 1, A, X
	// Paper 2, B, Y
	// Scissors 3, C, Z
	// win 6
	// draw 3
	// lose 0

	outcomes := map[string]int{
		"A X": 4, // Rock vs Rock = Draw (3 + 1)
		"A Y": 8, // Rock vs Paper = Win (6 + 2)
		"A Z": 3, // Rock vs Scissors = Lose (0 + 3)
		"B X": 1, // Paper vs Rock = Lose (0 + 1)
		"B Y": 5, // Paper vs Paper = Draw (3 + 2)
		"B Z": 9, // Paper vs Scissors = Win (6 + 3)
		"C X": 7, // Scissors vs Rock = Win (6 + 1)
		"C Y": 2, // Scissors vs Paper = Lose (0 + 2)
		"C Z": 6, // Scissors vs Scissors = Draw (3 + 3)
	}

	// X lose 0
	// Y draw 3
	// X win 6
	// Rock 1
	// Paper 2
	// Scissors 3
	// win 6
	// draw 3
	// lose 0

	outcomes2 := map[string]int{
		"A X": 3, // Rock Lose, Scissors (0 + 3)
		"A Y": 4, // Rock Draw, Rock (3 + 1)
		"A Z": 8, // Rock Win, Paper (6 + 2)
		"B X": 1, // Paper Lose, Rock (0 + 1)
		"B Y": 5, // Paper Draw, Paper (3 + 2)
		"B Z": 9, // Paper Win, Scissors (6 + 3)
		"C X": 2, // Scissors Lose, Paper (0 + 2)
		"C Y": 6, // Scissors Draw, Scissors (3 + 3)
		"C Z": 7, // Scissors Win, Rock (6 + 1)
	}

	result := 0
	result2 := 0

	for _, line := range lines {
		result += outcomes[line]
		result2 += outcomes2[line]
	}

	// First part
	fmt.Println(result)

	// Second part
	fmt.Println(result2)

}
