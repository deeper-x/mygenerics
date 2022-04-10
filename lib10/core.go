package lib10

import "fmt"

// Number numeric object
type Number interface {
	~int | ~int32 | ~int64 | ~float64
}

// Text default string representation
type Text interface {
	~string | ~[]byte
}

// Car default object
type Car[N Number, T Text] struct {
	Serial N
	Name   T
}

// NewCar builder
func NewCar[N Number, T Text](n N, t T) Car[N, T] {
	return Car[N, T]{
		Serial: n,
		Name:   t,
	}
}

func (c Car[N, T]) Details() string {
	return fmt.Sprintf("%v, %v", c.Name, c.Serial)
}
