package main

import (
	"log"

	"github.com/deeper-x/mygenerics/lib1"
	"github.com/deeper-x/mygenerics/lib2"
	"github.com/deeper-x/mygenerics/lib3"
	"github.com/deeper-x/mygenerics/lib4"
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

	// lib1 demo
	o2 := lib1.NewNumericOp(10, 30)
	log.Println(o2.GetSum())

	o3 := lib1.NewNumericOp(20.1, 30.1)
	log.Println(o3.GetSum())

	o4 := lib2.NewBook(10.1, "doo")
	o5 := o4.GetInstance()
	log.Println(o5.Author)

	o6 := lib2.NewBook(10, []string{"one", "two"})
	log.Println(o6)

	o7 := lib3.NewLog([]string{"demo foo"})
	log.Println(o7.GetArchive())

	toadd := []string{"one", "two"}
	o7.AddArchive(toadd)

	o8 := lib4.NewMessage(10, "title demo")
	log.Println(o8.GetTitle())
}

// typical G usage on funcs...
func sumInts(p1, p2 int) int {
	return p1 + p2
}

// ... with numbers manipulation
func sumNums[T myNum](p1, p2 T) T {
	return p1 + p2
}

// struct pattern
func newOperator[T myNum](x T, y T) operator[T, T] {
	return operator[T, T]{p1: x, p2: y}
}

func (o operator[P1, P2]) SaySomething() {
	log.Println(o.p1, o.p2)
}
