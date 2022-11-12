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

func getTags(usecase usecase.TagUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading tag"
		data, err := usecase.Get()
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
		var toJ []*presenter.Tag
		for _, d := range data {
			toJ = append(toJ, &presenter.Tag{
				TagId:   d.TagId,
				TagName: d.TagName,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func searchTags(usecase usecase.TagUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error searching tag"
		vars := mux.Vars(r)
		name := vars["tag_name"]
		data, err := usecase.Search(name)
		if err != nil && err != model.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		toJ := make([]*presenter.Tag, 0)
		for _, d := range data {
			toJ = append(toJ, &presenter.Tag{
				TagId:   d.TagId,
				TagName: d.TagName,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func getTagById(usecase usecase.TagUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading tag"
		vars := mux.Vars(r)
		id := vars["tag_id"]
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
		toJ := &presenter.Tag{
			TagId:   data.TagId,
			TagName: data.TagName,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func getTagByName(usecase usecase.TagUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading tag"
		vars := mux.Vars(r)
		name := vars["tag_name"]
		data, err := usecase.GetByName(name)
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
		if data.TagId == "" {
			toJ := []string{}
			if err := json.NewEncoder(w).Encode(toJ); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(errorMessage))
			}
			return
		}
		toJ := &presenter.Tag{
			TagId:   data.TagId,
			TagName: data.TagName,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func createTag(usecase usecase.TagUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding book"
		var input struct {
			TagName string `json:"tagName"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		id := model.NewID()
		book, err := model.NewTag(id, input.TagName)
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
		toJ := &presenter.Tag{
			TagId:   id,
			TagName: input.TagName,
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

func editTag(usecase usecase.TagUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error editting tag"
		vars := mux.Vars(r)
		id := vars["tag_id"]
		var input struct {
			TagName string `json:"tagName"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		tag, err := model.NewTag(id, input.TagName)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		if err := usecase.Edit(tag); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Tag{
			TagId:   id,
			TagName: input.TagName,
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

func deleteTag(usecase usecase.TagUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error deleting tag"
		fmt.Println("deletemethod")
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

// MakeTagHandlers make url handlers
func MakeTagHandlers(r *mux.Router, n negroni.Negroni, usecase usecase.TagUsecase) {
	r.Handle("/tag/", n.With(
		negroni.Wrap(getTags(usecase)),
	)).Methods("GET", "OPTIONS").Name("getTags")

	r.Handle("/tag/search/q={tag_name}", n.With(
		negroni.Wrap(searchTags(usecase)),
	)).Methods("GET", "OPTIONS").Name("searchTags")

	r.Handle("/tag/{tag_id}", n.With(
		negroni.Wrap(getTagById(usecase)),
	)).Methods("GET", "OPTIONS").Name("getTag")

	r.Handle("/tag/name/{tag_name}", n.With(
		negroni.Wrap(getTagByName(usecase)),
	)).Methods("GET", "OPTIONS").Name("getTagByName")

	r.Handle("/tag/", n.With(
		negroni.Wrap(createTag(usecase)),
	)).Methods("POST", "OPTIONS").Name("createTag")

	r.Handle("/tag/{tag_id}", n.With(
		negroni.Wrap(editTag(usecase)),
	)).Methods("PUT", "OPTIONS").Name("editTag")

	r.Handle("/tag/{tag_id}", n.With(
		negroni.Wrap(deleteTag(usecase)),
	)).Methods("DELETE", "OPTIONS").Name("deleteTag")
}
