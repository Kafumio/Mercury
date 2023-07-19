package deduplicationService

import (
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-job/internal/handler/services/deduplication/limit"
	"Mercury/app/Mercury-job/internal/handler/services/deduplication/srv"
	"Mercury/app/Mercury-job/internal/handler/services/deduplication/structs"
	"Mercury/app/Mercury-job/internal/svc"
	"context"
)

type contentDeduplicationService struct {
	svcCtx *svc.ServiceContext
}

func NewContentDeduplicationService(svcCtx *svc.ServiceContext) structs.DuplicationService {
	return &contentDeduplicationService{svcCtx: svcCtx}
}

func (c contentDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) error {
	return srv.NewContentDeduplicationService(c.svcCtx, limit.NewSimpleLimitService(c.svcCtx)).
		Deduplication(ctx, taskInfo, param)
}
