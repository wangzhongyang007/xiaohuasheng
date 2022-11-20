package controller

import (
	"context"
	v1 "xiaohuasheng/api/v1"
	"xiaohuasheng/internal/model"
	"xiaohuasheng/internal/service"
)

// Event 内容管理
var Event = cEvent{}

type cEvent struct{}

func (a *cEvent) Create(ctx context.Context, req *v1.EventReq) (res *v1.EventRes, err error) {
	out, err := service.Event().Create(ctx, model.EventCreateInput{
		EventCreateUpdateBase: model.EventCreateUpdateBase{
			OrgId:         req.OrgId,
			CardName:      req.CardName,
			SaleStartDate: req.SaleStartDate,
			SaleStopDate:  req.SaleStopDate,
			Usage:         req.Usage,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.EventRes{EventId: out.EventId}, nil
}

func (a *cEvent) Delete(ctx context.Context, req *v1.EventDeleteReq) (res *v1.EventDeleteRes, err error) {
	err = service.Event().Delete(ctx, req.Id)
	return
}

func (a *cEvent) Update(ctx context.Context, req *v1.EventUpdateReq) (res *v1.EventUpdateRes, err error) {
	err = service.Event().Update(ctx, model.EventUpdateInput{
		Id: req.Id,
		EventCreateUpdateBase: model.EventCreateUpdateBase{
			OrgId:         req.OrgId,
			CardName:      req.CardName,
			SaleStartDate: req.SaleStartDate,
			SaleStopDate:  req.SaleStopDate,
			Usage:         req.Usage,
		},
	})
	return &v1.EventUpdateRes{Id: req.Id}, nil
}

func (a *cEvent) List(ctx context.Context, req *v1.EventGetListCommonReq) (res *v1.EventGetListCommonRes, err error) {
	getListRes, err := service.Event().GetList(ctx, model.EventGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &v1.EventGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
