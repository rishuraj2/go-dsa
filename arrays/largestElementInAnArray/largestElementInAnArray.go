package largestelementinanarray

func LargestElementInAnArray(nums []int) (int) {

	if len(nums) == 0 {
		return 0
	}

	maxValue := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
	}

	return maxValue
}