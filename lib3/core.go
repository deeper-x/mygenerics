package lib3

// Stringer is the string based interface
type Stringer interface {
	~string | ~[]string
}

// Log is the default log
type Log[T Stringer] struct {
	Entry   T
	Archive T
}

// NewLog is the Log builder
func NewLog[T Stringer](t T) *Log[T] {
	return &Log[T]{
		Entry: t,
	}
}

func (l *Log[T]) GetEntry() T {
	return l.Entry
}

// GetArchive returns all log entries
func (l *Log[T]) GetArchive() T {
	return l.Archive
}
