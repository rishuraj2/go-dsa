package recursion

func SubsetSums(nums []int) []int {
	var res []int

	var helper func(int, int)

	helper = func(index, sum int) {
		if index == len(nums) {
			res = append(res, sum)
			return
		}

		helper(index+1, sum)
		helper(index+1, sum+nums[index])
	}

	helper(0, 0)

	return res
}