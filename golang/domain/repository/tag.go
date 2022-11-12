package repository

import (
	"local-book-reader/domain/model"
)

//go:generate mockhandler -source=$GOFILE -destination=./mock/$GOFILE -package=mock
type TagRepository interface {
	Read() (tags []*model.Tag, err error)
	ReadById(id string) (tag *model.Tag, err error)
	ReadByName(name string) (tag *model.Tag, err error)
	Search(name string) (tags []*model.Tag, err error)
	Create(tag *model.Tag) (*model.Tag, error)
	Update(tag *model.Tag) (*model.Tag, error)
	Delete(id string) (string, error)
}
