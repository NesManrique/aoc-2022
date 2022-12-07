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

	// map of directory and size
	sizes := make(map[string]int)

	// path of directories
	var path []string

	for _, line := range lines {
		tokens := strings.Fields(line)
		// fmt.Println(tokens)
		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if tokens[2] == ".." {
					sizes[strings.Join(path[:len(path)-1], "/")] += sizes[strings.Join(path, "/")]
					path = path[:len(path)-1]
				} else {
					path = append(path, tokens[2])
					dir := strings.Join(path, "/")
					sizes[dir] = 0
				}
			}
			if tokens[1] == "ls" {
				continue
			}
		} else if tokens[0] == "dir" {
			continue
		} else if size, err := strconv.Atoi(tokens[0]); err == nil {
			sizes[strings.Join(path, "/")] += size
		} else {
			fmt.Println("ERROR", line)
		}
	}

	for dirs := len(path); dirs > 1; dirs = len(path) {
		sizes[strings.Join(path[:len(path)-1], "/")] += sizes[strings.Join(path, "/")]
		path = path[:len(path)-1]
	}

	var result1 int
	for _, size := range sizes {
		if size <= 100000 {
			result1 = result1 + size
		}
	}

	fmt.Println("Result 1:", result1)

	result2 := sizes["/"]

	// huehuehue
	// 70000000 - sizes["/"] + directory_size = 30000000
	// directory_size = - 70000000 + sizes["/"] + 30000000
	// directory_size >= 40000000 - sizes["/"]

	for _, size := range sizes {
		if size >= sizes["/"]-40000000 && size < result2 {
			result2 = size
		}
	}

	fmt.Println("Result 2:", result2)

}
