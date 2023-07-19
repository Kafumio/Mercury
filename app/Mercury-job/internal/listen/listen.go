package listen

import (
	"Mercury/app/Mercury-job/internal/cron"
	"Mercury/app/Mercury-job/internal/mqs"
	"Mercury/app/Mercury-job/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/service"
)

// Mqs 返回所有消费者
func Mqs(svcCtx *svc.ServiceContext) []service.Service {
	ctx := context.Background()

	var services []service.Service
	services = append(services, []service.Service{
		mqs.NewRabbitTask(ctx, svcCtx),
		cron.NewCornTask(ctx, svcCtx),
	}...)

	return services
}
