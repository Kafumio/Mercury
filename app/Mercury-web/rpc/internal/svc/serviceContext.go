package svc

import (
	"Mercury/app/Mercury-web/rpc/internal/config"
	"Mercury/common/mq"
)

type ServiceContext struct {
	Config   config.Config
	MqClient mq.IMessagingClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := mq.NewMessagingClientURL(c.Rabbit.URL)
	handleErr(err)
	return &ServiceContext{
		Config:   c,
		MqClient: client,
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
