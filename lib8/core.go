package lib8

import "fmt"

// Number is the numeric type
type Number interface {
	~int | ~int64 | ~float64
}

// Text every kind of string or slice representation
type Text interface {
	~string | ~[]string
}

// Log is the logging object
type Log[N Number, T Text] struct {
	ID   N
	Text T
}

// NewLog is the log builder
func NewLog[N Number, T Text](n N, t T) Log[N, T] {
	return Log[N, T]{
		ID:   n,
		Text: t,
	}
}

func (l Log[N, T]) Details() string {
	return fmt.Sprintf("%v %v", l.ID, l.Text)
}
