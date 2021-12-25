package main

import "fmt"

func sum_subset(arr []int, sum int, i int) int {
	if sum == 0 {
		return 1
	} else if i == len(arr) || sum < 0 {
		return 0
	}
	return sum_subset(arr, sum-arr[i], i+1) + sum_subset(arr, sum, i+1)
}

func dp_sum_subset(arr []int, sum int) int {
	n := len(arr)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
	}

	for j := 0; j <= sum; j++ {
		dp[0][j] = 0
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-arr[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][sum]
}

func main() {
	l := []int{1, 5, 2, 3, 4}
	fmt.Println("Total number of ways (Recursion) ", sum_subset(l, 5, 0))
	fmt.Println("Total number of ways (Dp) ", dp_sum_subset(l, 5))
}
