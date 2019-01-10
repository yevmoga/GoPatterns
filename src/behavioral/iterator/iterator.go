package iterator

type Iterator interface {
	Next()
	HasMore() bool
	GetCurrent() Book
}

type BookIterator struct {
	shelf *BookShelf
	index uint
}

func (i *BookIterator) Next() {
	if i.HasMore() {
		i.index++
	}
}

func (i *BookIterator) HasMore() bool {
	return i.index < uint(len(i.shelf.Books))
}

func (i *BookIterator) GetCurrent() Book {
	return i.shelf.Books[i.index]
}
