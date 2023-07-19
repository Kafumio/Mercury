package content_model

import (
	"Mercury/app/Mercury-common/model"
	"Mercury/app/Mercury-common/types"
)

type PushContentModel struct {
}

func NewPushContentModel() *PushContentModel {
	return &PushContentModel{}
}

func (d PushContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	return d
}
