package main

import (
	"fmt"
	"go-dsa/recursion"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(recursion.PowerSet(nums))
}
