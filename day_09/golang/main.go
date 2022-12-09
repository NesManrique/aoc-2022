package main

import (
	"aoc-2022/common"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

//        -1, 2  0, 2  1, 2
// -2, 1  -1, 1  0, 1  1, 1  2, 1
// -2, 0  -1, 0  0, 0  1, 0  2, 0
// -2,-1  -1,-1  0,-1  1,-1  2,-1
//        -1,-2  0,-2  1,-2

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func main() {
	lines := common.ReadFile(os.Args[1])

	directions := map[rune]image.Point{'U': {0, -1}, 'R': {1, 0}, 'D': {0, 1}, 'L': {-1, 0}}
	rope := make([]image.Point, 10)

	tailVisited1 := map[image.Point]bool{}
	tailVisited2 := map[image.Point]bool{}
	// fmt.Println(len(tailVisited))

	for _, line := range lines {
		movement := strings.Fields(line)
		direction := movement[0]
		distance, _ := strconv.Atoi(movement[1])
		// fmt.Println("direction", direction, "distance", distance)

		for i := 0; i < distance; i++ {
			// fmt.Println("head", head, "tail", tail)
			rope[0] = rope[0].Add(directions[rune(direction[0])])
			// fmt.Println("head", head, "tail", tail)

			for knot := 1; knot < len(rope); knot++ {
				// Check if distance between head and tail (last tail) is > 1
				if d := rope[knot-1].Sub(rope[knot]); abs(d.X) > 1 || abs(d.Y) > 1 {
					// fmt.Println("d: ", d)
					// Move last tail to the right direction to be 1 away from head (last tail)
					rope[knot] = rope[knot].Add(image.Point{sgn(d.X), sgn(d.Y)})
					// fmt.Println("head", head, "tail", tail)
				}
			}
			// fmt.Println("---")
			tailVisited1[rope[1]] = true
			tailVisited2[rope[len(rope)-1]] = true
		}
		// fmt.Println()
	}

	fmt.Println("Result 1: ", len(tailVisited1))
	fmt.Println("Result 2: ", len(tailVisited2))
}
