package main

import "fmt"

// Unbounded KP problem
// Number of ways
func change(coins []int, N int, c int) int {
	if N == 0 {
		return 1
	} else if c < 0 {
		return 0
	}

	if N >= coins[c] {
		return change(coins, N-coins[c], c) + change(coins, N, c-1)
	} else {
		return change(coins, N, c-1)
	}
}

func dp_change(coins []int, N int) int {
	total_coins := len(coins)

	dp := make([][]int, total_coins+1)

	for i := 0; i <= total_coins; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 0; i <= N; i++ {
		dp[0][i] = 0
	}
	for i := 0; i <= total_coins; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= total_coins; i++ {
		for j := 1; j <= N; j++ {
			if j >= coins[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[total_coins][N]

}

func main() {
	coins := []int{1, 2, 4}
	N := 3

	fmt.Println("Number of ways (Recursion)", change(coins, N, len(coins)-1))
	fmt.Println("Number of ways (DP)", dp_change(coins, N))
}
