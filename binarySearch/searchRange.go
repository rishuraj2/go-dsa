package binarysearch

func searchRange(nums []int, target int) []int {
	if !checkPresence(nums, target) {
		return []int{-1, -1}
	}
	
	start := -1
	end := -1

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high - low)/2
		if (target <= nums[mid]) {
			high = mid-1
		} else {
			low = mid+1
		}
	}

	start = low
	low = 0
	high = len(nums) - 1

	for low <= high {
		mid := low + (high - low)/2
		if (target < nums[mid]) {
			high = mid-1
		} else {
			low = mid+1
		}
	}
	end = high

	return []int{start, end}
}

func checkPresence(nums []int, target int) bool {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high - low)/2

		if target == nums[mid] {
			return true
		}

		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}