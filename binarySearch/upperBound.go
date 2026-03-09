package binarysearch

func upperBound(nums []int, x int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high - low)/2

		if x < nums[mid] {
			high = mid-1
		} else {
			low = mid+1
		}
	}
	return low
}