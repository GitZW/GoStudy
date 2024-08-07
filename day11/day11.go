package main

import "golang.org/x/tour/tree"

// Walk 遍历树 t，并树中所有的值发送到信道 ch。
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t != nil {
		ch <- t.Value
	}

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 判断 t1 和 t2 是否包含相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)

	if Same(t1, t2) {
		println("same")
	}
	if !Same(t1, t2) {
		println("not same")
	}

}
