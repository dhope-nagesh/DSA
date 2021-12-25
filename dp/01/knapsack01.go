package main

import "fmt"

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// max profit
func knapsack(wt []int, val []int, w int, n int) int {
	if n < 0 {
		return 0
	} else if wt[n] <= w {
		return max(val[n]+knapsack(wt, val, w-wt[n], n-1), knapsack(wt, val, w, n-1))
	} else {
		return knapsack(wt, val, w, n-1)
	}
}

func dp_knapsack(wt []int, val []int, w int) int {
	n := len(wt)
	m := w
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for j := 0; j <= m; j++ {
		dp[0][j] = 0
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if wt[i-1] <= j {
				dp[i][j] = max(dp[i-1][j], val[i-1]+dp[i-1][j-wt[i-1]])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][m]
}

func main() {
	wt := []int{1, 3, 4, 5}
	val := []int{1, 4, 5, 7}
	w := 10
	p := knapsack(wt, val, w, len(wt)-1)
	p2 := dp_knapsack(wt, val, w)
	fmt.Println("Max Profit (Recursion)", p)
	fmt.Println("Max Profit (DP)", p2)
}
