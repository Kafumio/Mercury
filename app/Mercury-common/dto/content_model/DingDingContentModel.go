package content_model

import (
	"Mercury/app/Mercury-common/model"
	"Mercury/app/Mercury-common/taskUtil"
	"Mercury/app/Mercury-common/types"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type DingDingContentModel struct {
	// SendType string `json:"sendType"`
	Content string `json:"content"`
	// MediaId  string `json:"mediaId"`
}

func NewDingDingContentModel() *DingDingContentModel {
	return &DingDingContentModel{}
}

func (d DingDingContentModel) BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables
	var content DingDingContentModel
	_ = jsonx.Unmarshal([]byte(messageTemplate.MsgContent), &content)
	newVariables := getStringVariables(variables)
	content.Content = taskUtil.ReplaceByMap(content.Content, newVariables)
	// content.SendType = newVariables["sendType"]
	// content.MediaId = newVariables["mediaId"]
	return content
}
