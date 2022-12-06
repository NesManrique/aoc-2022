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

	var stacks *stackList
	var stacks2 *stackList

	firstLine := true
	stacknumber := 0
	var index int
	var line string
	for index, line = range lines {

		if string(line[1]) == "1" {
			break
		}

		if stacknumber == 0 {
			for i := 1; i < len(line); i = i + 4 {
				stacknumber++
			}
			// fmt.Println("number of stacks: ", stacknumber)
		}

		if firstLine {
			stacks = NewStackList(stacknumber)
			stacks2 = NewStackList(stacknumber)
			firstLine = false
		}

		for i := 1; i < len(line); i = i + 4 {
			token := string(line[i])

			// fmt.Println("token: ", token)
			// fmt.Println("Line elem: ", i)

			if token != " " {
				// fmt.Println("Pushing: ", token)
				// fmt.Println("Pushing to stack: ", i/4)
				stacks.PushBottom(token, i/4)
				stacks2.PushBottom(token, i/4)
				// fmt.Println("Len: ", stacks.StackSize(i/4))
			}
			// fmt.Println("")
		}

	}

	// stacks.Print()

	// moves
	for _, line := range lines[index+1:] {
		if string(line) == "" {
			continue
		}
		tokens := strings.Fields(line)
		// fmt.Println(tokens)

		amount, _ := strconv.Atoi(string(tokens[1]))
		orig, _ := strconv.Atoi(string(tokens[3]))
		dest, _ := strconv.Atoi(string(tokens[5]))
		stacks.MoveToAnotherList(amount, orig-1, dest-1)
		err := stacks2.MoveToAnotherListMany(amount, orig-1, dest-1)
		if err != nil {
			continue
		}
	}

	stacks.PrintTop()
	fmt.Println("")
	stacks2.PrintTop()

}
