package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EventReq struct {
	g.Meta        `path:"/backend/event/add" tags:"Event" method:"post" summary:"You first event api"`
	CardName      string `json:"card_name" v:"required#E通卡名称不能为空" dc:"图片链接"`
	SaleStartDate string `json:"sale_start_date"    v:"required#售卖开始时间不能为空" dc:"售卖开始时间"`
	SaleStopDate  string `json:"sale_stop_date"    v:"required#售卖截止时间不能为空" dc:"售卖截止时间"`
	OrgId         int    `json:"org_id"   v:"required#主办id不能为空" dc:"主办id"`
	Usage         int    `json:"usage"   v:"required#使用方式 1、预约 2、免预约不能为空" dc:"使用方式 1、预约 2、免预约"`
}

type EventRes struct {
	EventId int `json:"event_id"`
}

type EventDeleteReq struct {
	g.Meta `path:"/backend/event/delete" method:"delete" tags:"event" summary:"删除event接口"`
	Id     uint `v:"min:1#请选择需要删除的event" dc:"eventid"`
}
type EventDeleteRes struct{}

type EventUpdateReq struct {
	g.Meta        `path:"/backend/event/update/{Id}" method:"post" tags:"event" summary:"修改event接口"`
	Id            uint   `json:"id"      v:"min:1#请选择需要修改的event" dc:"eventId"`
	CardName      string `json:"card_name" v:"required#E通卡名称不能为空" dc:"图片链接"`
	SaleStartDate string `json:"sale_start_date"    v:"required#售卖开始时间不能为空" dc:"售卖开始时间"`
	SaleStopDate  string `json:"sale_stop_date"    v:"required#售卖截止时间不能为空" dc:"售卖截止时间"`
	OrgId         int    `json:"org_id"   v:"required#主办id不能为空" dc:"主办id"`
	Usage         int    `json:"usage"   v:"required#使用方式 1、预约 2、免预约不能为空" dc:"使用方式 1、预约 2、免预约"`
}
type EventUpdateRes struct {
	Id uint `json:"id"`
}
type EventGetListCommonReq struct {
	g.Meta `path:"/backend/event/list" method:"get" tags:"event" summary:"event列表接口"`
	CommonPaginationReq
}
type EventGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
