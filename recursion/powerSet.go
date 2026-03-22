package recursion

func PowerSet(nums []int) [][]int {
	var res [][]int

	var setGenerator func(int, []int)

	setGenerator = func(index int, mem []int) {

		if index == len(nums) {
			temp := make([]int, len(mem))
			copy(temp, mem)
			res = append(res, temp)
			return
		}

		mem = append(mem, nums[index])
		setGenerator(index+1, mem)
		mem = mem[:len(mem)-1]

		setGenerator(index+1, mem)
	}

	setGenerator(0, []int{})

	return res
}