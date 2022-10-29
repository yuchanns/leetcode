func duplicateZeros(arr []int) {
	l := len(arr)
	i := 0
	for i < l {
		if arr[i] == 0 {
			arr = append(arr[:i+1], arr[i:l-1]...)
			i++
		}
		i++
	}
}
