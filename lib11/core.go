package lib11

import "fmt"

// Number is numeric representation
type Number interface {
	~int64 | ~float64 | ~int
}

// Text string representation
type Text interface {
	~string | ~[]byte
}

// Product object generalized
type Product[N Number, T Text] struct {
	Serial T
	Name   T
	Price  N
}

// NewProduct builder
func NewProduct[N Number, T Text](s T, n T, p N) *Product[N, T] {
	return &Product[N, T]{
		Serial: s,
		Name:   n,
		Price:  p,
	}
}

func (p *Product[N, T]) String() string {
	return fmt.Sprintf("%v, %v, [%v]", p.Name, p.Price, p.Serial)
}
