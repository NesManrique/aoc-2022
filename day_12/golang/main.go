package main

import (
	"aoc-2022/common"
	"fmt"
	"image"
	"os"
)

func main() {
	lines := common.ReadFile(os.Args[1])
	var start image.Point
	var end image.Point
	matrix := map[image.Point]rune{}

	for y, line := range lines {
		for x, char := range line {
			matrix[image.Point{x, y}] = rune(char)

			if char == 'S' {
				start.X = x
				start.Y = y
			}
			if char == 'E' {
				end.X = x
				end.Y = y
			}
		}
	}

	matrix[start], matrix[end] = 'a', 'z'
	distance := map[image.Point]int{end: 0}
	queue := []image.Point{end}
	var bestStart *image.Point

	for len(queue) > 0 {
		fmt.Println("qstart", queue)
		current := queue[0]
		queue = queue[1:]

		if matrix[current] == 'a' && bestStart == nil {
			bestStart = &current
		}

		for _, direction := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			next := current.Add(direction)
			_, visited := distance[next]
			_, valid := matrix[next]

			if valid && !visited && matrix[current] <= matrix[next]+1 {
				distance[next] = distance[current] + 1
				queue = append(queue, next)
			}
		}
	}

	fmt.Println(distance[start])
	fmt.Println(distance[*bestStart])

}
