package handlers

import (
	"Mercury/app/Mercury-common/dto/content_model"
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-job/internal/script"
	"context"
	"github.com/pkg/errors"
)

type smsHandler struct {
	BaseHandler
}

func NewSmsHandler() IHandler {
	return &smsHandler{}
}

func (h smsHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	err = script.NewTencentSms().Send(ctx, script.SmsParams{
		MessageTemplateId: taskInfo.MessageTemplateId,
		Phones:            taskInfo.Receiver,
		Content:           getContent(taskInfo),
		SendAccount:       taskInfo.SendAccount,
	})
	if err != nil {
		return errors.Wrap(err, "smsHandler send err")
	}
	return nil
}

func getContent(taskInfo types.TaskInfo) string {
	var content content_model.SmsContentModel
	getContentModel(taskInfo.ContentModel, &content)
	if content.Url != "" {
		return content.Content + " " + content.Url
	}
	return content.Content
}
