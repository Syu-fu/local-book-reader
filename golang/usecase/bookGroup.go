package usecase

import (
	"local-book-reader/domain/model"
	"local-book-reader/domain/repository"
)

type BookGroupUsecase interface {
	Search(string) (bookGroups []*model.BookGroup, err error)
	GetList() (bookGroups []*model.BookGroup, err error)
	GetById(string) (bookGroup *model.BookGroup, err error)
	Add(*model.BookGroup) (err error)
	Edit(*model.BookGroup) (err error)
	Delete(string) (err error)
}

type bookGroupUsecase struct {
	bookGroupRepo repository.BookGroupRepository
}

func NewBookGroupUsecase(bookGroupRepo repository.BookGroupRepository) BookGroupUsecase {
	bookGroupUsecase := bookGroupUsecase{bookGroupRepo: bookGroupRepo}
	return &bookGroupUsecase
}

// Get a list of book group
func (usecase *bookGroupUsecase) Search(word string) (bookGroups []*model.BookGroup, err error) {
	bookGroups, err = usecase.bookGroupRepo.Search(word)
	return
}

// Get a list of book group
func (usecase *bookGroupUsecase) GetList() (bookGroups []*model.BookGroup, err error) {
	bookGroups, err = usecase.bookGroupRepo.Read()
	return
}

// Get a book group
func (usecase *bookGroupUsecase) GetById(id string) (bookGroup *model.BookGroup, err error) {
	bookGroup, err = usecase.bookGroupRepo.ReadById(id)
	return
}

// Add a book group
func (usecase *bookGroupUsecase) Add(bookGroup *model.BookGroup) (err error) {
	_, err = usecase.bookGroupRepo.Create(bookGroup)
	return
}

// Edit a book group
func (usecase *bookGroupUsecase) Edit(bookGroup *model.BookGroup) (err error) {
	_, err = usecase.bookGroupRepo.Update(bookGroup)
	return
}

// Delete a book group
func (usecase *bookGroupUsecase) Delete(id string) (err error) {
	_, err = usecase.bookGroupRepo.Delete(id)
	return
}
