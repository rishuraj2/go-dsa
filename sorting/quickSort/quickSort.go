package quicksort

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums) - 1)
}

func quickSort(nums []int, low int, high int) {
	if (low < high) {
		mid := partition(nums, low, high);
		quickSort(nums, low, mid-1);
		quickSort(nums, mid+1, high);
	}
}

func partition(nums []int, low int, high int) int {
	pivot := nums[high];
	i := low-1;

	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++;
			temp := nums[i];
			nums[i] = nums[j]
			nums[j] = temp;
		}
	}

	i++;
	temp := nums[high];
	nums[high] = nums[i];
	nums[i] = temp;

	return i;

}