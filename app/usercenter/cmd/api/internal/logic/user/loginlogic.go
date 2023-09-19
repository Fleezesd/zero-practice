package user

import (
	"context"

	"github.com/fleezesd/zero-practice/app/usercenter/cmd/api/internal/svc"
	"github.com/fleezesd/zero-practice/app/usercenter/cmd/api/internal/types"
	"github.com/fleezesd/zero-practice/app/usercenter/cmd/rpc/usercenter"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	// 调用rpc服务
	loginResp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &usercenter.LoginReq{
		AuthType: "",
		AuthKey:  req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var resp types.LoginResp
	_ = copier.Copy(&resp, loginResp)
	return &resp, nil
}
