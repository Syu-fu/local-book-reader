package usecase

import (
	"local-book-reader/domain/model"
	"local-book-reader/domain/repository"
)

type BookUsecase interface {
	GetById(string) (book []*model.Book, err error)
	GetByIdAndVolume(string, string) (book *model.Book, err error)
	Add(*model.Book) (err error)
	Edit(*model.Book) (err error)
	Delete(string, string) (err error)
}

type bookUsecase struct {
	bookRepo repository.BookRepository
}

func NewBookUsecase(bookRepo repository.BookRepository) BookUsecase {
	bookUsecase := bookUsecase{bookRepo: bookRepo}
	return &bookUsecase
}

// Get a list of books in the book group
func (usecase *bookUsecase) GetById(id string) (book []*model.Book, err error) {
	book, err = usecase.bookRepo.ReadById(id)
	return
}

// Get a book
func (usecase *bookUsecase) GetByIdAndVolume(id string, volume string) (book *model.Book, err error) {
	book, err = usecase.bookRepo.ReadByIdAndVolume(id, volume)
	return
}

// Add a book
func (usecase *bookUsecase) Add(book *model.Book) (err error) {
	_, err = usecase.bookRepo.Create(book)
	return
}

// Edit a book
func (usecase *bookUsecase) Edit(book *model.Book) (err error) {
	_, err = usecase.bookRepo.Update(book)
	return
}

// Delete a book
func (usecase *bookUsecase) Delete(id string, volume string) (err error) {
	_, _, err = usecase.bookRepo.Delete(id, volume)
	return
}
