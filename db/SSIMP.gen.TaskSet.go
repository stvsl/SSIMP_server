package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TaskSetMgr struct {
	*_BaseMgr
}

// TaskSetMgr open func
func TaskSetMgr(db *gorm.DB) *_TaskSetMgr {
	if db == nil {
		panic(fmt.Errorf("TaskSetMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TaskSetMgr{_BaseMgr: &_BaseMgr{DB: db.Table("TaskSet"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TaskSetMgr) GetTableName() string {
	return "TaskSet"
}

// Reset 重置gorm会话
func (obj *_TaskSetMgr) Reset() *_TaskSetMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TaskSetMgr) Get() (result TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TaskSetMgr) Gets() (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TaskSetMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTaskSetID taskSetID获取 任务集ID
func (obj *_TaskSetMgr) WithTid(tid int) Option {
	return optionFunc(func(o *options) { o.query["tid"] = tid })
}

// WithName name获取 任务名称
func (obj *_TaskSetMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithContent content获取 任务内容
func (obj *_TaskSetMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithArea area获取 区域
func (obj *_TaskSetMgr) WithArea(area string) Option {
	return optionFunc(func(o *options) { o.query["area"] = area })
}

// WithPoslo poslo获取 经度
func (obj *_TaskSetMgr) WithPoslo(poslo float64) Option {
	return optionFunc(func(o *options) { o.query["poslo"] = poslo })
}

// WithPosli posli获取 纬度
func (obj *_TaskSetMgr) WithPosli(posli float64) Option {
	return optionFunc(func(o *options) { o.query["posli"] = posli })
}

// WithCycle cycle获取 周期(每周完成多少次)
func (obj *_TaskSetMgr) WithCycle(cycle int) Option {
	return optionFunc(func(o *options) { o.query["cycle"] = cycle })
}

// GetByOption 功能选项模式获取
func (obj *_TaskSetMgr) GetByOption(opts ...Option) (result TaskSet, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TaskSetMgr) GetByOptions(opts ...Option) (results []*TaskSet, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TaskSetMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TaskSet, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

// ////////////////////////enume case ////////////////////////////////////////////

// GetFromTaskSetTid 通过taskSetTid获取内容 任务集ID
func (obj *_TaskSetMgr) GetFromTid(tid int) (result TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`tid` = ?", tid).Find(&result).Error

	return
}

// GetBatchFromTaskSetTid 批量查找 任务集ID
func (obj *_TaskSetMgr) GetBatchFromTid(tids []int) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`tid` IN (?)", tids).Find(&results).Error

	return
}

// GetFromTaskSetName 通过taskSetName获取内容 任务名称
func (obj *_TaskSetMgr) GetFromName(name string) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromTaskSetName 批量查找 任务名称
func (obj *_TaskSetMgr) GetBatchFromName(names []string) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromTaskSetContent 通过taskSetContent获取内容 任务内容
func (obj *_TaskSetMgr) GetFromContent(content string) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromTaskSetContent 批量查找 任务内容
func (obj *_TaskSetMgr) GetBatchFromContent(contents []string) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromTaskSetArea 通过taskSetArea获取内容 区域
func (obj *_TaskSetMgr) GetFromArea(area string) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`area` = ?", area).Find(&results).Error

	return
}

// GetBatchFromTaskSetArea 批量查找 区域
func (obj *_TaskSetMgr) GetBatchFromArea(areas []string) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`area` IN (?)", areas).Find(&results).Error

	return
}

// GetFromTaskSetPoslo 通过taskSetPoslo获取内容 经度
func (obj *_TaskSetMgr) GetFromPoslo(poslo float64) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`poslo` = ?", poslo).Find(&results).Error

	return
}

// GetBatchFromTaskSetPoslo 批量查找 经度
func (obj *_TaskSetMgr) GetBatchFromPoslo(poslos []float64) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`poslo` IN (?)", poslos).Find(&results).Error

	return
}

// GetFromTaskSetPosli 通过taskSetPosli获取内容 纬度
func (obj *_TaskSetMgr) GetFromPosli(posli float64) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`posli` = ?", posli).Find(&results).Error

	return
}

// GetBatchFromTaskSetPosli 批量查找 纬度
func (obj *_TaskSetMgr) GetBatchFromPosli(poslis []float64) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`posli` IN (?)", poslis).Find(&results).Error

	return
}

// GetFromTaskSetCycle 通过taskSetCycle获取内容 周期(每周完成多少次)
func (obj *_TaskSetMgr) GetFromCycle(cycle int) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`cycle` = ?", cycle).Find(&results).Error

	return
}

// GetBatchFromTaskSetCycle 批量查找 周期(每周完成多少次)
func (obj *_TaskSetMgr) GetBatchFromCycle(cycles []int) (results []*TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`cycle` IN (?)", cycles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TaskSetMgr) FetchByPrimaryKey(TaskSetID string) (result TaskSet, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TaskSet{}).Where("`TaskSetID` = ?", TaskSetID).Find(&result).Error

	return
}
