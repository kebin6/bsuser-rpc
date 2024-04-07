package base

import (
	"context"
	bsUserEnt "github.com/kebin6/bsuser-rpc/ent/bsuser"
	"github.com/kebin6/bsuser-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/bsuser-rpc/internal/svc"
	"github.com/kebin6/bsuser-rpc/types/bsuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByMobileLogic {
	return &GetByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByMobileLogic) GetByMobile(in *bsuser.MobileReq) (*bsuser.BsUserInfo, error) {
	info, err := l.svcCtx.DB.Bsuser.Query().Where(bsUserEnt.MobileEQ(in.Mobile)).First(l.ctx)
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
