package process

import (
	"Mercury/app/Mercury-common/interfaces"
	"Mercury/app/Mercury-common/types"
	"context"
)

type BusinessProcess struct {
	process []interfaces.Process
}

func NewBusinessProcess() *BusinessProcess {
	return &BusinessProcess{
		process: make([]interfaces.Process, 0),
	}
}
func (p *BusinessProcess) Process(ctx context.Context, sendTaskModel *types.SendTaskModel) error {
	for _, pr := range p.process {
		err := pr.Process(ctx, sendTaskModel)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *BusinessProcess) AddProcess(pr ...interfaces.Process) error {
	if len(pr) > 0 {
		p.process = append(p.process, pr...)
	}
	return nil
}
