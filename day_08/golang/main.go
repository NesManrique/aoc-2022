package main

import (
	"aoc-2022/common"
	"fmt"
	"os"
	"strconv"
)

// A tree is visible if all of the other trees between it and an edge of the grid are shorter than it.
// Score = how many trees are visible from that tree, i.e. how many trees are shorter or the same height as it.
// Trees with height 0 are also trees and count for the score.

func isVisibleNorthAndScore(x, y int, numbers [][]int) (bool, int) {
	score := 0

	if x != 0 {

		tree := numbers[x][y]

		for i := x - 1; i >= 0; i-- {
			ntree := numbers[i][y]
			if ntree >= tree {
				score++
				return false, score
			} else {
				score++
			}
		}
	}

	return true, score

}

func isVisibleSouthAndScore(x, y int, numbers [][]int) (bool, int) {
	score := 0

	if x != len(numbers)-1 {

		tree := numbers[x][y]

		for i := x + 1; i < len(numbers); i++ {
			stree := numbers[i][y]
			if stree >= tree {
				score++
				return false, score
			} else {
				score++
			}
		}
	}

	return true, score

}

func isVisibleEastAndScore(x, y int, numbers [][]int) (bool, int) {
	score := 0

	if y != len(numbers[0])-1 {

		tree := numbers[x][y]

		for i := y + 1; i < len(numbers[0]); i++ {
			etree := numbers[x][i]
			if etree >= tree {
				score++
				return false, score
			} else {
				score++
			}
		}
	}

	return true, score

}

func isVisibleWestAndScore(x, y int, numbers [][]int) (bool, int) {
	score := 0

	if y != 0 {

		tree := numbers[x][y]

		for i := y - 1; i >= 0; i-- {
			wtree := numbers[x][i]
			if wtree >= tree {
				score++
				return false, score
			} else {
				score++
			}
		}
	}

	return true, score

}

func isVisibleAndScore(x, y int, numbers [][]int) (bool, int) {
	north, nscore := isVisibleNorthAndScore(x, y, numbers)
	south, sscore := isVisibleSouthAndScore(x, y, numbers)
	west, wscore := isVisibleWestAndScore(x, y, numbers)
	east, escore := isVisibleEastAndScore(x, y, numbers)

	return north || south || west || east, nscore * sscore * wscore * escore
}

func main() {
	lines := common.ReadFile(os.Args[1])
	numbers := make([][]int, len(lines))

	for i, line := range lines {
		numbers[i] = make([]int, len(line))
		for j, char := range line {
			numbers[i][j], _ = strconv.Atoi(string(char))
		}
	}

	result1 := 0
	result2 := -1
	// scores := make([][]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		// scores[i] = make([]int, len(numbers[i]))
		for j := 0; j < len(numbers); j++ {
			visible, score := isVisibleAndScore(i, j, numbers)
			if visible {
				result1 += 1
			}

			if score > result2 {
				result2 = score
			}
			// scores[i][j] = score
			// fmt.Print(scores[i][j])
		}
		// fmt.Println()
	}

	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)

}
