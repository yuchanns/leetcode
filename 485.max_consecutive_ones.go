func findMaxConsecutiveOnes(nums []int) int {
	cnt, max := 0, 0
	for _, num := range nums {
		if num == 1 {
			cnt += 1
		} else {
			if max < cnt {
				max = cnt
			}
			cnt = 0
		}
	}
	if max < cnt {
		max = cnt
	}
	return max
}
