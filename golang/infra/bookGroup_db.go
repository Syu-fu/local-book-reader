package infra

import (
	"fmt"
	"local-book-reader/domain/model"
	"local-book-reader/domain/repository"
)

type BookGroupRepository struct {
	SqlHandler
}

func NewBookGroupRepository(sqlHandler SqlHandler) repository.BookGroupRepository {
	bookGroupRepository := BookGroupRepository{sqlHandler}
	return &bookGroupRepository
}

func (bookGroupRepo *BookGroupRepository) Read() ([]*model.BookGroup, error) {
	var bookgroups []*model.BookGroup
	rows, err := bookGroupRepo.SqlHandler.Conn.Query("SELECT book_groups.book_id, book_groups.title, book_groups.title_reading, book_groups.author, book_groups.author_reading, book_groups.thumbnail, tags.tag_id, tags.tag_name " +
		"FROM book_groups LEFT OUTER JOIN tagging ON book_groups.book_id JOIN tags ON tagging.tag_id = tags.tag_id ORDER BY book_groups.book_id")
	if err != nil {
		return nil, err
	}
	cnt := -1
	fmt.Println("getList running")
	for rows.Next() {
		var bg model.BookGroup
		var tag model.Tag
		err = rows.Scan(&bg.BookId, &bg.Title, &bg.TitleReading, &bg.Author, &bg.AuthorReading, &bg.Thumbnail, &tag.TagId, &tag.TagName)
		if err != nil {
			return nil, err
		}
		if cnt == -1 || bookgroups[cnt].BookId != bg.BookId {
			bookgroups = append(bookgroups, &bg)
			cnt++
			bookgroups[cnt].AddTag(&tag)
		} else {
			bookgroups[cnt].AddTag(&tag)
		}
	}
	fmt.Print("bookgroup list length:")
	fmt.Sprintln(len(bookgroups))
	return bookgroups, err
}

func (bookGroupRepo *BookGroupRepository) ReadById(id string) (*model.BookGroup, error) {
	var bg *model.BookGroup = new(model.BookGroup)
	rows, err := bookGroupRepo.SqlHandler.Conn.Query("SELECT book_groups.book_id, book_groups.title, book_groups.title_reading, book_groups.author, book_groups.author_reading, book_groups.thumbnail, tags.tag_id, tags.tag_name "+
		"FROM book_groups LEFT OUTER JOIN tagging ON book_groups.book_id JOIN tags ON tagging.tag_id = tags.tag_id WHERE book_groups.book_id = ? ORDER BY book_groups.book_id", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tag *model.Tag = new(model.Tag)
		err = rows.Scan(&bg.BookId, &bg.Title, &bg.TitleReading, &bg.Author, &bg.AuthorReading, &bg.Thumbnail, &tag.TagId, &tag.TagName)
		if err != nil {
			return nil, err
		}
		bg.AddTag(tag)
	}
	return bg, err
}

func (bookGroupRepo *BookGroupRepository) Create(bg *model.BookGroup) (*model.BookGroup, error) {
	_, err := bookGroupRepo.SqlHandler.Conn.Exec(
		"INSERT INTO book_groups (book_id, title, title_reading, author, author_reading, thumbnail) VALUES (?, ?, ?, ?, ?, ?)",
		bg.BookId, bg.Title, bg.TitleReading, bg.Author, bg.AuthorReading, bg.Thumbnail)
	fmt.Println(err.Error())
	return bg, err
}

func (bookGroupRepo *BookGroupRepository) Update(bg *model.BookGroup) (*model.BookGroup, error) {
	_, err := bookGroupRepo.SqlHandler.Conn.Exec(
		"UPDATE book_groups SET title = ?, title_reading = ?, author = ?, author_reading = ?, thumbnail = ? WHERE book_id = ?",
		bg.Title, bg.TitleReading, bg.Author, bg.AuthorReading, bg.Thumbnail, bg.BookId)
	return bg, err
}

func (bookGroupRepo *BookGroupRepository) Delete(id string) (string, error) {
	_, err := bookGroupRepo.SqlHandler.Conn.Exec("DELETE FROM book_groups WHERE book_id = ?", id)
	return id, err
}
