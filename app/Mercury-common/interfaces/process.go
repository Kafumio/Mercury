package interfaces

import (
	"Mercury/app/Mercury-common/types"
	"context"
)

type Process interface {
	Process(ctx context.Context, sendTaskModel *types.SendTaskModel) error
}
