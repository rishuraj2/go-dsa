package recursion

func GenerateParenthesis(n int) []string {
	var res []string

	var helper func(int, int, int, string)

	helper = func(index int, open int, close int, mem string) {
		if index >= n*2 {
			res = append(res, mem)
			return
		}

		if open < n {
			helper(index+1, open+1, close, mem+"(")
		}

		if close < open {
			helper(index+1, open, close+1, mem+")")
		}

	}

	helper(0, 0, 0, "")

	return res
}
