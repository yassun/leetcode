func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	rob, nonRob := 0, 0
	for _, v := range nums {
		rob, nonRob = nonRob+v, getMax(rob, nonRob)
	}

	return getMax(rob, nonRob)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
