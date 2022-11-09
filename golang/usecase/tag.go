package usecase

import (
	"local-book-reader/domain/model"
	"local-book-reader/domain/repository"
)

type TagUsecase interface {
	Get() (tags []*model.Tag, err error)
	GetById(string) (tag *model.Tag, err error)
	Search(string) (tags []*model.Tag, err error)
	Add(*model.Tag) (err error)
	Edit(*model.Tag) (err error)
	Delete(string) (err error)
}

type tagUsecase struct {
	tagRepo repository.TagRepository
}

func NewTagUsecase(tagRepo repository.TagRepository) TagUsecase {
	tagUsecase := tagUsecase{tagRepo: tagRepo}
	return &tagUsecase
}

// Get a list of tags in the book group
func (usecase *tagUsecase) Search(name string) (tags []*model.Tag, err error) {
	tags, err = usecase.tagRepo.Search(name)
	return
}

// Get a list of tags in the book group
func (usecase *tagUsecase) Get() (tags []*model.Tag, err error) {
	tags, err = usecase.tagRepo.Read()
	return
}

// Get a list of tags in the book group
func (usecase *tagUsecase) GetById(id string) (tag *model.Tag, err error) {
	tag, err = usecase.tagRepo.ReadById(id)
	return
}

// Add a tag
func (usecase *tagUsecase) Add(tag *model.Tag) (err error) {
	_, err = usecase.tagRepo.Create(tag)
	return
}

// Edit a tag
func (usecase *tagUsecase) Edit(tag *model.Tag) (err error) {
	_, err = usecase.tagRepo.Update(tag)
	return
}

// Delete a tag
func (usecase *tagUsecase) Delete(id string) (err error) {
	_, err = usecase.tagRepo.Delete(id)
	return
}
