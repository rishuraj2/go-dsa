package binarysearch

func minEatingSpeed(piles []int, h int) int {
	hourCounter := 0
	low := 1
	high := max(piles)

	for low <= high {
		mid := low + (high - low)/2
		hourCounter = 0

		for _, val := range piles {
			hourCounter += (val + mid - 1) / mid
		}

		if hourCounter > h {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func max(arr []int) int {
	arrLen := len(arr)
	max := arr[0]

	for i := 1; i < arrLen; i++ {

		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}