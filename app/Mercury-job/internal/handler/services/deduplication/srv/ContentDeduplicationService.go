package srv

import (
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-job/internal/handler/services/deduplication/structs"
	"Mercury/app/Mercury-job/internal/svc"
	"Mercury/common/zutils/encrypt"
	"context"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type contentDeduplicationService struct {
	svcCtx *svc.ServiceContext
	limit  structs.LimitService

	deduplicationService
}

func NewContentDeduplicationService(svcCtx *svc.ServiceContext, limit structs.LimitService) structs.DeduplicationService {
	return &contentDeduplicationService{svcCtx: svcCtx, limit: limit}
}

func (c contentDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) error {
	return c.deduplicationService.Deduplication(ctx, c.limit, c, taskInfo, param)
}

func (c contentDeduplicationService) DeduplicationSingleKey(taskInfo *types.TaskInfo, receiver string) string {
	str, _ := jsonx.Marshal(taskInfo.ContentModel)
	return encrypt.MD5(cast.ToString(taskInfo.MessageTemplateId) + receiver + string(str))
}
