package removeduplicatesfromsortedarray

func RemoveDuplicatesFromSortedArray(nums []int) int {
	var uniqueRegister []int

	uniqueRegister = append(uniqueRegister, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			uniqueRegister = append(uniqueRegister, nums[i])
		}
	}

	copy(nums, uniqueRegister)

	return len(uniqueRegister)
}