package lib4

// Numeric is the base number interface
type Numeric interface {
	~int | ~int64 | float64
}

// Text is the textual generic interface
type Text interface {
	~string | ~[]string
}

// Message is the core message struct
type Message[N Numeric, T Text] struct {
	ID    N
	Title T
}

// NewMessage message builder
func NewMessage[N Numeric, T Text](ID N, t T) *Message[N, T] {
	return &Message[N, T]{
		ID:    ID,
		Title: t,
	}
}

// GetTitle return title for given message
func (m *Message[N, T]) GetTitle() T {
	return m.Title
}
