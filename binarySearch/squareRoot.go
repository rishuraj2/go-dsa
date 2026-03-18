package binarysearch

func squareRoot(n int) int {
	if n == 0 {
		return 0
	}

	low := 1
	high := n

	for low <= high {
		mid := low + (high - low)/2

		if mid == n/mid {
			return mid
		}

		if mid > n/mid {
			high = mid-1
		} else {
			low = mid+1
		}
	}

	return high
}