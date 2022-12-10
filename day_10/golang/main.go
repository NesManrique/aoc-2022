package main

import (
	"aoc-2022/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cpuCycle struct {
	initReg int
	addVal  int
	endReg  int
}

func main() {

	lines := common.ReadFile(os.Args[1])

	cycles := map[int]cpuCycle{}
	cycle := 1
	xregister := 1
	result1 := 0
	for _, line := range lines {

		command := strings.Fields(line)
		instruction := command[0]

		sprite := strings.Repeat(".", 40)
		if xregister-1 >= 0 && xregister-1 < 40 {
			sprite = common.ReplaceAtIndex(sprite, '#', xregister-1)
		}
		if xregister >= 0 && xregister < 40 {
			sprite = common.ReplaceAtIndex(sprite, '#', xregister)
		}
		if xregister+1 >= 0 && xregister+1 < 40 {
			sprite = common.ReplaceAtIndex(sprite, '#', xregister+1)
		}

		if (cycle-20)%40 == 0 && cycle <= 220 {
			result1 = result1 + (xregister * cycle)
		}

		if cycle%40 == 0 {
			// fmt.Printf("%3d \n", cycle)
			fmt.Printf("%s\n", string(sprite[(cycle-1)%40]))
		} else {
			// fmt.Printf("%3d ", cycle)
			fmt.Printf("%s", string(sprite[(cycle-1)%40]))
		}
		// fmt.Println(sprite)

		if instruction == "noop" {
			cycles[cycle] = cpuCycle{xregister, 0, xregister}
			cycle += 1
			continue
		}

		if instruction == "addx" {
			value, _ := strconv.Atoi(command[1])
			cycles[cycle] = cpuCycle{xregister, 0, xregister}
			// fmt.Println(cycle, cycles[cycle].initReg, cycles[cycle].addVal, cycles[cycle].endReg)
			cycle += 1

			if (cycle-20)%40 == 0 && cycle <= 220 {
				result1 = result1 + (xregister * cycle)
			}

			if cycle%40 == 0 {
				// fmt.Printf("%3d \n", cycle)
				fmt.Printf("%s\n", string(sprite[(cycle-1)%40]))
			} else {
				// fmt.Printf("%3d ", cycle)
				fmt.Printf("%s", string(sprite[(cycle-1)%40]))
			}
			// fmt.Println(sprite)

			cycles[cycle] = cpuCycle{xregister, value, xregister}
			xregister = xregister + value
			// fmt.Println(cycle, cycles2[cycle].initReg, cycles2[cycle].addVal, cycles2[cycle].endReg)
			cycle += 1
			continue
		}

	}

	fmt.Println(result1)
}
