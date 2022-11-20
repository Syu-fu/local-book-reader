package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"local-book-reader/domain/model"
	"local-book-reader/handler/presenter"
	"local-book-reader/usecase"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func searchBookGroups(usecase usecase.BookGroupUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error searching bookgroups"
		vars := mux.Vars(r)
		word := vars["word"]
		data, err := usecase.Search(word)
		if err != nil && err != model.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		toJ := make([]*presenter.BookGroup, 0)
		for _, d := range data {
			toJ = append(toJ, &presenter.BookGroup{
				BookId:        d.BookId,
				Title:         d.Title,
				TitleReading:  d.TitleReading,
				Author:        d.Author,
				AuthorReading: d.AuthorReading,
				Thumbnail:     d.Thumbnail,
				Tags:          d.Tags,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func getBookGroups(usecase usecase.BookGroupUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading bookgroups"
		data, err := usecase.GetList()
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
		var toJ []*presenter.BookGroup
		for _, d := range data {
			toJ = append(toJ, &presenter.BookGroup{
				BookId:        d.BookId,
				Title:         d.Title,
				TitleReading:  d.TitleReading,
				Author:        d.Author,
				AuthorReading: d.AuthorReading,
				Thumbnail:     d.Thumbnail,
				Tags:          d.Tags,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func getBookGroupById(usecase usecase.BookGroupUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading bookgroup"
		vars := mux.Vars(r)
		id := vars["book_id"]
		data, err := usecase.GetById(id)
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
		toJ := &presenter.BookGroup{
			BookId:        data.BookId,
			Title:         data.Title,
			TitleReading:  data.TitleReading,
			Author:        data.Author,
			AuthorReading: data.AuthorReading,
			Thumbnail:     data.Thumbnail,
			Tags:          data.Tags,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func createBookGroup(usecase usecase.BookGroupUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding bookgroup"
		var input struct {
			Title         string       `json:"title"`
			TitleReading  string       `json:"titleReading"`
			Author        string       `json:"author"`
			AuthorReading string       `json:"authorReading"`
			Thumbnail     string       `json:"thumnail"`
			Tags          []*model.Tag `json:"tags"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		id := model.NewID()
		book, err := model.NewBookGroup(id, input.Title, input.TitleReading, input.Author, input.AuthorReading, input.Thumbnail)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		for _, v := range input.Tags {
			book.AddTag(v)
		}

		if err := usecase.Add(book); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.BookGroup{
			BookId:        id,
			Title:         input.Title,
			TitleReading:  input.TitleReading,
			Author:        input.Author,
			AuthorReading: input.AuthorReading,
			Thumbnail:     input.Thumbnail,
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

func editBookGroup(usecase usecase.BookGroupUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error editting bookgroup"
		vars := mux.Vars(r)
		id := vars["tag_id"]
		var input struct {
			Title         string `json:"title"`
			TitleReading  string `json:"titleReading"`
			Author        string `json:"author"`
			AuthorReading string `json:"authorReading"`
			Thumbnail     string `json:"thumnail"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		bg, err := model.NewBookGroup(id, input.Title, input.TitleReading, input.Author, input.AuthorReading, input.Thumbnail)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		if err := usecase.Edit(bg); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.BookGroup{
			BookId:        id,
			Title:         input.Title,
			TitleReading:  input.TitleReading,
			Author:        input.Author,
			AuthorReading: input.AuthorReading,
			Thumbnail:     input.Thumbnail,
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

func deleteBookGroup(usecase usecase.BookGroupUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error deleting bookgroup"
		vars := mux.Vars(r)
		id := vars["tag_id"]
		err := usecase.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeBookGroupHandlers make url handlers
func MakeBookGroupHandlers(r *mux.Router, n negroni.Negroni, usecase usecase.BookGroupUsecase) {
	r.Handle("/bookgroup/search/{word}", n.With(
		negroni.Wrap(searchBookGroups(usecase)),
	)).Methods("GET", "OPTIONS").Name("searchBookGroups")

	r.Handle("/bookgroup/", n.With(
		negroni.Wrap(getBookGroups(usecase)),
	)).Methods("GET", "OPTIONS").Name("getBookGroups")

	r.Handle("/bookgroup/{book_id}", n.With(
		negroni.Wrap(getBookGroupById(usecase)),
	)).Methods("GET", "OPTIONS").Name("getBookGroupById")

	r.Handle("/bookgroup/", n.With(
		negroni.Wrap(createBookGroup(usecase)),
	)).Methods("POST", "OPTIONS").Name("createBookGroup")

	r.Handle("/bookgroup/{book_id}", n.With(
		negroni.Wrap(editBookGroup(usecase)),
	)).Methods("PUT", "OPTIONS").Name("editBookGroup")

	r.Handle("/bookgroup/{book_id}", n.With(
		negroni.Wrap(deleteBookGroup(usecase)),
	)).Methods("DELETE", "OPTIONS").Name("deleteBookGroup")
}
