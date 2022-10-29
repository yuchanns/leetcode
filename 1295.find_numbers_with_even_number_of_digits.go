func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		n := 0
		for num > 0 {
			num /= 10
			n++
		}
		if n&1 == 0 {
			count++
		}
	}
	return count
}
