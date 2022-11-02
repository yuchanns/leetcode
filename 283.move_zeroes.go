func moveZeroes(nums []int) {
	w, cnt, l := -1, 0, len(nums)
	for r := 0; r < l; r++ {
		if nums[r] == 0 {
			cnt++
			if w == -1 {
				w = r
			}
		} else {
			if w != -1 {
				nums[w] = nums[r]
				w++
			}
		}
	}
	for i := 0; i < cnt; i++ {
		nums[l-1-i] = 0
	}
}
