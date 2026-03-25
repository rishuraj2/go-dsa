package removeduplicatesfromsortedarray

func RemoveDuplicatesFromSortedArray(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	uniqueIndex := -1

	for i := 0; i < numsLen-1; i++ {
		if nums[i] != nums[i+1] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}

	uniqueIndex++
	nums[uniqueIndex] = nums[numsLen-1]

	return uniqueIndex+1
}