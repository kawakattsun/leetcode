package p0001

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	if output := twoSum(nums, target); !reflect.DeepEqual(output, expected) {
		t.Fatalf("test failed. output: %#v, expected: %#v", output, expected)
	}
}
