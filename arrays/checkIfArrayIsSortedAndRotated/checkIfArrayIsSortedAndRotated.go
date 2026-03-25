package checkifarrayissortedandrotated

func Check(nums []int) bool {
	numDrops := 0
	numsLen := len(nums)

	for i := 0; i < numsLen; i++ {
		if nums[i] > nums[(i+1)%numsLen] {
			numDrops++
		}
	}

	return numDrops <= 1
}