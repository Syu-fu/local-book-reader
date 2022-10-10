package repository

import (
	"local-book-reader/domain/model"
)

//go:generate mockhandler -source=$GOFILE -destination=./mock/$GOFILE -package=mock
type BookRepository interface {
	ReadById(id string) (books []*model.Book, err error)
	ReadByIdAndVolume(id string, volume string) (books *model.Book, err error)
	Create(book *model.Book) (*model.Book, error)
	Update(book *model.Book) (*model.Book, error)
	Delete(id string, volume string) (string, string, error)
}
