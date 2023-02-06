package main

import (
	"fmt"
)

func varf() {
	var (
		age    int
		flt    float32
		bol    bool
		introD string
	)
	flt = 1.000002
	age = 10
	bol = true
	introD = "Hello, GO!"
	amaze := 'a'
	fmt.Printf("%d %T\n", age, age)
	fmt.Printf("%v %T\n", bol, bol)
	fmt.Printf("%s %T\n", introD, introD)
	fmt.Printf("%c %T\n", amaze, amaze)
	fmt.Printf("%f %T\n", flt, flt)
}
