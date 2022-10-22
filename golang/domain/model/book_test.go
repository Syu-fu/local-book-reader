package model_test

import (
	"local-book-reader/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	b, err := model.NewBook(model.NewID(), "1", 1, "path/to/thumbnail", "The Hitchhiker's Guide to the Galaxy", "path/to/filepath", "Douglas Adams", "Pan Books", "ltr")
	assert.Nil(t, err)
	assert.Equal(t, b.Title, "The Hitchhiker's Guide to the Galaxy")
}

func TestBookValidate(t *testing.T) {
	type test struct {
		id        string
		title     string
		volume    string
		filepath  string
		direction string
		want      error
	}

	tests := []test{
		{
			id:        model.NewID(),
			title:     "The Hitchhiker's Guide to the Galaxy",
			volume:    "1",
			filepath:  "path/to/filepath",
			direction: "ltr",
			want:      nil,
		},
		{
			id:        model.NewID(),
			title:     "",
			volume:    "1",
			filepath:  "path/to/filepath",
			direction: "ltr",
			want:      model.ErrInvalidEntity,
		},
		{
			id:        model.NewID(),
			title:     "The Hitchhiker's Guide to the Galaxy",
			volume:    "",
			filepath:  "path/to/filepath",
			direction: "ltr",
			want:      model.ErrInvalidEntity,
		},
		{
			id:        model.NewID(),
			title:     "The Hitchhiker's Guide to the Galaxy",
			volume:    "1",
			filepath:  "",
			direction: "ltr",
			want:      model.ErrInvalidEntity,
		},
		{
			id:        "",
			title:     "The Hitchhiker's Guide to the Galaxy",
			volume:    "1",
			filepath:  "path/to/filepath",
			direction: "ltr",
			want:      model.ErrInvalidEntity,
		},
		{
			id:        model.NewID(),
			title:     "The Hitchhiker's Guide to the Galaxy",
			volume:    "1",
			filepath:  "path/to/filepath",
			direction: "",
			want:      model.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := model.NewBook(tc.id, tc.volume, 0, "thumbnail", tc.title, tc.filepath, "author", "publisher", tc.direction)
		assert.Equal(t, err, tc.want)
	}

}
