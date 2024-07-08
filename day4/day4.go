package main

import "fmt"

// 切片
//切片是引用，修改切片会影响到原始数组
// make 函数
//len 计算的是当前数组元素的个数，
//cap计算的是当前切片开始位到数组最后一个元素的个数

//append 追加元素

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Println(s[1:4])
	a := s[1:4]
	a[0] = 0
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(len(a))
	fmt.Println(len(s))

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Println(b, len(b), cap(b))

	s = append(s, 5)

	fmt.Println(s)

}
