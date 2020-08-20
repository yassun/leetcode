func moveZeroes(nums []int) {
	for nz, cur := 0, 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[nz], nums[cur] = nums[cur], nums[nz]
			nz++
		}
	}
}
