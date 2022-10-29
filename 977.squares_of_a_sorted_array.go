func sortedSquares(nums []int) []int {
	nl := len(nums)
	sq := make([]int, nl)
	i, j := 0, nl-1
	for p := nl - 1; p >= 0; p-- {
		ni, nj := nums[i], nums[j]
		if abs(ni) > abs(nj) {
			sq[p] = ni * ni
			i++
		} else {
			sq[p] = nj * nj
			j--
		}
	}
	return sq
}

func abs(i int) int {
	mask := i >> 63
	return (mask + i) ^ mask
}
