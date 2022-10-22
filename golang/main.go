package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"local-book-reader/handler"
	"local-book-reader/infra"
	"local-book-reader/usecase"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {

	sqlHandler := infra.NewSqlHandler()
	bookRepo := infra.NewBookRepository(*sqlHandler)
	bookService := usecase.NewBookUsecase(bookRepo)

	tagRepo := infra.NewTagRepository(*sqlHandler)
	tagService := usecase.NewTagUsecase(tagRepo)

	bookGroupRepo := infra.NewBookGroupRepository(*sqlHandler)
	bookGroupService := usecase.NewBookGroupUsecase(bookGroupRepo)

	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(handler.Cors),
		negroni.NewLogger(),
	)
	//book
	handler.MakeBookHandlers(r, *n, bookService)
	handler.MakeTagHandlers(r, *n, tagService)
	handler.MakeBookGroupHandlers(r, *n, bookGroupService)

	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8000",
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err.Error())
	}
}
