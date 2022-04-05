package lib6

import (
	"fmt"
)

// Number is a generic interface
type Number interface {
	~int | ~int64 | ~float64
}

// Text is the message interface
type Text interface {
	~string | ~int
}

// Identity is the core object
type Identity[N Number, T Text] struct {
	ID   N
	Name T
}

// NewIdentity returns identity instance
func NewIdentity[N Number, T Text](n N, t T) Identity[N, T] {
	return Identity[N, T]{
		ID:   n,
		Name: t,
	}
}

// Print returns identity data
func (i Identity[N, T]) Print() string {
	return fmt.Sprintf("%v->%v", i.ID, i.Name)
}
