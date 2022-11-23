package model_test

import (
	"local-book-reader/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBookGroup(t *testing.T) {
	bg, err := model.NewBookGroup(model.NewID(), "The Hitchhiker's Guide to the Galaxy", "The Hitchhiker's Guide to the Galaxy", "Douglas Adams", "Douglas Adams")
	assert.Nil(t, err)
	assert.Equal(t, bg.Title, "The Hitchhiker's Guide to the Galaxy")
	assert.NotNil(t, bg.BookId)
}

func TestBookGroupValidate(t *testing.T) {
	type test struct {
		bookId string
		title  string
		want   error
	}

	tests := []test{
		{
			bookId: model.NewID(),
			title:  "The Hitchhiker's Guide to the Galaxy",
			want:   nil,
		},
		{
			bookId: "",
			title:  "The Hitchhiker's Guide to the Galaxy",
			want:   model.ErrInvalidEntity,
		},
		{
			bookId: model.NewID(),
			title:  "",
			want:   model.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := model.NewBookGroup(tc.bookId, tc.title, "titleReading", "Author", "AuthorReading")
		assert.Equal(t, err, tc.want)
	}
}

func TestAddTag(t *testing.T) {
	bg, _ := model.NewBookGroup(model.NewID(), "The Hitchhiker's Guide to the Galaxy", "The Hitchhiker's Guide to the Galaxy", "Douglas Adams", "Douglas Adams")
	tID := model.NewID()
	tag, _ := model.NewTag(tID, "tagname")
	err := bg.AddTag(tag)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(bg.Tags))
	err = bg.AddTag(tag)
	assert.Equal(t, model.ErrTagAlreadyTagging, err)
}

func TestRemoveTag(t *testing.T) {
	bg, _ := model.NewBookGroup(model.NewID(), "The Hitchhiker's Guide to the Galaxy", "The Hitchhiker's Guide to the Galaxy", "Douglas Adams", "Douglas Adams")
	err := bg.RemoveTag(model.NewID())
	assert.Equal(t, model.ErrNotFound, err)
	tID := model.NewID()
	tag, _ := model.NewTag(tID, "tagname")
	_ = bg.AddTag(tag)
	err = bg.RemoveTag(tID)
	assert.Nil(t, err)
}

func TestGetTag(t *testing.T) {
	bg, _ := model.NewBookGroup(model.NewID(), "The Hitchhiker's Guide to the Galaxy", "The Hitchhiker's Guide to the Galaxy", "Douglas Adams", "Douglas Adams")
	tID := model.NewID()
	tag, _ := model.NewTag(tID, "tagname")
	_ = bg.AddTag(tag)
	outTag, err := bg.GetTag(tID)
	assert.Nil(t, err)
	assert.Equal(t, tag, outTag)
	_, err = bg.GetTag(model.NewID())
	assert.Equal(t, model.ErrNotFound, err)
}
