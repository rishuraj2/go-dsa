package recursion

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var helper func(int, []int, int)

	helper = func(index int, mem []int, sum int) {
		if sum >= target || index == len(candidates) {
			if sum == target {
				temp := make([]int, len(mem))
				copy(temp, mem)
				res = append(res, temp)
			}
			return
		}

		mem = append(mem, candidates[index])
		helper(index, mem, sum+candidates[index])
		mem = mem[:len(mem)-1]

		helper(index+1, mem, sum)
	}

	helper(0, []int{}, 0)

	return res
}