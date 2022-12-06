package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"aoc-2022/common"
)

func main() {

	lines := common.ReadFile(os.Args[1])

	var sumByElf []int
	var elfSum int

	elfSum = 0

	for _, line := range lines {

		if line == "" {
			sumByElf = append(sumByElf, elfSum)
			elfSum = 0
			continue
		}

		number, _ := strconv.Atoi(line)
		elfSum += number
	}

	sumByElf = append(sumByElf, elfSum)

	sort.Ints(sumByElf)

	// First part
	fmt.Println(sumByElf[len(sumByElf)-1])

	// Second part
	sumTopThreeElf := sumByElf[len(sumByElf)-1] + sumByElf[len(sumByElf)-2] + sumByElf[len(sumByElf)-3]
	fmt.Println(sumTopThreeElf)

}
