package usecase_test

import (
	"fmt"
	"local-book-reader/domain/model"
	"local-book-reader/injector"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newFixtureBook() *model.Book {
	return &model.Book{
		BookId:       "12345678-90ab-cdef-gh-ijklmnopqrst",
		Volume:       "1",
		DisplayOrder: 1,
		Thumbnail:    "path/to/thumbnail",
		Title:        "title",
		Filepath:     "path/to/filepath",
		Author:       "",
		Publisher:    "",
	}
}

func newFixtureBook2() *model.Book {
	return &model.Book{
		BookId:       "12345678-90ab-cdef-gh-ijklmnopqrst",
		Volume:       "2",
		DisplayOrder: 1,
		Thumbnail:    "path/to/thumbnail",
		Title:        "title2",
		Filepath:     "path/to/filepath",
		Author:       "",
		Publisher:    "",
	}
}

func Test_Add(t *testing.T) {
	m := injector.InjectInmemBookUsecase()
	b := newFixtureBook()
	err := m.Add(b)
	assert.Nil(t, err)
}

func Test_Get(t *testing.T) {
	m := injector.InjectInmemBookUsecase()
	b1 := newFixtureBook()
	b2 := newFixtureBook2()

	if err := m.Add(b1); err != nil {
		fmt.Print(err.Error())
	}
	if err := m.Add(b2); err != nil {
		fmt.Print(err.Error())
	}

	t.Run("getById", func(t *testing.T) {
		books, err := m.GetById("12345678-90ab-cdef-gh-ijklmnopqrst")
		assert.Nil(t, err)
		assert.Equal(t, 2, len(books))
	})

	t.Run("getByIdAndVolume", func(t *testing.T) {
		book, err := m.GetByIdAndVolume("12345678-90ab-cdef-gh-ijklmnopqrst", "2")
		assert.Nil(t, err)
		assert.Equal(t, b2.Title, book.Title)
	})
}

func Test_Update(t *testing.T) {
	m := injector.InjectInmemBookUsecase()
	b1 := newFixtureBook()
	b2 := newFixtureBook2()

	if err := m.Add(b1); err != nil {
		fmt.Print(err.Error())
	}
	if err := m.Add(b2); err != nil {
		fmt.Print(err.Error())
	}

	book, _ := m.GetByIdAndVolume("12345678-90ab-cdef-gh-ijklmnopqrst", "1")
	book.Title = "updated-title"
	assert.Nil(t, m.Edit(book))
	updated, _ := m.GetByIdAndVolume("12345678-90ab-cdef-gh-ijklmnopqrst", "1")
	unchanged, err := m.GetByIdAndVolume("12345678-90ab-cdef-gh-ijklmnopqrst", "2")
	t.Run("willUpdate", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Equal(t, book.Title, updated.Title)
	})
	t.Run("willNotUpdate", func(t *testing.T) {
		assert.NotEqual(t, book.Title, unchanged.Title)
	})
}

func Test_Delete(t *testing.T) {
	m := injector.InjectInmemBookUsecase()

	b1 := newFixtureBook()

	if err := m.Add(b1); err != nil {
		fmt.Print(err.Error())
	}

	t.Run("willDelete", func(t *testing.T) {
		assert.Nil(t, m.Delete("12345678-90ab-cdef-gh-ijklmnopqrst", "1"))
		_, err := m.GetByIdAndVolume("12345678-90ab-cdef-gh-ijklmnopqrst", "1")
		assert.Equal(t, err, model.ErrNotFound)
	})
	t.Run("willNotDelete", func(t *testing.T) {
		assert.Equal(t, m.Delete("12345678-90ab-cdef-gh-ijklmnopqrst", "2"), model.ErrNotFound)
	})
}
