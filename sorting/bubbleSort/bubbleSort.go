package bubblesort

func BubbleSort(nums []int) {
	for i := len(nums)-1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if (nums[j] > nums[j+1]) {
				temp := nums[j];
				nums[j] = nums[j+1];
				nums[j+1] = temp;
			}
		}
	}
}
