func removeElement(nums []int, val int) int {
	index := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[index] = nums[index], nums[i]
			index--
		}
	}
	return index + 1
}
