package action

import (
	"Mercury/app/Mercury-common/enums/channelType"
	"Mercury/app/Mercury-common/enums/messageType"
	"Mercury/app/Mercury-common/taskUtil"
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-web/rpc/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type SendMqAction struct {
	svcCtx *svc.ServiceContext
}

func NewSendMqAction(svcCtx *svc.ServiceContext) *SendMqAction {
	return &SendMqAction{svcCtx: svcCtx}
}

func (p SendMqAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel) error {
	marshal, err := jsonx.Marshal(sendTaskModel.TaskInfo)
	if err != nil {
		return err
	}
	channel := channelType.TypeCodeEn[sendTaskModel.TaskInfo[0].SendChannel]
	msgType := messageType.TypeCodeEn[sendTaskModel.TaskInfo[0].MsgType]
	return p.svcCtx.MqClient.Publish(marshal, taskUtil.GetMqKey(channel, msgType))
}
