package findmissingnumber

func missingNumber(nums []int) int {
	numsLength := len(nums)
	result := numsLength * (numsLength + 1) / 2

	for _, value := range nums {
		result -= value
	}

	return result
}
