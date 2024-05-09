package base

import (
	"context"
	"github.com/kebin6/bsuser-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/bsuser-rpc/internal/svc"
	"github.com/kebin6/bsuser-rpc/types/bsuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *bsuser.BsUserInfo) (*bsuser.BaseIDResp, error) {
	query := l.svcCtx.DB.Bsuser.UpdateOneID(*in.Id)

	if in.Status != nil && *in.Status != 0 {
		query.SetNotNilStatus(pointy.GetStatusPointer(in.Status))
	}

	if in.Pwd != nil && *in.Pwd != "" {
		query.SetNotNilPwd(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Pwd)))
	}

	if in.Name != nil && *in.Name != "" {
		query.SetNotNilName(pointy.GetPointer(*in.Name))
	}

	if in.Mobile != nil && *in.Mobile != "" {
		query.SetNotNilMobile(pointy.GetPointer(*in.Mobile))
	}

	if in.TotalAmount != nil {
		query.SetNotNilTotalAmount(pointy.GetPointer(*in.TotalAmount))
	}

	if in.ValidAmount != nil {
		query.SetNotNilValidAmount(pointy.GetPointer(*in.ValidAmount))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &bsuser.BaseIDResp{Msg: i18n.UpdateSuccess}, nil
}
