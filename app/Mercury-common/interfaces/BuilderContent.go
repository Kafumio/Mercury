package interfaces

import (
	"Mercury/app/Mercury-common/model"
	"Mercury/app/Mercury-common/types"
)

type BuilderContent interface {
	BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{}
}
