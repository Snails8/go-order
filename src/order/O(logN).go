package order

import "fmt"

func logN() {
	array := []int{1, 3, 6, 9, 100, 101}
	target := 100

	p := solve(array, target)
	fmt.Println("target", p)
}

func solve(array []int, target int) int {
	start := 0
	end := len(array) -1

	var index int 

	for {
		if end < start {
			return -1
		}
		index = (start + end) / 2

		if array[index] == target {
			return index
		}

		if array[index] < target {
			start = index + 1 
		} else {
			end = index - 1
		}
	}
}