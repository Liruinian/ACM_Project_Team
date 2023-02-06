package main

import (
	"fmt"
)

func main() {
	//demo
	l := NewDuLinkList()
	x1 := l.PushBack(3)
	x2 := l.PushBack(4)
	x5 := l.PushBack(114514)
	x3 := l.PushFront(2)
	x4 := l.PushFront(1)

	l.Remove(x5)
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(l)
	fmt.Printf("l.Len() = %d\n", l.Len())
	fmt.Printf("l.First() = %+v\n", l.First())
	fmt.Printf("l.Last() = %+v\n", l.Last())
	fmt.Println("----------")
	l.Init()
	a1 := l.PushBack(-2)
	a2 := l.PushBack(-1)
	a5 := l.PushBack(1919810)
	l0 := NewDuLinkList()
	a3 := l0.PushFront(-3)
	a4 := l0.PushFront(-4)
	l.PushFrontList(l0)
	l.Remove(a5)
	l.PopFirst()
	l.PopLast()

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(l)
	fmt.Println(l.Len())
	fmt.Println(l.First())

	fmt.Println("------------------")
	fmt.Println(x5)
	fmt.Println(a5)
}
