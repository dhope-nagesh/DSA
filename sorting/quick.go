package main

import "fmt"

func partition(a *[]int, s int, e int) int {
	p := s
	for s < e {
		for s < len(*a) && (*a)[s] <= (*a)[p] {
			s++
		}
		for (*a)[e] > (*a)[p] {
			e--
		}

		if s < e {
			(*a)[s], (*a)[e] = (*a)[e], (*a)[s]
		}
	}
	(*a)[p], (*a)[e] = (*a)[e], (*a)[p]
	return e
}

func quicksort(a *[]int, s int, e int) {
	if s < e {
		p := partition(a, s, e)
		quicksort(a, s, p-1)
		quicksort(a, p+1, e)
	}
}

func main() {
	a := []int{4, 8, 1, 3, 2, 5, 7, 6}
	fmt.Println(a)
	quicksort(&a, 0, len(a)-1)
	fmt.Println(a)
}
