package searchmatrix

func searchMatrix(matrix [][]int, target int) bool {
	rowSearchStart := 0
	rowSearchEnd := len(matrix) - 1

	lengthRow := len(matrix[0])

	for rowSearchStart < rowSearchEnd {
		mid := rowSearchStart + (rowSearchEnd - rowSearchStart)/2

		if (target > matrix[mid][lengthRow - 1]) {
			rowSearchStart = mid + 1
		} else if (target < matrix[mid][0]) {
			rowSearchEnd = mid - 1
		} else {
			rowSearchEnd = mid
		}
	}

	colSearchStart := 0
	colSearchEnd := lengthRow - 1

	for colSearchStart <= colSearchEnd {
		mid := colSearchStart + (colSearchEnd - colSearchStart)/2
		if (matrix[rowSearchEnd][mid] == target) {
			return true
		}

		if (target < matrix[rowSearchEnd][mid]) {
			colSearchEnd = mid - 1
		} else {
			colSearchStart = mid + 1
		}
	}

	return false
	
}