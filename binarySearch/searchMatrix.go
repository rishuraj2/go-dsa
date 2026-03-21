package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	low := 0
	high := len(matrix) - 1

	for low <= high {
		mid := low + (high - low)/2

		if target == matrix[mid][0] {
			return true
		}

		if target > matrix[mid][0] {
			low = mid+1
		} else {
			high = mid-1
		}
	}

	if high < 0 {
		return false
	}

	foundRow := high

	low = 0
	high = len(matrix[high]) - 1

	for low <= high {
		mid := low + (high - low)/2

		if matrix[foundRow][mid] == target {
			return true
		}

		if target > matrix[foundRow][mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}