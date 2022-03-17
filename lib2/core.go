package lib2

// ID is the type for book ID management
type ID interface {
	~int | ~float64
}

// Name is the name management type
type Name interface {
	~string | ~[]string
}

// Book is the central object
type Book[T1 ID, T2 Name] struct {
	Pages  T1
	Author T2
}

// NewBook is the book builder
func NewBook[T1 ID, T2 Name](p T1, a T2) *Book[T1, T2] {
	return &Book[T1, T2]{
		Pages:  p,
		Author: a,
	}
}

func (b *Book[T1, T2]) GetInstance() *Book[T1, T2] {
	return &Book[T1, T2]{
		Pages:  b.Pages,
		Author: b.Author,
	}
}
