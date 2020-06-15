package p0001

func twoSum(nums []int, target int) []int {
	indexies := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := indexies[target-n]; ok {
			return []int{j, i}
		}
		indexies[n] = i
	}
	return nil
}
