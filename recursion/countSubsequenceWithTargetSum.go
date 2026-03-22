package recursion

func CountSubsequenceWithTargetSum(nums []int, k int) int {
	count := 0

	var helper func(int, []int, int)

	helper = func(index int, mem []int, sum int) {
		if index == len(nums) {
			if sum == k {
				count++
			}
			return
		}

		mem = append(mem, nums[index])
		helper(index+1, mem, sum+nums[index])
		mem = mem[:len(mem)-1]

		helper(index+1, mem, sum)
	}

	helper(0, []int{}, 0)
	
	return count
}