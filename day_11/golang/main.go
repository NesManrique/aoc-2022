package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

// type Monkey

type Monkey struct {
	Items     []int
	Operation func(int) int
	Test      func(int) int
}

func main() {

	readFile, _ := os.ReadFile(os.Args[1])

	split := strings.Split(strings.TrimSpace(string(readFile)), "\n\n")

	monkeys := make([]Monkey, len(split))
	mcm := 1

	for _, s := range split {
		var items, operation string
		var id, opValue, testValue, valueIfTrue, valueIfFalse int
		replaced := strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(s)
		fmt.Sscanf(replaced,
			`Monkey %d:
  Starting items: %s
  Operation: new = old %s %d
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`,
			&id, &items, &operation, &opValue, &testValue, &valueIfTrue, &valueIfFalse,
		)

		json.Unmarshal([]byte("["+items+"]"), &monkeys[id].Items)

		monkeys[id].Operation = map[string]func(int) int{
			"+": func(old int) int { return old + opValue },
			"*": func(old int) int { return old * opValue },
			"^": func(old int) int { return old * old },
		}[operation]

		monkeys[id].Test = func(obj int) int {
			if obj%testValue == 0 {
				return valueIfTrue
			}
			return valueIfFalse
		}
		mcm *= testValue
	}

	monkeys2 := append([]Monkey{}, monkeys...)

	rounds := 20
	inspected := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for k, m := range monkeys {
			for _, item := range m.Items {
				item = m.Operation(item) / 3
				monkeys[m.Test(item)].Items = append(monkeys[m.Test(item)].Items, item)
				inspected[k]++
			}
			monkeys[k].Items = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	result1 := inspected[0] * inspected[1]
	fmt.Println(result1)

	// Find an operation that doesn't change the result of the modulo
	// f(x), f(w) % D == w % D

	// https://www.youtube.com/watch?v=Zybk9cfAR4k

	// https://en.wikipedia.org/wiki/Modulo_operation

	// Applying modulo twice (or more) yields the same result as applying it once because
	// of congruence. If a ≡ b (mod n) (a is congruent to b modulo n), then
	// a - b = kn for some integer k. In other words, a - b is a multiple of n.
	//
	// In practice, what I wanted to prove is
	// a ≡ (a mod n) (mod n)
	// a - (a - kn)  = kn
	// a - a + kn = kn
	// kn = kn

	rounds2 := 10000
	inspected2 := make([]int, len(monkeys2))
	for i := 0; i < rounds2; i++ {
		for k, m := range monkeys2 {
			for _, item := range m.Items {
				item = m.Operation(item) % mcm
				monkeys2[m.Test(item)].Items = append(monkeys2[m.Test(item)].Items, item)
				inspected2[k]++
			}
			monkeys2[k].Items = nil
		}
		// fmt.Println(monkeys)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected2)))
	result2 := inspected2[0] * inspected2[1]
	fmt.Println(result2)

}
