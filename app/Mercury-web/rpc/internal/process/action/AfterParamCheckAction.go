package action

import (
	"Mercury/app/Mercury-common/enums/channelType"
	"Mercury/app/Mercury-common/enums/idType"
	"Mercury/app/Mercury-common/taskUtil"
	"Mercury/app/Mercury-common/types"
	"context"
	"github.com/pkg/errors"
	"regexp"
)

type AfterParamCheckAction struct {
}

func NewAfterParamCheckAction() *AfterParamCheckAction {
	return &AfterParamCheckAction{}
}

func (p AfterParamCheckAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel) error {

	// 1. 过滤掉不合法的手机号

	if sendTaskModel.TaskInfo[0].IdType == idType.Phone && sendTaskModel.TaskInfo[0].SendChannel == channelType.Sms {
		var newTask []types.TaskInfo
		for _, item := range sendTaskModel.TaskInfo {
			for _, tel := range item.Receiver {
				matched, _ := regexp.Match(taskUtil.PhoneRegex, []byte(tel))
				if matched {
					newTask = append(newTask, item)
				}
			}
		}
		if len(newTask) <= 0 {
			return errors.Wrapf(sendErr, "AfterParamCheckAction sendTaskModel:%v", sendTaskModel)
		}
		sendTaskModel.TaskInfo = newTask
	}

	return nil
}
