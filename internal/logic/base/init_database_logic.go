package base

import (
	"context"

	"github.com/kebin6/bsuser-rpc/internal/svc"
	"github.com/kebin6/bsuser-rpc/types/bsuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *bsuser.Empty) (*bsuser.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &bsuser.BaseResp{}, nil
}
