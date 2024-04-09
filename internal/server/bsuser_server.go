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

func (s *BsuserServer) Create(ctx context.Context, in *bsuser.BsUserInfo) (*bsuser.BaseIDResp, error) {
	l := base.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *BsuserServer) Update(ctx context.Context, in *bsuser.BsUserInfo) (*bsuser.BaseIDResp, error) {
	l := base.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *BsuserServer) GetById(ctx context.Context, in *bsuser.IDReq) (*bsuser.BsUserInfo, error) {
	l := base.NewGetByIdLogic(ctx, s.svcCtx)
	return l.GetById(in)
}

func (s *BsuserServer) GetByMobile(ctx context.Context, in *bsuser.MobileReq) (*bsuser.BsUserInfo, error) {
	l := base.NewGetByMobileLogic(ctx, s.svcCtx)
	return l.GetByMobile(in)
}

func (s *BsuserServer) GetOne(ctx context.Context, in *bsuser.BsUserInfo) (*bsuser.BsUserInfo, error) {
	l := base.NewGetOneLogic(ctx, s.svcCtx)
	return l.GetOne(in)
}

func (s *BsuserServer) GetList(ctx context.Context, in *bsuser.BsUserListReq) (*bsuser.BsUserListResp, error) {
	l := base.NewGetListLogic(ctx, s.svcCtx)
	return l.GetList(in)
}
