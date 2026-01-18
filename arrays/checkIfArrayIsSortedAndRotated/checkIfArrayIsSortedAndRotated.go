package checkifarrayissortedandrotated

func Check(nums []int) bool {

	numsLen := len(nums)

	if (numsLen == 0) {
		return false
	}

	anomalyIndex := 0
	
	for i := 0; i < numsLen - 1; i++ {
		if nums[i] > nums[i+1] {
			anomalyIndex = i+1
		}
	}

	for i := anomalyIndex; i < (anomalyIndex + numsLen) - 1; i++ {
		if (nums[i % numsLen] > nums[(i+1) % numsLen]) {
			return false
		}
	}

	return true
}