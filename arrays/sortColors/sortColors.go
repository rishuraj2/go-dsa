package sortcolors

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, low, high int) {
	if (low < high) {
		pivot := partition(nums, low, high)

		quickSort(nums, low, pivot - 1)
		quickSort(nums, pivot + 1, high)
	}
}

func partition(nums []int, low, high int) (i int) {
	pivot := nums[high]
	i = low - 1

	for j := low; j < high; j++ {
		if (nums[j] <= pivot) {
			i++
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
		}
	}

	i++
	nums[high] = nums[i]
	nums[i] = pivot

	return
}