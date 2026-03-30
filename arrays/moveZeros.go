package arrays

func moveZeros(nums []int) {
	current := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			current++
			nums[current] = nums[i]
		}
	}
	current++
	for i := current; i < len(nums); i++ {
		nums[i] = 0
	}
}