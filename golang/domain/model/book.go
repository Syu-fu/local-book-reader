package model

//Book is BookModel
type Book struct {
	BookId       string
	Volume       string
	DisplayOrder int
	Thumbnail    string
	Title        string
	Filepath     string
	Author       string
	Publisher    string
	Direction    string
}

//NewBook create a new book
func NewBook(bookId, volume string, displayOrder int, thumbnail, title, filepath, author, publisher, direction string) (*Book, error) {
	b := &Book{
		BookId:       bookId,
		Volume:       volume,
		DisplayOrder: displayOrder,
		Thumbnail:    thumbnail,
		Title:        title,
		Filepath:     filepath,
		Author:       author,
		Publisher:    publisher,
		Direction:    direction,
	}
	err := b.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return b, nil
}

//Validate validate book
func (b *Book) Validate() error {
	if b.BookId == "" || b.Volume == "" || b.Title == "" || b.Filepath == "" || b.Direction == "" {
		return ErrInvalidEntity
	}
	return nil
}
