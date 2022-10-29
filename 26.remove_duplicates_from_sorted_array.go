func removeDuplicates(nums []int) int {
	j := 0
	for i := range nums {
		if i == 0 {
			continue
		}
		i -= j
		if nums[i] == nums[i-1] {
			nums = removeIndex(nums, i-1)
			j++
		}
	}
	return len(nums)
}

func removeIndex(nums []int, index int) []int {
	return append(nums[:index], nums[index+1:]...)
}
