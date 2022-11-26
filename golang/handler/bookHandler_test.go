package handler

import (
	"testing"

	mock "local-book-reader/usecase/mock"

	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_getBooksById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock.NewMockBookUsecase(ctrl)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookHandlers(r, *n, service)
	path, err := r.GetRoute("getBooksById").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/book/{book_id}", path)
}

func Test_getBookByIdAndVolume(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock.NewMockBookUsecase(ctrl)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookHandlers(r, *n, service)
	path, err := r.GetRoute("getBookByIdAndVolume").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/book/{book_id}/{volume}", path)
}

func Test_createBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	bookUsecase := mock.NewMockBookUsecase(ctrl)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookHandlers(r, *n, bookUsecase)
	path, err := r.GetRoute("createBook").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/book/", path)
}

func Test_updateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	bookUsecase := mock.NewMockBookUsecase(ctrl)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookHandlers(r, *n, bookUsecase)
	path, err := r.GetRoute("editBook").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/book/{book_id}/{volume}", path)
}

func Test_deleteBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mock.NewMockBookUsecase(ctrl)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookHandlers(r, *n, service)
	path, err := r.GetRoute("deleteBook").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/book/{book_id}/{volume}", path)
}
