package model

// EventCreateUpdateBase 创建/修改内容基类
type EventCreateUpdateBase struct {
	OrgId         int
	CardName      string
	SaleStartDate string
	SaleStopDate  string
	Usage         int
}

// EventCreateInput 创建内容
type EventCreateInput struct {
	EventCreateUpdateBase
}

// EventCreateOutput 创建内容返回结果
type EventCreateOutput struct {
	EventId int `json:"event_id"`
}

// EventUpdateInput 修改内容
type EventUpdateInput struct {
	EventCreateUpdateBase
	Id uint
}

// EventGetListInput 获取内容列表
type EventGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// EventGetListOutput 查询列表结果
type EventGetListOutput struct {
	List  []EventGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

// EventSearchInput 搜索列表
type EventSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// EventSearchOutput 搜索列表结果
type EventSearchOutput struct {
	List  []EventSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int          `json:"stats"` // 搜索统计
	Page  int                     `json:"page"`  // 分页码
	Size  int                     `json:"size"`  // 分页数量
	Total int                     `json:"total"` // 数据总数
}

type EventGetListOutputItem struct {
	Event  *EventListItem         `json:"event"`
	Extend []*EventListExtendItem `json:"extend"`
	//Id        uint        `json:"id"` // 自增ID
	//PicUrl    string      `json:"pic_url"`
	//Link      string      `json:"link"`
	//Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	//CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	//UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type EventSearchOutputItem struct {
	EventGetListOutputItem
}

// EventListItem 主要用于列表展示
type EventListItem struct {
	Id            uint   `json:"id"` // 自增ID
	OrgId         uint   `json:"org_id"`
	CardName      string `json:"card_name"`
	SaleStartDate string `json:"sale_start_date"`
	SaleStopDate  string `json:"sale_stop_date"`
	Usage         uint   `json:"usage"`
}

type EventListExtendItem struct {
	Id           uint   `json:"id"` // 自增ID
	OrgId        uint   `json:"org_id"`
	CardId       uint   `json:"card_id"`
	EventId      uint   `json:"event_id"`
	ScreeningsId uint   `json:"screenings_id"`
	PriceId      string `json:"price_id"`
}
