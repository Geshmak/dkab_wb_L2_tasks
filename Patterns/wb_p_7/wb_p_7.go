package main

import "fmt"

type strategySort interface {
	sort([]int)
}

type bubbleSort struct {
}

func (s *bubbleSort) sort(a []int) {
	fmt.Println("bubble")
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

type InsertionSort struct {
}

func (s *InsertionSort) sort(a []int) {
	fmt.Println("insert")
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

type context struct {
	strategy strategySort
}

func (c *context) Algorithm(a strategySort) {
	c.strategy = a
}

func (c *context) sort(s []int) {
	c.strategy.sort(s)
}

func main() {

	ctx := &context{}
	ctx.Algorithm(&InsertionSort{})
	ctx.sort([]int{2, 5, 8, 2, 4, 6, 7, 2, 3})

	ctx.Algorithm(&bubbleSort{})
	ctx.sort([]int{2, 5, 8, 2, 4, 6, 7, 2, 3})

}
