package deduplicationService

import (
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-job/internal/handler/services/deduplication/limit"
	"Mercury/app/Mercury-job/internal/handler/services/deduplication/srv"
	"Mercury/app/Mercury-job/internal/handler/services/deduplication/structs"
	"Mercury/app/Mercury-job/internal/svc"
	"context"
)

type frequencyDeduplicationService struct {
	svcCtx *svc.ServiceContext
}

func NewFrequencyDeduplicationService(svcCtx *svc.ServiceContext) structs.DuplicationService {
	return &frequencyDeduplicationService{svcCtx: svcCtx}
}

func (c frequencyDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) error {
	return srv.NewFrequencyDeduplicationService(c.svcCtx, limit.NewSlideWindowLimitService(c.svcCtx)).
		Deduplication(ctx, taskInfo, param)
}
