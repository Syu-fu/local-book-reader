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
	rows, err := bookGroupRepo.SqlHandler.Conn.Query("SELECT book_groups.book_id, book_groups.title, book_groups.title_reading, book_groups.author, book_groups.author_reading, book_groups.thumbnail, ifnull(tags.tag_id, ''), ifnull(tags.tag_name, '') " +
		"FROM book_groups LEFT OUTER JOIN tagging ON book_groups.book_id = tagging.book_id LEFT OUTER JOIN tags ON tagging.tag_id = tags.tag_id ORDER BY book_groups.book_id")
	if err != nil {
		return nil, err
	}
	cnt := -1
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
			if tag.TagId != "" {
				_ = bookgroups[cnt].AddTag(&tag)
			} else {
				bookgroups[cnt].Tags = []*model.Tag{}
			}
		} else {
			if tag.TagId != "" {
				_ = bookgroups[cnt].AddTag(&tag)
			}
		}
	}
	return bookgroups, err
}

func (bookGroupRepo *BookGroupRepository) Search(word string) ([]*model.BookGroup, error) {
	var bookgroups []*model.BookGroup
	searchWord := "%" + word + "%"
	rows, err := bookGroupRepo.SqlHandler.Conn.Query("SELECT book_groups.book_id, book_groups.title, book_groups.title_reading, book_groups.author, book_groups.author_reading, book_groups.thumbnail, ifnull(tags.tag_id, ''), ifnull(tags.tag_name, '') "+
		"FROM book_groups LEFT OUTER JOIN tagging ON book_groups.book_id = tagging.book_id LEFT OUTER JOIN tags ON tagging.tag_id = tags.tag_id "+
		"WHERE book_groups.title LIKE ? OR book_groups.title_reading LIKE ? OR book_groups.author LIKE ? OR book_groups.author_reading LIKE ? ORDER BY book_groups.book_id", searchWord, searchWord, searchWord, searchWord)
	if err != nil {
		return nil, err
	}
	cnt := -1
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
			if tag.TagId != "" {
				_ = bookgroups[cnt].AddTag(&tag)
			} else {
				bookgroups[cnt].Tags = []*model.Tag{}
				fmt.Println("emp")
			}
		} else {
			if tag.TagId != "" {
				_ = bookgroups[cnt].AddTag(&tag)
			}
		}
	}
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
		_ = bg.AddTag(tag)
	}
	return bg, err
}

func (bookGroupRepo *BookGroupRepository) Create(bg *model.BookGroup) (*model.BookGroup, error) {
	_, err := bookGroupRepo.SqlHandler.Conn.Exec(
		"INSERT INTO book_groups (book_id, title, title_reading, author, author_reading, thumbnail) VALUES (?, ?, ?, ?, ?, ?)",
		bg.BookId, bg.Title, bg.TitleReading, bg.Author, bg.AuthorReading, bg.Thumbnail)
	if err != nil {
		return nil, err
	}
	if len(bg.Tags) != 0 {
		vals := []interface{}{}
		taggingSQL := "INSERT INTO tagging (book_id, tag_id) VALUES "
		for _, v := range bg.Tags {
			fmt.Println(v.TagId)
			taggingSQL += "(?, ?),"
			vals = append(vals, bg.BookId, v.TagId)
		}
		// trim the last comma
		taggingSQL = taggingSQL[0 : len(taggingSQL)-1]

		if _, err := bookGroupRepo.SqlHandler.Conn.Exec(taggingSQL, vals...); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}

	return bg, err
}

func (bookGroupRepo *BookGroupRepository) Update(bg *model.BookGroup) (*model.BookGroup, error) {
	if _, err := bookGroupRepo.SqlHandler.Conn.Exec(
		"UPDATE book_groups SET title = ?, title_reading = ?, author = ?, author_reading = ?, thumbnail = ? WHERE book_id = ?",
		bg.Title, bg.TitleReading, bg.Author, bg.AuthorReading, bg.Thumbnail, bg.BookId); err != nil {
		return nil, err
	}

	if _, err := bookGroupRepo.SqlHandler.Conn.Exec("DELETE FROM tagging WHERE book_id = ?", bg.BookId); err != nil {
		return nil, err
	}

	if len(bg.Tags) != 0 {
		vals := []interface{}{}
		taggingSQL := "INSERT INTO tagging (book_id, tag_id) VALUES "
		for _, v := range bg.Tags {
			fmt.Println(v.TagId)
			taggingSQL += "(?, ?),"
			vals = append(vals, bg.BookId, v.TagId)
		}
		// trim the last comma
		taggingSQL = taggingSQL[0 : len(taggingSQL)-1]

		if _, err := bookGroupRepo.SqlHandler.Conn.Exec(taggingSQL, vals...); err != nil {
			return nil, err
		}
	}
	return bg, nil
}

func (bookGroupRepo *BookGroupRepository) Delete(id string) (string, error) {
	if _, err := bookGroupRepo.SqlHandler.Conn.Exec("DELETE FROM tagging WHERE book_id = ?", id); err != nil {
		return "", err
	}
	if _, err := bookGroupRepo.SqlHandler.Conn.Exec("DELETE FROM books WHERE book_id = ?", id); err != nil {
		return "", err
	}
	if _, err := bookGroupRepo.SqlHandler.Conn.Exec("DELETE FROM book_groups WHERE book_id = ?", id); err != nil {
		return "", err
	}
	return id, nil
}
