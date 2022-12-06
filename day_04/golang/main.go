package main

import (
	"aoc-2022/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SectionRange struct {
	start, end int
}

func isContained(range1 SectionRange, range2 SectionRange) bool {
	return range1.start >= range2.start && range1.end <= range2.end
}

// func overlapped(range1 SectionRange, range2 SectionRange) bool {
// -	| --	|  --	| ---	|
// -	|  -- | --	|	 -  |
// return isContained(range1, range2) || isContained(range2, range1) || (range1.start < range2.start && range1.end > range2.start) || (range2.start < range1.start && range2.end > range1.start)
// }

func areSeparate(range1 SectionRange, range2 SectionRange) bool {
	// ---	  |    --- |
	//    --- | ---		 |
	return (range1.end < range2.start) || (range2.end < range1.start)
}

func main() {

	lines := common.ReadFile(os.Args[1])

	// part1
	result := 0

	// part2
	result2 := 0

	for _, line := range lines {
		sections := strings.Split(line, ",")
		sectionRanges := make([]SectionRange, 2)
		for i, sec := range sections {
			section := strings.Split(sec, "-")
			sectionRanges[i].start, _ = strconv.Atoi(section[0])
			sectionRanges[i].end, _ = strconv.Atoi(section[1])
		}
		// fmt.Println(sectionRanges)
		if isContained(sectionRanges[0], sectionRanges[1]) || isContained(sectionRanges[1], sectionRanges[0]) {
			result += 1
		}

		if !areSeparate(sectionRanges[0], sectionRanges[1]) {
			result2 += 1
		}

	}

	fmt.Println(result)
	fmt.Println(result2)

}
