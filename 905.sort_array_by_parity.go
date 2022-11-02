func sortArrayByParity(nums []int) []int {
	l := len(nums) - 1
	ne := l
	for i := l; i >= 0; i-- {
		if nums[i]&1 != 0 {
			nums[ne], nums[i] = nums[i], nums[ne]
			ne--
		}
	}
	return nums
}
