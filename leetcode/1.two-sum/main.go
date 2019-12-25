package main

func twoSum(nums []int, target int) []int {
	// key:value-lookup  val:index
	m := make(map[int]int)
	for n := range nums {
		if idx, ok := m[nums[n]]; ok {
			return []int{idx, n}
		}
		m[target-nums[n]] = n
	}
	return nil
}
