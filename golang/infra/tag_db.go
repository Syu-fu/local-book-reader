package infra

import (
	"local-book-reader/domain/model"
	"local-book-reader/domain/repository"
)

type TagRepository struct {
	SqlHandler
}

func NewTagRepository(sqlHandler SqlHandler) repository.TagRepository {
	tagRepository := TagRepository{sqlHandler}
	return &tagRepository
}

func (tagRepo *TagRepository) Read() ([]*model.Tag, error) {
	var tags []*model.Tag
	rows, err := tagRepo.SqlHandler.Conn.Query("SELECT tag_id, tag_name FROM tags")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t model.Tag
		err = rows.Scan(&t.TagId, &t.TagName)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &t)
	}
	return tags, err
}

func (tagRepo *TagRepository) ReadById(id string) (*model.Tag, error) {
	var tag *model.Tag = new(model.Tag)
	err := tagRepo.SqlHandler.Conn.QueryRow(
		"SELECT tag_id, tag_name FROM tags WHERE tag_id = ?", id).Scan(&tag.TagId, &tag.TagName)
	return tag, err
}

func (tagRepo *TagRepository) Create(tag *model.Tag) (*model.Tag, error) {
	_, err := tagRepo.SqlHandler.Conn.Exec(
		"INSERT INTO tags (tag_id, tag_name) VALUES (?, ?)",
		tag.TagId, tag.TagName)
	return tag, err
}

func (tagRepo *TagRepository) Update(tag *model.Tag) (*model.Tag, error) {
	_, err := tagRepo.SqlHandler.Conn.Exec(
		"UPDATE tags SET tag_name = ? WHERE tag_id = ?",
		tag.TagName, tag.TagId)
	return tag, err
}

func (tagRepo *TagRepository) Delete(id string) (string, error) {
	_, err := tagRepo.SqlHandler.Conn.Exec("DELETE FROM tags WHERE tag_id = ?", id)
	return id, err
}
