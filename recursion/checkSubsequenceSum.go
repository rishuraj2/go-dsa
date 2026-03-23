package recursion

func checkSubsequenceSum(nums []int, k int) bool {
	res := false
	var helper func(int, int)

	helper = func(index, sum int) {
		if index == len(nums){
			if sum == k {
				res = true
			}
			return
		}

		if res == false {
			helper(index+1, sum+nums[index])
			helper(index+1, sum)
		}
	}

	helper(0, 0)

	return res
}