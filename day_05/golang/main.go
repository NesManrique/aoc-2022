package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var stacks *stackList
	var stacks2 *stackList

	firstLine := true
	stacknumber := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

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
	for fileScanner.Scan() {
		line := fileScanner.Text()
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
