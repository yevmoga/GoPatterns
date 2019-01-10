package iterator

type IterableCollection interface {
	Iterator() Iterator
}

type BookShelf struct {
	Books []Book
}

func (b *BookShelf) Iterator() Iterator {
	return &BookIterator{shelf: b}
}

func (b *BookShelf) Add(book Book) {
	b.Books = append(b.Books, book)
	return
}

