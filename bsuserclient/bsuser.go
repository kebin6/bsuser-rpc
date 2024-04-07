// Code generated by goctl. DO NOT EDIT.
// Source: bsuser.proto

package bsuserclient

import (
	"context"

	"github.com/kebin6/bsuser-rpc/types/bsuser"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp     = bsuser.BaseIDResp
	BaseResp       = bsuser.BaseResp
	BaseUUIDResp   = bsuser.BaseUUIDResp
	BsUserInfo     = bsuser.BsUserInfo
	BsUserListReq  = bsuser.BsUserListReq
	BsUserListResp = bsuser.BsUserListResp
	Empty          = bsuser.Empty
	IDReq          = bsuser.IDReq
	IDsReq         = bsuser.IDsReq
	MobileReq      = bsuser.MobileReq
	PageInfoReq    = bsuser.PageInfoReq
	UUIDReq        = bsuser.UUIDReq
	UUIDsReq       = bsuser.UUIDsReq

	Bsuser interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		Create(ctx context.Context, in *BsUserInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		Update(ctx context.Context, in *BsUserInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		GetById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BsUserInfo, error)
		GetByMobile(ctx context.Context, in *MobileReq, opts ...grpc.CallOption) (*BsUserInfo, error)
		GetList(ctx context.Context, in *BsUserListReq, opts ...grpc.CallOption) (*BsUserListResp, error)
	}

	defaultBsuser struct {
		cli zrpc.Client
	}
)

func NewBsuser(cli zrpc.Client) Bsuser {
	return &defaultBsuser{
		cli: cli,
	}
}

func (m *defaultBsuser) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := bsuser.NewBsuserClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

func (m *defaultBsuser) Create(ctx context.Context, in *BsUserInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := bsuser.NewBsuserClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultBsuser) Update(ctx context.Context, in *BsUserInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := bsuser.NewBsuserClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (m *defaultBsuser) GetById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BsUserInfo, error) {
	client := bsuser.NewBsuserClient(m.cli.Conn())
	return client.GetById(ctx, in, opts...)
}

func (m *defaultBsuser) GetByMobile(ctx context.Context, in *MobileReq, opts ...grpc.CallOption) (*BsUserInfo, error) {
	client := bsuser.NewBsuserClient(m.cli.Conn())
	return client.GetByMobile(ctx, in, opts...)
}

func (m *defaultBsuser) GetList(ctx context.Context, in *BsUserListReq, opts ...grpc.CallOption) (*BsUserListResp, error) {
	client := bsuser.NewBsuserClient(m.cli.Conn())
	return client.GetList(ctx, in, opts...)
}
