package main

import "testing"

func TestChange(t *testing.T) {
	cases := []struct {
		amount int
		coins  []int
		want   int
	}{
		{5, []int{1, 2, 5}, 4}, // 5, 2+2+1, 2+1+1+1, 1*5
		{3, []int{2}, 0},
		{10, []int{10}, 1},
		{0, []int{7}, 1}, // the empty combination
		{4, []int{1, 2}, 3},
		{500, []int{3, 5, 7, 8, 9, 10, 11}, 35502874},
	}
	for _, c := range cases {
		if got := Change(c.amount, c.coins); got != c.want {
			t.Errorf("Change(%d,%v) = %d, want %d", c.amount, c.coins, got, c.want)
		}
	}
}
