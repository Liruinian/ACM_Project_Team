package main

import "fmt"

func strc(str string) {
	var (
		numCount    int
		letterCount int
		otherCount  int
	)
	for _, s := range str {
		if s <= '9' && s >= '0' {
			numCount++
		} else if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') {
			letterCount++
		} else {
			otherCount++
		}
	}
	fmt.Printf("数字个数：%d 字母个数：%d 其他字符个数：%d \n", numCount, letterCount, otherCount)
}
