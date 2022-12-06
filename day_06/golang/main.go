package main

import (
	"bufio"
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
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		solution1 := findMarker(line, 4)
		solution2 := findMarker(line, 14)
		fmt.Println(solution1)
		fmt.Println(solution2)
	}

}
