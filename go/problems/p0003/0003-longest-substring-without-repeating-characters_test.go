package p0003

import "testing"

type Case struct {
	Str      string
	Expected int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []Case{
		Case{Str: "abcabcbb", Expected: 3},
		Case{Str: "bbbbb", Expected: 1},
		Case{Str: "pwwkew", Expected: 3},
		Case{Str: "dvdf", Expected: 3},
	}
	for _, c := range cases {
		if output := lengthOfLongestSubstring(c.Str); output != c.Expected {
			t.Fatalf("test failed. output: %d, expected: %d", output, c.Expected)
		}

	}
}
