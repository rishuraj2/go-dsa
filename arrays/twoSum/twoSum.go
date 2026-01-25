package twosum

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i, num := range nums {
		if j, ok := numsMap[target-num]; ok {
			return []int{i, j}
		}
		numsMap[num] = i
	}
	return nil
}