package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"local-book-reader/domain/model"
	"local-book-reader/handler/presenter"
	"local-book-reader/usecase"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func getBookById(usecase usecase.BookUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading books"
		var data []*model.Book
		vars := mux.Vars(r)
		id := vars["book_id"]
		data, err := usecase.GetById(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != model.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.Book
		for _, d := range data {
			toJ = append(toJ, &presenter.Book{
				BookId:       d.BookId,
				Volume:       d.Volume,
				DisplayOrder: d.DisplayOrder,
				Thumbnail:    d.Thumbnail,
				Title:        d.Title,
				Filepath:     d.Filepath,
				Author:       d.Author,
				Publisher:    d.Publisher,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func getBookByIdAndVolume(usecase usecase.BookUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading book"
		vars := mux.Vars(r)
		id := vars["book_id"]
		volume := vars["volume"]
		fmt.Println("id:" + id + " vol:" + volume)
		data, err := usecase.GetByIdAndVolume(id, volume)
		if err != nil && err != model.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Book{
			BookId:       data.BookId,
			Volume:       data.Volume,
			DisplayOrder: data.DisplayOrder,
			Thumbnail:    data.Thumbnail,
			Title:        data.Title,
			Filepath:     data.Filepath,
			Author:       data.Author,
			Publisher:    data.Publisher,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func createBook(usecase usecase.BookUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding book"
		var input struct {
			Volume       string `json:"volume"`
			DisplayOrder int    `json:"displayOrder"`
			Thumbnail    string `json:"thumnail"`
			Title        string `json:"title"`
			Filepath     string `json:"filepath"`
			Author       string `json:"author"`
			Publisher    string `json:"publisher"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		id := model.NewID()
		book, err := model.NewBook(id, input.Volume, input.DisplayOrder, input.Thumbnail, input.Title, input.Filepath, input.Author, input.Publisher)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		if err := usecase.Add(book); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Book{
			BookId:       id,
			Volume:       input.Volume,
			DisplayOrder: input.DisplayOrder,
			Thumbnail:    input.Thumbnail,
			Title:        input.Title,
			Filepath:     input.Filepath,
			Author:       input.Author,
			Publisher:    input.Publisher,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func editBook(usecase usecase.BookUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error editting book"
		var input struct {
			BookId       string `json:"bookId"`
			Volume       string `json:"volume"` //number of volumes
			DisplayOrder int    `json:"displayOrder"`
			Thumbnail    string `json:"thumnail"` //path of thumbnail
			Title        string `json:"title"`
			Filepath     string `json:"filepath"` //path of books
			Author       string `json:"author"`
			Publisher    string `json:"publisher"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		book, err := model.NewBook(input.BookId, input.Volume, input.DisplayOrder, input.Thumbnail, input.Title, input.Filepath, input.Author, input.Publisher)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		if err := usecase.Edit(book); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Book{
			BookId:       input.BookId,
			Volume:       input.Volume,
			DisplayOrder: input.DisplayOrder,
			Thumbnail:    input.Thumbnail,
			Title:        input.Title,
			Filepath:     input.Filepath,
			Author:       input.Author,
			Publisher:    input.Publisher,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func deleteBook(usecase usecase.BookUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error deleting book"
		vars := mux.Vars(r)
		id := vars["id"]
		volume := vars["volume"]
		err := usecase.Delete(id, volume)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeBookHandlers make url handlers
func MakeBookHandlers(r *mux.Router, n negroni.Negroni, usecase usecase.BookUsecase) {
	r.Handle("/book/{book_id}", n.With(
		negroni.Wrap(getBookById(usecase)),
	)).Methods("GET", "OPTIONS").Name("getBooksById")

	r.Handle("/book/{book_id}/{volume}", n.With(
		negroni.Wrap(getBookByIdAndVolume(usecase)),
	)).Methods("GET", "OPTIONS").Name("getBookByIdAndVolume")

	r.Handle("/book/", n.With(
		negroni.Wrap(createBook(usecase)),
	)).Methods("POST", "OPTIONS").Name("createBook")

	r.Handle("/book/", n.With(
		negroni.Wrap(editBook(usecase)),
	)).Methods("PUT", "OPTIONS").Name("editBook")

	r.Handle("/book/{book_id}", n.With(
		negroni.Wrap(deleteBook(usecase)),
	)).Methods("DELETE", "OPTIONS").Name("deleteBook")
}
