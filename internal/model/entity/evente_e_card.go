// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EventeECard is the golang structure for table evente_e_card.
type EventeECard struct {
	Id            uint        `json:"id"            description:""`
	OrgId         int         `json:"orgId"         description:"主办id"`
	CardName      string      `json:"cardName"      description:"E通卡名称"`
	SaleStartDate *gtime.Time `json:"saleStartDate" description:"售卖开始时间"`
	SaleStopDate  *gtime.Time `json:"saleStopDate"  description:"售卖结束时间"`
	Usage         int         `json:"usage"         description:"使用方式 1、预约 2、免预约"`
}
