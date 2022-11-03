func heightChecker(heights []int) int {
	count := 0
	sortedHeights := qs(heights)
	for i := 0; i < len(heights); i++ {
		if heights[i] != sortedHeights[i] {
			count++
		}
	}
	return count
}

func qs(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := (len(arr) - 1) >> 1
	left, right := make([]int, 0, len(arr)), make([]int, 0, len(arr))
	for i := range arr {
		if i == mid {
			continue
		}
		if arr[i] <= arr[mid] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(qs(left), arr[mid]), qs(right)...)
}
