package content_model

import (
	"Mercury/app/Mercury-common/model"
	"Mercury/app/Mercury-common/types"
)

type ImContentModel struct {
}

func NewImContentModel() *ImContentModel {
	return &ImContentModel{}
}

func (d ImContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	return d
}
