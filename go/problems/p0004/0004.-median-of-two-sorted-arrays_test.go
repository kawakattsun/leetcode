package p0004

import "testing"

type Case struct {
	Input1   []int
	Input2   []int
	Expected float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []Case{
		Case{Input1: []int{1, 3}, Input2: []int{2}, Expected: 2.0},
		Case{Input1: []int{1, 2}, Input2: []int{3, 4}, Expected: 2.5},
	}
	for _, c := range cases {
		output := findMedianSortedArrays(c.Input1, c.Input2)
		if output != c.Expected {
			t.Fatalf("test failed. output: %f", output)
		}
	}
}
