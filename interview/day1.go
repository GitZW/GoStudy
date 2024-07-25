package main

import "fmt"

// for循环select时，如果通道已经关闭会怎么样？如果select中的case只有一个，又会怎么样？

// 读已经关闭的 chan 能一直读到东西，但是读到的内容根据通道内关闭前是否有元素而不同。
// 如果 chan 关闭前，buffer 内有元素还未读 , 会正确读到 chan 内的值，且返回的第二个 bool 值（是否读成功）为 true。
// 如果 chan 关闭前，buffer 内有元素已经被读完，chan 内无值，接下来所有接收的值都会非阻塞直接成功，返回 channel 元素的零值，但是第二个 bool 值一直为 false。
// 写已经关闭的 chan 会 panic

// for 循环select时，如果其中一个 case 通道已经关闭，则每次都会执行到这个 case。
// 如果 select 里边只有一个 case，而这个 case 被关闭了，则会出现死循环。
// select中如果任意某个通道有值可读时，它就会被执行，其他被忽略。
// 如果没有default字句，select将有可能阻塞，直到某个通道有值可以运行，所以select里最好有一个default，否则将有一直阻塞的风险。
func main() {
	ch := make(chan int, 1)

	ch <- 1
	close(ch)

	fmt.Println(".....")
	fmt.Println(<-ch)
	fmt.Println(".....")
	fmt.Println(<-ch)
	fmt.Println(".....")
	fmt.Println(<-ch)
	fmt.Println(".....")
	fmt.Println(<-ch)
	fmt.Println(".....")
	fmt.Println(<-ch)

}
