package binarysearch

// needs optimization
func pow(m, n int) int {
	if n == 0 {
		return 1
	}

	half := pow(m, n/2)

	if n % 2 == 0 {
		return half * half
	} 

	return m * half * half
}

func nthRoot(n, m int) int {
	if m == 0 {
		return 0
	}

	if n == 1 {
		return m
	}


	low, high := 1, m

	for low <= high {
		mid := low + (high - low)/2

		midPowN := pow(mid, n)

		if midPowN == m {
			return mid
		}

		if midPowN > m {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return high
}