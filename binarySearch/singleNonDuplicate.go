package binarysearch

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high - low)/2

		if nums[mid] != nums[mid+1] && nums[mid] != nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] == nums[mid+1] {
			low = mid + 1
		} else {
			high = mid - 2
		}
	}

	return -1
}