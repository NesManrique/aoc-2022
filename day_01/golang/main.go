package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var sumByElf []int
	var elfSum int

	elfSum = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			sumByElf = append(sumByElf, elfSum)
			elfSum = 0
			continue
		}

		number, _ := strconv.Atoi(line)
		elfSum += number
	}

	readFile.Close()

	sumByElf = append(sumByElf, elfSum)

	sort.Ints(sumByElf)

	// First part
	fmt.Println(sumByElf[len(sumByElf)-1])

	// Second part
	sumTopThreeElf := sumByElf[len(sumByElf)-1] + sumByElf[len(sumByElf)-2] + sumByElf[len(sumByElf)-3]
	fmt.Println(sumTopThreeElf)

}
