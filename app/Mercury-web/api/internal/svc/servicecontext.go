package svc

import (
	"Mercury/app/Mercury-web/api/internal/config"
	"Mercury/app/Mercury-web/rpc/mercury"
	"Mercury/app/Mercury-web/rpc/mercuryClient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	SendRpc mercury.MercuryClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		SendRpc: mercuryClient.NewMercury(zrpc.MustNewClient(c.SendRpcConf)),
	}
}
