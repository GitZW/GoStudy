package main

import (
	"fmt"
	"strconv"
)

// 方法
// 指针接收者的方法可以修改接收者指向的值
// 值接收者的方法只能修改接收者指向的值的副本

//带指针参数的函数必须接受一个指针
//接收者为指针的的方法被调用时，接收者既能是值又能是指针

//接受一个值作为参数的函数必须接受一个指定类型的值：
//而以值为接收者的方法被调用时，接收者既能为值又能为指针：

type Vertex struct {
	X int
	Y int
}

func (v Vertex) String() string {
	return "(" + strconv.Itoa(v.X) + "," + strconv.Itoa(v.Y) + ")"
}

func (v *Vertex) Scale(f float64) {
	v.X = int(float64(v.X) * f)
	v.Y = int(float64(v.Y) * f)

}

func ScaleFunc(v *Vertex, f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	fmt.Println(Vertex{1, 2}.String())
	v := Vertex{3, 4}
	v.Scale(2)
	fmt.Println(v.String())

	ScaleFunc(&v, 2)
	fmt.Println((&v).String())
}
