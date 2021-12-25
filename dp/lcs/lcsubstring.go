package main

import "fmt"

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// Longest Common Substring
// Reset to 0 when no match found for i, j
func lcs(a string, b string, i int, j int, count int) int {
	if i == len(a) || j == len(b) {
		return count
	}
	if a[i] == b[j] {
		return lcs(a, b, i+1, j+1, count+1)
	} else {
		return max(count, max(lcs(a, b, i, j+1, 0), lcs(a, b, i+1, j, 0)))
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
				dp[i][j] = 0
			}
		}
	}
	sub_length := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if sub_length < dp[i][j] {
				sub_length = dp[i][j]
			}
		}
	}
	return sub_length
}

func main() {
	a := "aabce"
	b := "habdabcaq"
	fmt.Println("LCSubstring length (Recursion): ", lcs(a, b, 0, 0, 0))
	fmt.Println("LCSubstring length (DP): ", dp_lcs(a, b))
}
