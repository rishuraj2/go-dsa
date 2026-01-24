package rotatearray

func rotate(nums []int, k int) {
	result := []int{}
	numsLength := len(nums)
	k = k%numsLength
	start := numsLength - k
	end := start + numsLength

	for i := start; i < end; i++ {
		result = append(result, nums[i%numsLength])
	}

	copy(nums, result)
}