func moveZeroes(nums []int) {
	max := len(nums) - 1
	for i := 0; max != 0; max-- {
		if nums[i] == 0 {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[len(nums)-1] = 0
			continue
		}
		i++
	}
}
