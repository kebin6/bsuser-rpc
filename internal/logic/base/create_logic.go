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

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *bsuser.BsUserInfo) (*bsuser.BaseIDResp, error) {
	query := l.svcCtx.DB.Bsuser.Create().
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetName(*in.Name).
		SetMobile(*in.Mobile).
		SetInviteCode(*in.InviteCode).
		SetInvitedBy(*in.InvitedBy)

	if in.Pwd == nil {
		in.Pwd = pointy.GetPointer("123456")
	}
	query.SetPwd(encrypt.BcryptEncrypt(*in.Pwd))
	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &bsuser.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
