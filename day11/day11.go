//package main
//
//import "golang.org/x/tour/tree"
//
//// Walk 遍历树 t，并树中所有的值发送到信道 ch。
//func Walk(t *tree.Tree, ch chan int) {
//	ch <- t.Value
//	if t.Left != nil {
//		Walk(t.Left, ch)
//	}
//	if t.Right != nil {
//		Walk(t.Right, ch)
//	}
//}
//
//// Same 判断 t1 和 t2 是否包含相同的值。
//func Same(t1, t2 *tree.Tree) bool {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	go Walk(t1, ch1)
//	go Walk(t2, ch2)
//	for {
//		v1, ok1 := <-ch1
//		v2, ok2 := <-ch2
//		if !ok1 || !ok2 {
//			return ok1 == ok2
//		}
//		if v1 != v2 {
//			return false
//		}
//	}
//	return true
//
//}
//
//func main() {
//	t1 := tree.New(1)
//	t2 := tree.New(1)
//
//	if Same(t1, t2) {
//		panic("same")
//	}
//	if !Same(t1, t2) {
//		panic("not same")
//	}
//	if Same(t1, nil) {
//		panic("not same")
//	}
//	if Same(nil, t2) {
//		panic("not same")
//	}
//
//}

package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		x, y := <-ch1, <-ch2
		if x != y {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
