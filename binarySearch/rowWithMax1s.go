package binarysearch

func rowWithMax1s(mat [][]int) int {
	maxOnes := 0
	rowMaxOnes := -1

	for i := 0; i < len(mat); i++ {
		countOnes := 0
		
		low := 0
		high := len(mat[i])-1

		for low <= high {
			mid := low + (high-low)/2

			if mat[i][mid] == 1 {
				high = mid-1
			} else {
				low = mid+1
			}
		}

		countOnes = len(mat[i]) - low

		if countOnes > maxOnes {
			rowMaxOnes = i
			maxOnes = countOnes
		}
	}

	return rowMaxOnes
}