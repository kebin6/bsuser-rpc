package base

import (
	"context"
	"github.com/kebin6/bsuser-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/bsuser-rpc/internal/svc"
	"github.com/kebin6/bsuser-rpc/types/bsuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdLogic {
	return &GetByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByIdLogic) GetById(in *bsuser.IDReq) (*bsuser.BsUserInfo, error) {
	info, err := l.svcCtx.DB.Bsuser.Get(l.ctx, in.GetId())
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &bsuser.BsUserInfo{
		Id:          pointy.GetPointer(info.ID),
		Name:        pointy.GetPointer(info.Name),
		Mobile:      pointy.GetPointer(info.Mobile),
		Pwd:         pointy.GetPointer(info.Pwd),
		InviteCode:  pointy.GetPointer(info.InviteCode),
		InvitedBy:   pointy.GetPointer(info.InvitedBy),
		TotalAmount: pointy.GetPointer(info.TotalAmount),
		ValidAmount: pointy.GetPointer(info.ValidAmount),
		Status:      pointy.GetPointer(uint32(info.Status)),
		CreatedAt:   pointy.GetPointer(info.CreatedAt.UnixMilli()),
		UpdatedAt:   pointy.GetPointer(info.UpdatedAt.UnixMilli()),
	}, nil
}
