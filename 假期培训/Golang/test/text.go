package main

import (
	"fmt"
	"unicode"
)

func textf(s string) {
	count := 0
	for _, i := range s {
		if unicode.Is(unicode.Han, i) {
			count++
		}
	}
	fmt.Printf("%s 中汉字的个数为 %d\n", s, count)

}
