package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int, flag string) {
	sum := 0
	for _, v := range s {
		fmt.Println(flag)
		sum += v
	}
	c <- sum // 发送 sum 到 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c, "1")
	go sum(s[len(s)/2:], c, "2")
	x, y := <-c, <-c // 从 c 接收

	fmt.Println(x, y, x+y)

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
