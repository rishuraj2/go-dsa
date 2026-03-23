package recursion

func subsetsWithDup(nums []int) [][]int {
	quickSort(nums, 0, len(nums)-1)
	var res [][]int
	var helper func(int, []int)

	helper = func(index int, mem []int) {
		temp := make([]int, len(mem))
		copy(temp, mem)
		res = append(res, temp)

		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}

			mem = append(mem, nums[i])
			helper(i+1, mem)
			mem = mem[:len(mem)-1]
		}
	}
	
	helper(0, []int{})

	return res
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
		if nums[i] <= pivot {
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