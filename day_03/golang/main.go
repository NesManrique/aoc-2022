package main

import (
	"aoc-2022/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadFile(os.Args[1])

	var comp1 string
	var comp2 string
	var val string
	var num int
	result := 0

	var elfGroup []string
	result2 := 0

	for _, line := range lines {

		// part1
		comp1 = line[0 : len(line)/2]
		comp2 = line[len(line)/2:]
		for i := 0; i < len(comp1); i++ {
			if strings.Contains(comp2, string(comp1[i])) {
				val = fmt.Sprintf("%d", comp1[i])
				num, _ = strconv.Atoi(val)
				if num <= 90 {
					result += num - 38
				} else {
					result += num - 96
				}
				break
			}
		}

		// part2
		elfGroup = append(elfGroup, line)
		if len(elfGroup) == 3 {
			for i := 0; i < len(elfGroup[0]); i++ {
				if strings.Contains(elfGroup[1], string(elfGroup[0][i])) && strings.Contains(elfGroup[2], string(elfGroup[0][i])) {
					val = fmt.Sprintf("%d", elfGroup[0][i])
					num, _ = strconv.Atoi(val)
					if num <= 90 {
						result2 += num - 38
					} else {
						result2 += num - 96
					}
					break
				}
			}
			elfGroup = []string{}
		}
	}

	fmt.Println(result)
	fmt.Println(result2)

}
