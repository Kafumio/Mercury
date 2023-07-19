package Mercury_job

import (
	"Mercury/app/Mercury-job/internal/config"
	"Mercury/app/Mercury-job/internal/handler/handlers"
	"Mercury/app/Mercury-job/internal/listen"
	"Mercury/app/Mercury-job/internal/svc"
	"Mercury/common/dbx"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/Mercury-job.yaml", "the config file")

func main() {
	flag.Parse()
	logx.DisableStat()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	group := service.NewServiceGroup()
	for _, mq := range listen.Mqs(ctx) {
		group.Add(mq)
	}
	handlers.SetUp(ctx)
	dbx.InitDb(c.Mysql)

	defer group.Stop()
	group.Start()
}
