package main

import "fmt"

// 指针
//类型 *T 是指向 T 类型值的指针，其零值为 nil。
//& 操作符会生成一个指向其操作数的指针。（& 取地址）
//* 操作符表示指针指向的底层值。 （* 取值）

// 结构体
type Vertex struct {
	X int
	Y int
}

// 数组
var a [10]int

func main() {
	i := 42
	p := &i        // 指向 i
	fmt.Println(p) // 通过指针读取 i 的值

	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(v.Y)
	x := &v
	fmt.Println((*x).X)
	fmt.Println(x.X) // 隐式解引用

	fmt.Println("--------------------------------------------------")

	var (
		v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
		v2 = Vertex{X: 1}  // Y:0 被隐式地赋予零值
		v3 = Vertex{}      // X:0 Y:0
		v4 = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)

	fmt.Println(v1, v2, v3, v4.X)

	fmt.Println("---------------------")
	a[0] = 0
	fmt.Println(a)
	fmt.Println(a[0], a[1])
	fmt.Println("---------------------")

}
