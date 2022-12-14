package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(filePath string) []string {
	// filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func ConvertRuneToInt(r rune) int {
	val := fmt.Sprintf("%d", r)
	num, _ := strconv.Atoi(val)
	return num
}
