package recursion

import "math"

func myPow(x float64, n int) float64 {
	if n == 0 { 
		return 1
	}

	if n < 0 {
		x = 1/x
		n = int(math.Abs(float64(n)))
	}

	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	half := pow(x, n/2)

	if n%2 == 0 {
		return half * half
	} else {
		return x * half * half
	}
}