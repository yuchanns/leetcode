func merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0, m+n)
	i, j := 0, 0
	for p := 0; p < m+n; p++ {
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				result = append(result, nums1[i])
				i++
			} else {
				result = append(result, nums2[j])
				j++
			}
		} else if i < m {
			result = append(result, nums1[i])
			i++
		} else if j < n {
			result = append(result, nums2[j])
			j++
		}
	}
	copy(nums1, result[:m+n])
}
