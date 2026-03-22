package recursion

func GenerateBinaryStrings(n int) []string {
	var result []string

	var helper func(int, string, int)

	helper = func(index int, mem string, prev int) {
		if index == n {
			result = append(result, mem)
			return
		}

		if prev == 0 {
			helper(index+1, mem+"0", 0)
			helper(index+1, mem+"1", 1)
		} else {
			helper(index+1, mem+"0", 0)
		}
	}

	helper(0, "", 0)
	return result
}