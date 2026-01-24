package maxconsicutiveones

func findMacConsicutiveOnes(nums []int) int {
	maxOnes := 0
	tempMaxOnes := 0

	for i := 0; i < len(nums); i++ {
		if (nums[i] == 1) {
			tempMaxOnes++
		} else {
			if (tempMaxOnes > maxOnes) {
				maxOnes = tempMaxOnes
			}
			tempMaxOnes = 0
		}
	}

	if (tempMaxOnes > maxOnes) {
		maxOnes = tempMaxOnes
	}

	return maxOnes
}