package main

import "fmt"

// Minimum number of coins required for a change

func min(x int, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	if x < y {
		return x
	}
	return y
}

// Unbounded KP problem
// Number of ways
func change(coins []int, N int, c int, tc int) int {
	if N == 0 {
		return tc
	} else if c < 0 {
		return 0
	}

	if N >= coins[c] {
		return min(change(coins, N-coins[c], c, tc+1), change(coins, N, c-1, tc))
	} else {
		return change(coins, N, c-1, tc)
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
		dp[i][0] = 0
	}

	for i := 1; i <= total_coins; i++ {
		for j := 1; j <= N; j++ {
			if j >= coins[i-1] {
				dp[i][j] = min(dp[i-1][j], 1+dp[i][j-coins[i-1]])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[total_coins][N]

}

func main() {
	coins := []int{1, 2, 3, 7}
	N := 7

	fmt.Println("Minimum coins (Recursion)", change(coins, N, len(coins)-1, 0))
	fmt.Println("Minimum coins (DP)", dp_change(coins, N))
}
