package base

import (
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

func (l *InitDatabaseLogic) insertApiData() (err error) {
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Bsuser"),
		Path:        pointy.GetPointer("/bsuser/create"),
		Description: pointy.GetPointer("apiDesc.createBsuser"),
		ApiGroup:    pointy.GetPointer("bsuser"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Bsuser"),
		Path:        pointy.GetPointer("/bsuser/update"),
		Description: pointy.GetPointer("apiDesc.updateBsuser"),
		ApiGroup:    pointy.GetPointer("bsuser"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Bsuser"),
		Path:        pointy.GetPointer("/bsuser/delete"),
		Description: pointy.GetPointer("apiDesc.deleteBsuser"),
		ApiGroup:    pointy.GetPointer("bsuser"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Bsuser"),
		Path:        pointy.GetPointer("/bsuser/list"),
		Description: pointy.GetPointer("apiDesc.getBsuserList"),
		ApiGroup:    pointy.GetPointer("bsuser"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Bsuser"),
		Path:        pointy.GetPointer("/bsuser"),
		Description: pointy.GetPointer("apiDesc.getBsuserById"),
		ApiGroup:    pointy.GetPointer("bsuser"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	return nil
}

func (l *InitDatabaseLogic) insertMenuData() error {
	// 目录
	menuData, err := l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(common.DefaultParentId),
		Name:      pointy.GetPointer("BsuserManagementDirectory"),
		Component: pointy.GetPointer("LAYOUT"),
		Path:      pointy.GetPointer("/bsuser_dir"),
		Sort:      pointy.GetPointer(uint32(1)),
		Disabled:  pointy.GetPointer(false),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("route.bsuserManagement"),
			Icon:               pointy.GetPointer("ic:outline-people-alt"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(0)),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(menuData.Id),
		Path:      pointy.GetPointer("/bsuser"),
		Name:      pointy.GetPointer("BsuserManagement"),
		Component: pointy.GetPointer("/bsuser/bsuser/index"),
		Sort:      pointy.GetPointer(uint32(1)),
		Disabled:  pointy.GetPointer(false),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("route.bsuserManagement"),
			Icon:               pointy.GetPointer("ic:outline-people-alt"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(1)),
	})

	if err != nil {
		return err
	}

	return err
}
