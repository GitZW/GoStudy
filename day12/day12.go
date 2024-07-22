package main

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 3

	if <-ch1 == <-ch2 {
		println("true")
	} else {
		println("false")
	}

}
