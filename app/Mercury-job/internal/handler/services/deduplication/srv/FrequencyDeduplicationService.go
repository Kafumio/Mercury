package srv

import (
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-job/internal/handler/services/deduplication/structs"
	"Mercury/app/Mercury-job/internal/svc"
	"context"
	"fmt"
)

const frequencyDeduplicationServicePrefix = "FRE"

type frequencyDeduplicationService struct {
	svcCtx *svc.ServiceContext
	limit  structs.LimitService

	deduplicationService
}

func NewFrequencyDeduplicationService(svcCtx *svc.ServiceContext, limit structs.LimitService) *frequencyDeduplicationService {
	return &frequencyDeduplicationService{svcCtx: svcCtx, limit: limit}
}

func (c frequencyDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) error {
	return c.deduplicationService.Deduplication(ctx, c.limit, c, taskInfo, param)
}

func (c frequencyDeduplicationService) DeduplicationSingleKey(taskInfo *types.TaskInfo, receiver string) string {
	return fmt.Sprintf("%s_%s_%d_%d", frequencyDeduplicationServicePrefix, receiver, taskInfo.MessageTemplateId, taskInfo.SendChannel)
}
