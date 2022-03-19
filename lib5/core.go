package lib5

import "fmt"

// Number generic interface
type Number interface {
	~int | ~int64 | ~int32 | ~float64
}

// Text generic interface
type Text interface {
	~string | ~[]string
}

// Stringer is the interface with string method
type Stringer interface {
	String() string
}

// Log base struct
type Log[N Number, T Text] struct {
	ID      N
	Message T
}

// NewLog is the log builder
func NewLog[N Number, T Text](id N, txt T) Log[N, T] {
	return Log[N, T]{
		ID:      id,
		Message: txt,
	}
}

// String is the formatting method
func (l *Log[N, T]) String() string {
	return fmt.Sprintf("%v->%s", l.ID, l.Message)
}

// Update message value
func (l *Log[N, T]) Update(v T) {
	l.Message = v
}
