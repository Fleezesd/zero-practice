package svc

import (
	"github.com/fleezesd/zero-practice/app/usercenter/cmd/api/internal/config"
	"github.com/fleezesd/zero-practice/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// rpc 逻辑stub嵌入
	UsercenterRpc usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUserCenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
