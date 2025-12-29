package selectionsort

func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minNumIndex := i;

		for j := i + 1; j < len(nums); j++ {
			if (nums[j] < nums[minNumIndex]) {
				minNumIndex = j;
			}
		}
		temp := nums[i];
		nums[i] = nums[minNumIndex];
		nums[minNumIndex] = temp;
	}
}