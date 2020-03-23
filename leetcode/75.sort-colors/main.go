package main

func sortColors(nums []int) {
	l := 0
	r := len(nums) - 1
	curr := 0
	for curr <= r {
		if nums[curr] == 0 {
			nums[l], nums[curr] = nums[curr], nums[l]
			l++
			curr++
			continue
		}
		if nums[curr] == 2 {
			nums[r], nums[curr] = nums[curr], nums[r]
			r--
			continue
		}
		curr++
	}
	return
}
