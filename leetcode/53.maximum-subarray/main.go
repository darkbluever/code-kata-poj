package main

func maxSubArray(nums []int) int {
	var sum, max = nums[0], nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		if sum > max {
			max = sum
		}
	}
	return max
}
