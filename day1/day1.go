package main

import (
	"fmt"
	"math"
)

//1. go 环境安装
//2. 函数定义
//3. var 定义
//4. 基础类型

func add(x int, y int) int {
	return x + y

}

func main() {
	var a string
	a = "1"
	fmt.Println(a)
	fmt.Println(add(1, 2))
	print(math.Pi)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
