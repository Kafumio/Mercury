package services

import (
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-job/internal/svc"
	"Mercury/common/zutils/arrayUtils"
	"Mercury/common/zutils/transform"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type discardMessageService struct {
	svcCtx *svc.ServiceContext
}

const discardMessageKey = "discard_message"

func NewDiscardMessageService(svcCtx *svc.ServiceContext) *discardMessageService {
	return &discardMessageService{svcCtx: svcCtx}
}

// IsDiscard 根据redis配置丢弃某个模板的所有消息
func (l discardMessageService) IsDiscard(ctx context.Context, taskInfo *types.TaskInfo) bool {
	discardMessageTemplateIds, err := l.svcCtx.RedisClient.SmembersCtx(ctx, discardMessageKey)
	if err != nil {
		logx.Errorw("discardMessageService smembers ", logx.Field("err", err))
		return false
	}
	if len(discardMessageTemplateIds) == 0 {
		return false
	}
	if arrayUtils.ArrayInt64In(transform.ArrayStringToInt64(discardMessageTemplateIds), taskInfo.MessageTemplateId) {
		return true
	}
	return false
}
