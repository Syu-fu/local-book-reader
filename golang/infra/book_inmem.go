package infra

import (
	"local-book-reader/domain/model"
)

type Inmem struct {
	m []*model.Book
}

func NewInmem() *Inmem {
	var m = []*model.Book{}
	return &Inmem{
		m: m,
	}
}

func (inmem *Inmem) ReadById(id string) ([]*model.Book, error) {
	var d []*model.Book
	for i := 0; i < len(inmem.m); i++ {
		if inmem.m[i].BookId == id {
			d = append(d, inmem.m[i])
		}
	}
	if len(d) == 0 {
		return nil, model.ErrNotFound
	}
	return d, nil
}

func (inmem *Inmem) ReadByIdAndVolume(id string, volume string) (*model.Book, error) {
	var b *model.Book
	for i := 0; i < len(inmem.m); i++ {
		if inmem.m[i].BookId == id && inmem.m[i].Volume == volume {
			b = inmem.m[i]
		}
	}
	if b == nil {
		return nil, model.ErrNotFound
	}
	return b, nil
}

func (inmem *Inmem) Create(book *model.Book) (*model.Book, error) {
	inmem.m = append(inmem.m, book)
	return book, nil
}

func (inmem *Inmem) Update(book *model.Book) (*model.Book, error) {
	for i := 0; i < len(inmem.m); i++ {
		if book.BookId == inmem.m[i].BookId {
			inmem.m[i] = book
			return book, nil
		}
	}
	return nil, model.ErrNotFound
}

func (inmem *Inmem) Delete(id string, volume string) (string, string, error) {
	for i := 0; i < len(inmem.m); i++ {
		if id == inmem.m[i].BookId {
			inmem.m = append(inmem.m[:i], inmem.m[i+1:]...)
			return id, volume, nil
		}
	}
	return "", "", model.ErrNotFound
}
