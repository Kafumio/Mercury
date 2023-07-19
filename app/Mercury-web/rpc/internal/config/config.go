package config

import (
	"Mercury/common/dbx"
	"Mercury/common/mq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      dbx.DbConfig
	Rabbit     mq.RabbitConf
	CacheRedis cache.CacheConf
}
