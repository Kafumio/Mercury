package logic

import (
	"Mercury/app/Mercury-common/interfaces"
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-web/rpc/internal/process"
	"Mercury/app/Mercury-web/rpc/internal/process/action"
	"Mercury/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Mercury/app/Mercury-web/rpc/internal/svc"
	"Mercury/app/Mercury-web/rpc/mercury"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchSendLogic {
	return &BatchSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchSendLogic) BatchSend(in *mercury.BatchSendRequest) (resp *mercury.SendResponse, err error) {
	// todo: add your logic here and delete this line
	if in.MessageParam == nil || len(in.MessageParam) <= 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("客户端参数错误"), "in:%v", in)
	}
	var messageParamList = make([]types.MessageParam, 0)
	for _, item := range in.MessageParam {
		variables := make(map[string]interface{})
		extra := make(map[string]interface{})
		err = jsonx.Unmarshal([]byte(item.Variables), &variables)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("客户端参数错误"), "in:%v err:%v", in, err)
		}
		err = jsonx.Unmarshal([]byte(item.Extra), &extra)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("客户端参数错误"), "in:%v err:%v", in, err)
		}
		messageParamList = append(messageParamList, types.MessageParam{
			Receiver:  item.Receiver,
			Variables: variables,
			Extra:     extra,
		})
	}
	var sendTaskModel = &types.SendTaskModel{
		MessageTemplateId: in.MessageTemplateId,
		MessageParamList:  messageParamList,
	}
	businessProcess := process.NewBusinessProcess()
	_ = businessProcess.AddProcess([]interfaces.Process{
		action.NewPreParamCheckAction(),    //前置参数校验
		action.NewAssembleAction(l.svcCtx), //拼装参数
		action.NewAfterParamCheckAction(),  //后置参数检查
		action.NewSendMqAction(l.svcCtx),   //发送到mq
	}...)

	err = businessProcess.Process(l.ctx, sendTaskModel)
	if err != nil {
		return nil, err
	}
	return &mercury.SendResponse{Code: in.Code}, nil
}
