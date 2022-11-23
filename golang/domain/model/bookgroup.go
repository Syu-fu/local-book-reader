package model

// bookgroup is BookGroupModel
type BookGroup struct {
	BookId        string
	Title         string
	TitleReading  string
	Author        string
	AuthorReading string
	Tags          []*Tag
}

// NewBookGroup create a new bookgroup
func NewBookGroup(bookId, title, titleReading, author, authorReading string) (*BookGroup, error) {
	bg := &BookGroup{
		BookId:        bookId,
		Title:         title,
		TitleReading:  titleReading,
		Author:        author,
		AuthorReading: authorReading,
	}
	err := bg.Validate()
	if err != nil {
		return nil, err
	}
	return bg, nil
}

func (bg *BookGroup) Validate() error {
	if bg.BookId == "" || bg.Title == "" {
		return ErrInvalidEntity
	}
	return nil
}

func (bg *BookGroup) GetTag(tagId string) (*Tag, error) {
	for _, v := range bg.Tags {
		if v.TagId == tagId {
			return v, nil
		}
	}
	return nil, ErrNotFound
}

// AddTag add a tag
func (bg *BookGroup) AddTag(tag *Tag) error {
	_, err := bg.GetTag(tag.TagId)
	if err == nil {
		return ErrTagAlreadyTagging
	}
	bg.Tags = append(bg.Tags, tag)
	return nil
}

// RemoveTag remove a tag
func (bg *BookGroup) RemoveTag(tagId string) error {
	for i, j := range bg.Tags {
		if j.TagId == tagId {
			bg.Tags = append(bg.Tags[:i], bg.Tags[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
