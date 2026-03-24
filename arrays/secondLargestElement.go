package arrays

func secondLargestElement(nums []int) int {
	quickSort(nums, 0, len(nums)-1)

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			return nums[i]
		}
	}

	return -1
}

func quickSort(nums []int, low, high int) {
	if low < high {
		pivot := partition(nums, low, high)
		quickSort(nums, low, pivot-1)
		quickSort(nums, pivot+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]

	partition := low-1

	for i := low; i < high; i++ {
		if nums[i] > pivot {
			partition++
			temp := nums[partition]
			nums[partition] = nums[i]
			nums[i] = temp
		}
	}

	partition++
	nums[high] = nums[partition]
	nums[partition] = pivot

	return partition
}