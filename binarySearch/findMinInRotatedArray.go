package binarysearch

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	minVal := nums[0]

	for low <= high {
		if nums[low] < nums[high] {
			minVal = min(minVal, nums[low])
		}

		mid := low + (high - low)/2
		minVal = min(minVal, nums[mid])
		if nums[low] <= nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return minVal
}

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}