package lib1

// aNumber is the numeric type used
type aNumber interface {
	~int | ~int64 | ~int32 | ~float32 | ~float64
}

// NumericOp is the numeric operation object
type NumericOp[T aNumber] struct {
	P1 T
	P2 T
}

// NewNumericOp is the numeric builder
func NewNumericOp[T aNumber](p1, p2 T) NumericOp[T] {
	return NumericOp[T]{
		P1: p1,
		P2: p2,
	}
}

func (no NumericOp[T]) GetSum() T {
	return no.P1 + no.P2
}
