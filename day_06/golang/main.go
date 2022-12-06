package main

import (
	"aoc-2022/common"
	"fmt"
	"os"
	"strings"
)

func findMarker(line string, markerLength int) int {
	for i := markerLength - 1; i < len(line); i++ {
		slice := line[i-(markerLength-1) : i+1]
		// fmt.Println(slice)
		if isMarker(slice) {
			return i + 1
		}
	}
	return 0
}

func isMarker(slice string) bool {
	seen := ""
	for _, r := range slice {
		if strings.Contains(seen, string(r)) {
			return false
		}
		seen = strings.Join([]string{seen, string(r)}, "")
	}
	return true
}

func main() {
	lines := common.ReadFile(os.Args[1])

	for _, line := range lines {
		solution1 := findMarker(line, 4)
		solution2 := findMarker(line, 14)
		fmt.Println(solution1)
		fmt.Println(solution2)
	}

}
