package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("--- homework1 ---")
	varf()
	fmt.Println("--- homework2 ---")
	textf("hello沙河小王子")
	fmt.Println("--- homework3 ---")
	num := []int{221, 221, 5, 5, 123, 123, 754, 148, 148, 496, 496}
	numc(num)
	fmt.Println("--- homework4 ---")
	strc("asq3241[\\]")
	fmt.Println("--- homework5 ---")
	multip()
	fmt.Println("--- homework6 ---")
	array1 := []int{1, 3, 5, 7, 8}
	arraySum(array1)
	arraySearchSum(array1, 8)

	fmt.Printf("Press ENTER to exit...")
	b := make([]byte, 1)
	_, err := os.Stdin.Read(b)
	if err != nil {
		return
	}

}
