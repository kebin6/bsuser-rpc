package base

import (
	"context"
	bsUserEnt "github.com/kebin6/bsuser-rpc/ent/bsuser"
	"github.com/kebin6/bsuser-rpc/ent/predicate"
	"github.com/kebin6/bsuser-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/bsuser-rpc/internal/svc"
	"github.com/kebin6/bsuser-rpc/types/bsuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListLogic) GetList(in *bsuser.BsUserListReq) (*bsuser.BsUserListResp, error) {
	var predicates []predicate.Bsuser
	if in.Name != nil {
		predicates = append(predicates, bsUserEnt.NameContains(*in.Name))
	}
	if in.Mobile != nil {
		predicates = append(predicates, bsUserEnt.MobileContains(*in.Mobile))
	}
	if in.Status != nil {
		predicates = append(predicates, bsUserEnt.StatusEQ(*pointy.GetStatusPointer(in.Status)))
	}
	if in.InvitedBy != nil {
		predicates = append(predicates, bsUserEnt.InvitedBy(*in.InvitedBy))
	}
	if in.InviteCode != nil {
		predicates = append(predicates, bsUserEnt.InviteCodeEQ(*in.InviteCode))
	}
	if in.MinTotalAmount != nil {
		predicates = append(predicates, bsUserEnt.TotalAmountGTE(*in.MinTotalAmount))
	}
	if in.MaxTotalAmount != nil {
		predicates = append(predicates, bsUserEnt.TotalAmountLTE(*in.MaxTotalAmount))
	}
	if in.MinValidAmount != nil {
		predicates = append(predicates, bsUserEnt.ValidAmountGTE(*in.MinValidAmount))
	}
	if in.MaxValidAmount != nil {
		predicates = append(predicates, bsUserEnt.ValidAmountLTE(*in.MaxValidAmount))
	}
	if in.StartTime != nil {
		predicates = append(predicates, bsUserEnt.CreatedAtGTE(*pointy.GetTimeMilliPointer(in.StartTime)))
	}
	if in.EndTime != nil {
		predicates = append(predicates, bsUserEnt.CreatedAtLTE(*pointy.GetTimeMilliPointer(in.EndTime)))
	}

	result, err := l.svcCtx.DB.Bsuser.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &bsuser.BsUserListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &bsuser.BsUserInfo{
			Id:          pointy.GetPointer(v.ID),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:      pointy.GetPointer(uint32(v.Status)),
			Name:        &v.Name,
			Mobile:      &v.Mobile,
			TotalAmount: &v.TotalAmount,
			ValidAmount: &v.ValidAmount,
			InviteCode:  &v.InviteCode,
			InvitedBy:   &v.InvitedBy,
		})
	}

	return resp, nil
}
