package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "blue"}
	fmt.Println(colors)
	//remove first
	colors = append(colors[1:])
	fmt.Println(colors)
	//remove last
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)
	//make func
	nums := make([]int, 5, 5)
	nums[0] = 8
	nums[1] = 5
	nums[2] = 23
	nums[3] = 2
	nums[4] = 123
	fmt.Println(nums)
	//print capacity of slice
	fmt.Println(cap(nums))
	nums = append(nums, 1233)
	//when we append above capacity capacity is doubled
	fmt.Println(cap(nums))
	//sorts the slice of ints
	sort.Ints(nums)
	fmt.Println(nums)
}
