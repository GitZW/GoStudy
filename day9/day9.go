package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Append(val T) {
	l.next = &List[T]{val: val}
}

func (l *List[T]) Length() int {
	i := 0
	for l.next != nil {
		i++
		l = l.next
	}
	return i

}

func main() {
	l := List[int]{}
	l.Append(1)
	fmt.Println(l.Length())

}
