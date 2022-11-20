package event

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"
	"xiaohuasheng/internal/dao"
	"xiaohuasheng/internal/model"
	"xiaohuasheng/internal/model/entity"
	"xiaohuasheng/internal/service"
)

type sEvent struct{}

func init() {
	service.RegisterEvent(New())
}

func New() *sEvent {
	return &sEvent{}
}

func (s *sEvent) Create(ctx context.Context, in model.EventCreateInput) (out model.EventCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.EventeECard.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.EventCreateOutput{EventId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sEvent) Delete(ctx context.Context, id uint) error {
	return dao.EventeECard.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除内容
		_, err := dao.EventeECard.Ctx(ctx).Where(g.Map{
			dao.EventeECard.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// Update 修改
func (s *sEvent) Update(ctx context.Context, in model.EventUpdateInput) error {
	return dao.EventeECard.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.EventeECard.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.EventeECard.Columns().Id).
			Where(dao.EventeECard.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sEvent) GetList(ctx context.Context, in model.EventGetListInput) (out *model.EventGetListOutput, err error) {
	var (
		m = dao.EventeECard.Ctx(ctx)
	)
	out = &model.EventGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.EventeECard
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Event
	if err := listModel.ScanList(&out.List, "Event"); err != nil {
		return out, err
	}
	// Extend
	err = dao.EventeECardExtend.Ctx(ctx).
		Fields(model.EventListExtendItem{}).
		Where(dao.EventeECardExtend.Columns().CardId, gutil.ListItemValuesUnique(out.List, "Event", "Id")).
		ScanList(&out.List, "Extend", "Event", "card_id:Id")
	if err != nil {
		return out, err
	}
	return
}
