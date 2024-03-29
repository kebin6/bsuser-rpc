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
	BaseIDResp   = bsuser.BaseIDResp
	BaseResp     = bsuser.BaseResp
	BaseUUIDResp = bsuser.BaseUUIDResp
	Empty        = bsuser.Empty
	IDReq        = bsuser.IDReq
	IDsReq       = bsuser.IDsReq
	PageInfoReq  = bsuser.PageInfoReq
	UUIDReq      = bsuser.UUIDReq
	UUIDsReq     = bsuser.UUIDsReq

	Bsuser interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
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