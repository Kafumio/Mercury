package action

import (
	"Mercury/app/Mercury-common/dto/content_model"
	"Mercury/app/Mercury-common/repo"
	"Mercury/app/Mercury-common/taskUtil"
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-web/rpc/internal/svc"
	"Mercury/common/zutils/transform"
	"context"
	"github.com/pkg/errors"
	"strings"
)

type AssembleAction struct {
	svcCtx *svc.ServiceContext
}

func NewAssembleAction(svcCtx *svc.ServiceContext) *AssembleAction {
	return &AssembleAction{svcCtx: svcCtx}
}

func (p AssembleAction) Process(ctx context.Context, sendTaskModel *types.SendTaskModel) error {
	messageParamList := sendTaskModel.MessageParamList

	messageTemplate, err := repo.NewMessageTemplateRepo(p.svcCtx.Config.CacheRedis).
		One(ctx, sendTaskModel.MessageTemplateId)
	if err != nil {
		return errors.Wrapf(sendErr, "查询模板异常 err:%v 模板id:%d", err, sendTaskModel.MessageTemplateId)
	}
	contentModel := content_model.GetBuilderContentBySendChannel(messageTemplate.SendChannel)

	var newTaskList []types.TaskInfo
	for _, param := range messageParamList {

		curTask := types.TaskInfo{
			MessageTemplateId: messageTemplate.ID,
			BusinessId:        taskUtil.GenerateBusinessId(messageTemplate.ID, messageTemplate.TemplateType),
			Receiver:          transform.ArrayStringUniq(strings.Split(param.Receiver, ",")),
			IdType:            messageTemplate.IDType,
			SendChannel:       messageTemplate.SendChannel,
			TemplateType:      messageTemplate.TemplateType,
			MsgType:           messageTemplate.MsgType,
			ShieldType:        messageTemplate.ShieldType,
			ContentModel:      contentModel.BuilderContent(messageTemplate, param),
			SendAccount:       messageTemplate.SendAccount,
		}

		newTaskList = append(newTaskList, curTask)
	}
	sendTaskModel.TaskInfo = newTaskList
	return nil
}
