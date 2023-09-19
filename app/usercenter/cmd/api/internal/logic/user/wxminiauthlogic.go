package user

import (
	"context"

	"github.com/fleezesd/zero-practice/app/usercenter/cmd/api/internal/svc"
	"github.com/fleezesd/zero-practice/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WXMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWXMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WXMiniAuthLogic {
	return &WXMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WXMiniAuthLogic) WXMiniAuth(req *types.WXMiniAuthReq) (resp *types.WXMiniAuthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
