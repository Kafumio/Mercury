package structs

import (
	"Mercury/app/Mercury-common/types"
	"context"
)

type DeduplicationConfigItem struct {
	Num  int   `json:"num"`  // 次数 当配置为 10(内容去重) 表示N秒内达到几次会消息被丢弃 20(一天内N次相同渠道去重) 一天内N次相同idType+渠道去重
	Time int64 `json:"time"` // 时间 当配置为 10(内容去重) 表示N秒内内容重复的消息会直接丢弃 20(一天内N次相同渠道去重) 无效
}

type DuplicationService interface {
	Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param DeduplicationConfigItem) error
}
