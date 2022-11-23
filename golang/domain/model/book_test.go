package model_test

import (
	"local-book-reader/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	b, err := model.NewBook(model.NewID(), "1", 1, "The Hitchhiker's Guide to the Galaxy", "Douglas Adams", "Pan Books", "ltr")
	assert.Nil(t, err)
	assert.Equal(t, b.Title, "The Hitchhiker's Guide to the Galaxy")
}

func TestBookValidate(t *testing.T) {
	type test struct {
		id        string
		title     string
		volume    string
		direction string
		want      error
	}

	tests := []test{
		{
			id:        model.NewID(),
			title:     "The Hitchhiker's Guide to the Galaxy",
			volume:    "1",
			direction: "ltr",
			want:      nil,
		},
		{
			id:        model.NewID(),
			title:     "",
			volume:    "1",
			direction: "ltr",
			want:      model.ErrInvalidEntity,
		},
		{
			id:        model.NewID(),
			title:     "The Hitchhiker's Guide to the Galaxy",
			volume:    "",
			direction: "ltr",
			want:      model.ErrInvalidEntity,
		},
		{
			id:        "",
			title:     "The Hitchhiker's Guide to the Galaxy",
			volume:    "1",
			direction: "ltr",
			want:      model.ErrInvalidEntity,
		},
		{
			id:        model.NewID(),
			title:     "The Hitchhiker's Guide to the Galaxy",
			volume:    "1",
			direction: "",
			want:      model.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := model.NewBook(tc.id, tc.volume, 0, tc.title, "author", "publisher", tc.direction)
		assert.Equal(t, err, tc.want)
	}

}
