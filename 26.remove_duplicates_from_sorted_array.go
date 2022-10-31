func removeDuplicates(nums []int) int {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == nums[i+1] {
			nums = removeIndex(nums, i+1)
		}
	}
	return len(nums)
}

func removeIndex(nums []int, index int) []int {
	return append(nums[:index], nums[index+1:]...)
}
