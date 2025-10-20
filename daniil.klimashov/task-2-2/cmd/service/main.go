package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	value, err := x.(int)

	if !err {
		fmt.Println("Invalid value passed to heap")

		return
	} else {
		*h = append(*h, value)
	}
}
func (h *IntHeap) Pop() any {
	old := *h
	length := len(old)

	if length == 0 {
		fmt.Println("Invalid length of heap")

		return nil
	}

	x := old[length-1]
	*h = old[0 : length-1]

	return x
}

func main() {
	var amountOfDishes int

	if _, err := fmt.Scan(&amountOfDishes); err != nil {
		fmt.Println("Invalid amount of dishes")

		return
	}

	preferredDishes := &IntHeap{}
	heap.Init(preferredDishes)

	for range amountOfDishes {
		var dish int

		if _, err := fmt.Scan(&dish); err != nil {
			fmt.Println("Invalid type of dish")

			return
		}

		heap.Push(preferredDishes, dish)
	}

	var kPreferredDish int

	if _, err := fmt.Scan(&kPreferredDish); err != nil {
		fmt.Println("Invalid value of k preferred dish")

		return
	}

	if kPreferredDish > amountOfDishes {
		fmt.Println("Preferred dish must be less than amount of dishes")

		return
	}

	for range kPreferredDish - 1 {
		heap.Pop(preferredDishes)
	}

	fmt.Println(heap.Pop(preferredDishes))
}
