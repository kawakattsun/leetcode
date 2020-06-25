package p0004

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	expected := 2.0
	output := findMedianSortedArrays(nums1, nums2)
	if output != expected {
		t.Fatalf("test failed. output: %f", output)
	}
}
