package findmissingnumber

func missingNumber(nums []int) int {
	quickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums); i++ {
		if (nums[i] != i) {
			return i
		}
	}
	return len(nums)
}

func quickSort(nums []int, low, high int) {
	if (low < high) {
		pivot := partition(nums, low, high)
		quickSort(nums, low, pivot - 1)
		quickSort(nums, pivot + 1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1

	for j := low; j < high; j++ {
		if (nums[j] <= pivot) {
			i++;
			temp := nums[j]
			nums[j] = nums[i]
			nums[i] = temp
		}
	}

	i++

	nums[high] = nums[i]
	nums[i] = pivot

	return i
}
