package lib9

import "fmt"

// Text is the textual representation interface
type Text interface {
	~string | ~[]string
}

// Number numeric representation
type Number interface {
	~int | ~int64 | ~float64
}

// Identity is the core struct
type Identity[N Number, T Text] struct {
	ID   N
	Name T
}

// NewIdentity identity builder
func NewIdentity[N Number, T Text](n N, t T) Identity[N, T] {
	return Identity[N, T]{
		ID:   n,
		Name: t,
	}
}

func (i Identity[N, T]) String() string {
	return fmt.Sprintf("%v, %v", i.ID, i.Name)
}
