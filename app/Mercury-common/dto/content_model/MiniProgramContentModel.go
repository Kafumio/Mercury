package content_model

import (
	"Mercury/app/Mercury-common/model"
	"Mercury/app/Mercury-common/types"
)

type MiniProgramContentModel struct {
}

func NewMiniProgramContentModel() *MiniProgramContentModel {
	return &MiniProgramContentModel{}
}

func (m MiniProgramContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	return m
}
