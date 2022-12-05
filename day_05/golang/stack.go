package main

import (
	"container/list"
	"fmt"
)

type customStack struct {
	stack *list.List
}

func (c *customStack) PushBottom(value string) {
	c.stack.PushFront(value)
}

func (c *customStack) PushTop(value string) {
	c.stack.PushBack(value)
}

func (c *customStack) Pop() (string, error) {
	if c.stack.Len() == 0 {
		return "", fmt.Errorf("stack is empty")
	}
	e := c.stack.Back()
	c.stack.Remove(e)
	return e.Value.(string), nil
}

func (c *customStack) PopN(n int) ([]string, error) {
	if c.stack.Len() < n {
		return nil, fmt.Errorf("stack is empty")
	}
	var result []string
	for i := 0; i < n; i++ {
		e := c.stack.Back()
		c.stack.Remove(e)
		result = append(result, e.Value.(string))
	}
	return result, nil
}

func (c *customStack) PushBottomMany(values []string) {
	for i := len(values) - 1; i >= 0; i-- {
		c.PushBottom(values[i])
	}
}

func (c *customStack) PushTopMany(values []string) {
	for i := len(values) - 1; i >= 0; i-- {
		c.PushTop(values[i])
	}
}

func (c *customStack) Top() (string, error) {
	if c.stack.Len() > 0 {
		ele := c.stack.Back()
		return ele.Value.(string), nil
	}
	return "", fmt.Errorf("Top Error: Stack is empty")
}

func (c *customStack) Bottom() (string, error) {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		return ele.Value.(string), nil
	}
	return "", fmt.Errorf("Bottom Error: Stack is empty")
}

func (c *customStack) Size() int {
	return c.stack.Len()
}

func (c *customStack) Empty() bool {
	return c.stack.Len() == 0
}

func (c *customStack) Print() {
	for e := c.stack.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
