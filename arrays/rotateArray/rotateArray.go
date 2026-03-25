package rotatearray

func rotate(nums []int, k int) {
	k = k%len(nums)
	if k == 0 || len(nums) == 1 {
		return
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, low, high int) {
	mid := low + (high-low)/2
	for i := low; i <= mid; i++ {
		temp := nums[i]
		nums[i] = nums[high-(i-low)]
		nums[high-(i-low)] = temp
	}
}