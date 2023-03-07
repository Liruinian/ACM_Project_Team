package main

import "sort"

func numc(nums []int) {
	var (
		i      int
		flag   int
		output int
	)
	sort.Ints(nums)
	flag = 0
	for i = 0; i < len(nums); i++ {
		if flag == 0 {
			output += nums[i]
			flag = 1
		} else {
			output -= nums[i]
			flag = 0
		}
	}
	if output < 0 {
		output = 0 - output
	}
	println(output)
}
