package main

import "testing"

func getFacePage(n int) int {
	return n/8 + 1
}
func TestFacePage(t *testing.T) {
	cases := []struct {
		given int
		want  int
	}{
		{0, 1},
		{7, 1},
		{8, 2},
		{15, 2},
		{72, 10},
		{79, 10},
	}
	for _, c := range cases {
		got := getFacePage(c.given)
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
