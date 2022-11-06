package model

//Tag is TagModel
type Tag struct {
	TagId   string `json:"tagId"`
	TagName string `json:"tagName"`
}

//NewTag create a new tag
func NewTag(tagId string, tagName string) (*Tag, error) {
	t := &Tag{
		TagId:   tagId,
		TagName: tagName,
	}
	err := t.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return t, nil
}

//Validate validate tag
func (t *Tag) Validate() error {
	if t.TagName == "" {
		return ErrInvalidEntity
	}
	return nil
}
