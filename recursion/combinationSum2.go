package recursion

func combinationSum2(candidates []int, target int) [][]int {
	quickSort(candidates, 0, len(candidates)-1)

	var res [][]int

	var helper func(int, []int, int)

	helper = func(index int, mem []int, sum int) {

		if index == len(candidates) || sum >= target {
			if sum == target {
				temp := make([]int, len(mem))
				copy(temp, mem)
				res = append(res, temp)
			}
			return 
		}

		for i := index; i < len(candidates); i++ {

			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			
			mem = append(mem, candidates[i])
			helper(i+1, mem, sum+candidates[i])
			mem = mem[:len(mem)-1]
		}
	}

	helper(0, []int{}, 0)

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
		if nums[i] < pivot {
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