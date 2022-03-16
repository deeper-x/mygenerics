package main

import "log"

// MyNumber is a number
type MyNumber interface {
	~int | ~float64
}

type obj[A, B MyNumber] struct {
	a A
	b B
}

func main() {
	var x = 2
	var y = 3.14

	log.Println(doubleIt(x))
	log.Println(doubleIt(y))
}

func doubleIt[T MyNumber](p T) T {
	return p * 2
}

func (p obj[A, B]) SayHello() A {
	return p.a
}
