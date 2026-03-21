package recursion

func countGoodNumbers(n int64) int {
	if n == 0 {
		return 0
	}

	result := int64(1)

	result = (result*powCustom(5, n/2))%1000000007
	result = (result*powCustom(4, n/2))%1000000007

	if n%2 != 0 {
		result = (result * 5)%1000000007
	}

	return int (result)
}

func powCustom(x, n int64) int64 {
	if n == 0 {
		return 1
	}

	half := powCustom(x, n/2)

	if n%2 == 0 {
		return (half * half)%1000000007
	} else {
		return (x * half * half)%1000000007
	}
}