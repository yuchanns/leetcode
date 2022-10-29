func removeElement(nums []int, val int) int {
	j := 0
	for i := range nums {
		i = i - j
		if nums[i] == val {
			nums = removeIndex(nums, i)
			j++
		}
	}

	return len(nums)
}

func removeIndex(nums []int, index int) []int {
	return append(nums[:index], nums[index+1:]...)
}
