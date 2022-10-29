func checkIfExist(arr []int) bool {
	for i := range arr {
		if arr[i]&1 == 0 {
			for j := range arr {
				if j == i {
					continue
				}
				if arr[i]>>1 == arr[j] {
					return true
				}
			}
		}
	}

	return false
}
