package presenter

import (
	"local-book-reader/domain/model"
)

type BookGroup struct {
	BookId        string       `json:"bookId"`
	Title         string       `json:"title"`
	TitleReading  string       `json:"titleReading"`
	Author        string       `json:"author"`
	AuthorReading string       `json:"authorReading"`
	Tags          []*model.Tag `json:"tags"`
}
