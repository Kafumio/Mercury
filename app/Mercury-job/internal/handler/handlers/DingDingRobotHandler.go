package handlers

import (
	"Mercury/app/Mercury-common/dto/account"
	"Mercury/app/Mercury-common/dto/content_model"
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-support/utils/accountUtils"
	"Mercury/common/zutils/arrayUtils"
	"context"
	"github.com/pkg/errors"
	"github.com/wanghuiyt/ding"
)

type dingDingRobotHandler struct {
	BaseHandler
}

const SendAll = "@all"

func NewDingDingRobotHandler() IHandler {
	return dingDingRobotHandler{}
}

func (h dingDingRobotHandler) DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error) {
	var content content_model.DingDingContentModel
	getContentModel(taskInfo.ContentModel, &content)

	var acc account.DingDingRobotAccount
	err = accountUtils.GetAccount(ctx, taskInfo.SendAccount, &acc)
	if err != nil {
		return errors.Wrap(err, "dingDingRobotHandler get account err")
	}
	var at []string
	d := ding.Webhook{
		AccessToken: acc.AccessToken,
		Secret:      acc.Secret,
		EnableAt:    true,
	}

	if arrayUtils.ArrayStringIn(taskInfo.Receiver, SendAll) {
		d.AtAll = true
	} else {
		at = taskInfo.Receiver
	}

	err = d.SendMessage(content.Content, at...)
	if err != nil {
		return errors.Wrap(err, "dingDingRobotHandler SendMessage err")
	}
	return nil
}
