package repository

import (
	"local-book-reader/domain/model"
)

type BookGroupRepository interface {
	Read() (tags []*model.BookGroup, err error)
	ReadById(id string) (tag *model.BookGroup, err error)
	Create(tag *model.BookGroup) (*model.BookGroup, error)
	Update(tag *model.BookGroup) (*model.BookGroup, error)
	Delete(id string) (string, error)
}
