package main

import "fmt"

func multip() {
	var (
		i int
		j int
	)
	for i = 1; i < 10; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}
