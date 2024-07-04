package main

import (
	"fmt"
	"math"
)

// 流程控制 for 循环,if,switch,defer
//defer 推迟调用的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的调用会按照后进先出的顺序调用

func ifTest() {
	if 1 == 1 {
		fmt.Println("1 == 1")
	}
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	ifTest()
	fmt.Println(pow(2, 3, 10))

	defer fmt.Println("world")

	fmt.Println("hello")
}
