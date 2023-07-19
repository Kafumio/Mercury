package srv

import (
	"Mercury/app/Mercury-common/types"
	"Mercury/app/Mercury-job/internal/handler/services/deduplication/structs"
	"Mercury/common/zutils/arrayUtils"
	"context"
)

type deduplicationService struct {
}

func (c deduplicationService) Deduplication(ctx context.Context,
	limit structs.LimitService,
	service structs.DeduplicationService,
	taskInfo *types.TaskInfo,
	param structs.DeduplicationConfigItem) error {

	var newRows []string
	filter, err := limit.LimitFilter(ctx, service, taskInfo, param)
	if err != nil {
		return err
	}
	for _, s := range taskInfo.Receiver {
		if !arrayUtils.ArrayStringIn(filter, s) {
			newRows = append(newRows, s)
		}
	}
	taskInfo.Receiver = newRows
	return nil
}
