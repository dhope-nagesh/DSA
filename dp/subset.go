package main

import "fmt"

func subsets(list []int, i int, s []int) [][]int {
	if i == len(list) {
		return [][]int{s}
	}
	l1 := subsets(list, i+1, append(s, list[i]))
	l2 := subsets(list, i+1, s)
	l3 := append(l1, l2...)
	return l3
}

func main() {
	l := []int{1, 1, 2, 3, 4}
	fmt.Println(subsets(l, 0, []int{}))
	// fmt.Println("Nagesh Dhope")
}
