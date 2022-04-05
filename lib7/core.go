package lib7

import "fmt"

// Number is the generic numeric type
type Number interface {
	~int | ~int64 | ~int32
}

// Text is the textual name's representation
type Text interface {
	~int | ~int64 | ~int32 | ~string
}

// Person represent the main identity
type Person[N Number, T Text] struct {
	ID   N
	Name T
}

// NewPerson builder
func NewPerson[N Number, T Text](n N, t T) *Person[N, T] {
	return &Person[N, T]{
		ID:   n,
		Name: t,
	}
}

func (p Person[N, T]) Details() string {
	return fmt.Sprintf("%v->%v", p.ID, p.Name)
}
