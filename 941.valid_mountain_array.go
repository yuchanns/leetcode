func validMountainArray(arr []int) bool {
	l := len(arr)
	if l < 3 {
		return false
	}
	leftTop, rightTop := 0, l-1
	for i := range arr {
		if i == 0 {
			continue
		}
		if arr[i] > arr[i-1] && leftTop+1 == i {
			leftTop = i
			if leftTop == l-1 {
				return false
			}
		}
		if arr[l-i] < arr[l-i-1] && rightTop-1 == l-i-1 {
			rightTop = l - i - 1
			if rightTop == 0 {
				return false
			}
		}
		if leftTop == rightTop {
			return true
		}
	}
	return false
}
