package model_test

import (
	"local-book-reader/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTag(t *testing.T) {
	tg, err := model.NewTag(model.NewID(), "favorite")
	assert.Nil(t, err)
	assert.NotNil(t, tg.TagId)
	assert.Equal(t, tg.TagName, "favorite")
}

func TestTagValidate(t *testing.T) {
	type test struct {
		tagid   string
		tagname string
		want    error
	}

	tests := []test{
		{
			tagid:   model.NewID(),
			tagname: "favorite",
			want:    nil,
		},
		{
			tagid:   model.NewID(),
			tagname: "",
			want:    model.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := model.NewTag(tc.tagid, tc.tagname)
		assert.Equal(t, err, tc.want)
	}

}
