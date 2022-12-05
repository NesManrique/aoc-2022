package main

import (
	"container/list"
	"fmt"
)

type stackList struct {
	stacks []*customStack
}

func NewStackList(size int) *stackList {
	stacks := make([]*customStack, size)
	for i := 0; i < size; i++ {
		stacks[i] = &customStack{stack: list.New()}
	}
	return &stackList{stacks: stacks}
}

func (s *stackList) PushBottom(value string, index int) {
	s.stacks[index].PushBottom(value)
}

func (s *stackList) PushTop(value string, index int) {
	s.stacks[index].PushTop(value)
}

func (s *stackList) StackSize(index int) int {
	return s.stacks[index].Size()
}

func (s *stackList) Size() int {
	return len(s.stacks)
}

func (s *stackList) StackPrint(index int) {
	s.stacks[index].Print()
}

func (s *stackList) Print() {
	for i := 0; i < s.Size(); i++ {
		fmt.Println("Stack: ", i)
		s.stacks[i].Print()
	}
}

func (s *stackList) PrintTop() {
	for i := 0; i < s.Size(); i++ {
		if !s.stacks[i].Empty() {
			first, err := s.stacks[i].Top()

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("First: ", first)
		}
	}
}

func (s *stackList) MoveToAnotherList(amount int, origin int, destination int) {
	for i := 0; i < amount; i++ {
		if s.stacks[origin].Empty() {
			break
		}
		val, _ := s.stacks[origin].Top()
		// fmt.Println("Move: ", val)
		s.stacks[origin].Pop()
		s.stacks[destination].PushTop(val)
	}
}

func (s *stackList) MoveToAnotherListMany(amount int, origin int, destination int) error {
	if s.stacks[origin].Empty() {
		return fmt.Errorf("origin stack is empty")
	}
	val, _ := s.stacks[origin].PopN(amount)
	// fmt.Println("Move: ", val)
	s.stacks[destination].PushTopMany(val)
	return nil
}
