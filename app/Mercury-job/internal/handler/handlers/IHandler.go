package handlers

import (
	"Mercury/app/Mercury-common/types"
	"context"
)

type ILimit interface {
	Limit(ctx context.Context, taskInfo types.TaskInfo) bool
}

type IHandler interface {
	ILimit
	DoHandler(ctx context.Context, taskInfo types.TaskInfo) (err error)
}
