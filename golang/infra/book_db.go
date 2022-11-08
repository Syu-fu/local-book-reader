package infra

import (
	"fmt"
	"local-book-reader/domain/model"
	"local-book-reader/domain/repository"
)

type BookRepository struct {
	SqlHandler
}

func NewBookRepository(sqlHandler SqlHandler) repository.BookRepository {
	bookRepository := BookRepository{sqlHandler}
	return &bookRepository
}

func (bookRepo *BookRepository) ReadById(id string) (books []*model.Book, err error) {
	rows, err := bookRepo.SqlHandler.Conn.Query("SELECT book_id, volume, display_order, thumbnail, title, filepath, author, publisher, direction FROM books WHERE book_id = ? ORDER BY display_order, volume", id)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Print(err)
		}
	}()
	for rows.Next() {
		book := model.Book{}

		if err := rows.Scan(&book.BookId, &book.Volume, &book.DisplayOrder, &book.Thumbnail, &book.Title, &book.Filepath, &book.Author, &book.Publisher, &book.Direction); err != nil {
			fmt.Print(err)
		}

		books = append(books, &book)
	}
	return
}

func (bookRepo *BookRepository) ReadByIdAndVolume(id string, volume string) (*model.Book, error) {
	var book *model.Book = new(model.Book)
	err := bookRepo.SqlHandler.Conn.QueryRow(
		"SELECT book_id, volume, display_order, thumbnail, title, filepath, author, publisher, direction FROM books WHERE book_id = ? AND volume = ?",
		id, volume).Scan(&book.BookId, &book.Volume, &book.DisplayOrder, &book.Thumbnail, &book.Title, &book.Filepath, &book.Author, &book.Publisher, &book.Direction)
	return book, err
}

func (bookRepo *BookRepository) Create(book *model.Book) (*model.Book, error) {
	_, err := bookRepo.SqlHandler.Conn.Exec(
		"INSERT INTO books (book_id, volume, display_order, thumbnail, title, filepath, author, publisher, direction) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		book.BookId, book.Volume, book.DisplayOrder, book.Thumbnail, book.Title, book.Filepath, book.Author, book.Publisher, book.Direction)
	return book, err
}

func (bookRepo *BookRepository) Update(book *model.Book) (*model.Book, error) {
	_, err := bookRepo.SqlHandler.Conn.Exec(
		"UPDATE books SET volume = ?, display_order = ?, thumbnail = ?, title = ?, filepath = ?, author = ?, publisher = ?, direction = ? WHERE book_id = ?",
		book.Volume, book.DisplayOrder, book.Thumbnail, book.Title, book.Filepath, book.Author, book.Publisher, book.Direction, book.BookId)
	return book, err
}

func (bookRepo *BookRepository) Delete(id string, volume string) (string, string, error) {
	_, err := bookRepo.SqlHandler.Conn.Exec("DELETE FROM books WHERE book_id = ? AND volume = ?", id, volume)
	return id, volume, err
}
