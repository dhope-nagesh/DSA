package main

import "fmt"

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// Longest Common Subquence
//
func lcs(a string, b string, i int, j int) int {
	if i == len(a) || j == len(b) {
		return 0
	}
	if a[i] == b[j] {
		return 1 + lcs(a, b, i+1, j+1)
	} else {
		return max(lcs(a, b, i, j+1), lcs(a, b, i+1, j))
	}
}

func dp_lcs(a string, b string) int {
	n := len(a)
	m := len(b)
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
			if a[i-1] == b[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

func main() {
	a := "abecdgh"
	b := "bcedfhr"
	fmt.Println("LCS length (Recursion): ", lcs(a, b, 0, 0))
	fmt.Println("LCS length (DP): ", dp_lcs(a, b))
}
