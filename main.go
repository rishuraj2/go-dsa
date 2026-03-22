package main

import (
	"fmt"
	"go-dsa/recursion"
)

func main() {
	nums := []int{4, 2, 10, 5, 1, 3}
	fmt.Println(recursion.CountSubsequenceWithTargetSum(nums, 5))
}
