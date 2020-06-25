package p0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}
	min, max, harf := 0, m, (m+n+1)/2
	for min <= max {
		i := (min + max) / 2
		j := harf - i
		if i < max && nums2[j-1] > nums1[i] {
			min = i + 1
		} else if i > min && nums1[i-1] > nums2[j] {
			max = i - 1
		} else {
			var (
				maxLeft  int
				maxRight int
			)
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					maxLeft = nums1[i-1]
				} else {
					maxLeft = nums2[j-1]
				}
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			if i == m {
				maxRight = nums2[j]
			} else if j == n {
				maxRight = nums1[i]
			} else {
				if nums1[i] < nums2[j] {
					maxRight = nums1[i]
				} else {
					maxRight = nums2[j]
				}
			}

			return float64((maxLeft + maxRight)) / 2.0
		}
	}
	return 0.0
}
