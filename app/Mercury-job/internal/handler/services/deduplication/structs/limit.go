package structs

import (
	"Mercury/app/Mercury-common/types"
	"context"
)

type LimitService interface {
	LimitFilter(ctx context.Context, duplication DeduplicationService, taskInfo *types.TaskInfo, param DeduplicationConfigItem) (filterReceiver []string, err error)
}
type DeduplicationService interface {
	Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param DeduplicationConfigItem) error
	DeduplicationSingleKey(taskInfo *types.TaskInfo, receiver string) string
}
