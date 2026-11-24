package main

import "fmt"

// Change counts the COMBINATIONS of coins (unlimited supply) summing to
// amount — {1,2,2} and {2,1,2} are the same combination. Coins form the
// OUTER loop so each coin multiset is counted exactly once.
func Change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // one way to make zero: take nothing
	for _, c := range coins {
		for a := c; a <= amount; a++ {
			dp[a] += dp[a-c]
		}
	}
	return dp[amount]
}

func main() {
	fmt.Println(Change(5, []int{1, 2, 5}))
	fmt.Println(Change(3, []int{2}))
}
