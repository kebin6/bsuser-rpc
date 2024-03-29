// Code generated by goctl. DO NOT EDIT.
// Source: bsuser.proto

package server

import (
	"context"

	"github.com/kebin6/bsuser-rpc/internal/logic/base"
	"github.com/kebin6/bsuser-rpc/internal/svc"
	"github.com/kebin6/bsuser-rpc/types/bsuser"
)

type BsuserServer struct {
	svcCtx *svc.ServiceContext
	bsuser.UnimplementedBsuserServer
}

func NewBsuserServer(svcCtx *svc.ServiceContext) *BsuserServer {
	return &BsuserServer{
		svcCtx: svcCtx,
	}
}

func (s *BsuserServer) InitDatabase(ctx context.Context, in *bsuser.Empty) (*bsuser.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}