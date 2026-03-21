package recursion

func myAtoi(s string) int {

	start := startString(s)
	sign, start := checkSign(s, start)

	total := 0

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] >'9' {
			break
		}

		total = (total*10) + (int(s[i]) - int('0'))

		if sign == 1 && total >= 2147483647 {
			return 2147483647
		}

		if sign == -1 && total >= 2147483648 {
			return 2147483648
		}
	}

	return sign*total
}

func startString(s string) int {
	start := 0

	for start < len(s) && s[start] == ' ' {
		start++
	}

	return start
}

func checkSign(s string, start int) (int, int) {
	if s[start] == '-' {
		return -1, start+1
	} 
	
	if s[start] == '+' {
		return 1, start+1
	}

	return 1, start
}