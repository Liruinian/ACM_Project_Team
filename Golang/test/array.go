package main

import "fmt"

func arraySum(array []int) {
	sum := 0
	for _, i := range array {
		sum += i
	}
	fmt.Println(sum)
}

func arraySearchSum(array []int, wSum int) {
	var ans [][2]int
	for i, ai := range array {
		for j, aj := range array {
			if j > i {
				if ai+aj == wSum {
					arrTemp := [2]int{i, j}
					ans = append(ans, arrTemp)
				}
			}
		}
	}
	fmt.Println(ans)
}
