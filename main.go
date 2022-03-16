package main

import (
	"log"
)

type myNum interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

type operator[P1, P2 myNum] struct {
	p1 P1
	p2 P2
}

type operSimple struct {
	p1 int64
	p2 int64
}

func main() {
	log.Println(sumNums(10, 20))
	log.Println(sumNums(29.1, 31.2))

	o := newOperator(1.1, 2.2)
	o.SaySomething()
}

func newOperator[T myNum](x T, y T) operator[T, T] {
	return operator[T, T]{p1: x, p2: y}
}

func newOperSimple(x, y int64) operSimple {
	return operSimple{p1: x, p2: y}
}

func sumInts(p1, p2 int) int {
	return p1 + p2
}

func sumNums[T myNum](p1, p2 T) T {
	return p1 + p2
}

func (o operator[P1, P2]) SaySomething() {
	log.Println(o.p1, o.p2)
}
