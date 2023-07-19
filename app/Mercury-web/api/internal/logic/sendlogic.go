package logic

import (
	"Mercury/app/Mercury-web/rpc/mercury"
	"Mercury/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Mercury/app/Mercury-web/api/internal/svc"
	"Mercury/app/Mercury-web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendLogic) Send(req types.SendRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	variables, err := jsonx.Marshal(req.MessageParam.Variables)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("客户端参数错误"), "send err:%v", err)
	}

	extra, err := jsonx.Marshal(req.MessageParam.Extra)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("客户端参数错误"), "send err:%v", err)
	}

	send, err := l.svcCtx.SendRpc.Send(l.ctx, &mercury.SendRequest{
		Code:              req.Code,
		MessageTemplateId: req.MessageTemplateId,
		MessageParam: &mercury.MessageParam{
			Receiver:  req.MessageParam.Receiver,
			Variables: string(variables),
			Extra:     string(extra),
		},
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("发送消息错误"), "send err:%v", err)
	}
	return &types.Response{Message: send.Code}, err
}
