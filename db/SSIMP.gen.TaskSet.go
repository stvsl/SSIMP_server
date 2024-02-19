package db

import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _TasksetMgr struct {
	*_BaseMgr
}

// TasksetMgr open func
func TasksetMgr(db *gorm.DB) *_TasksetMgr {
	if db == nil {
		panic(fmt.Errorf("TasksetMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TasksetMgr{_BaseMgr: &_BaseMgr{DB: db.Table("TaskSet"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_TasksetMgr) Debug() *_TasksetMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TasksetMgr) GetTableName() string {
	return "TaskSet"
}

// Reset 重置gorm会话
func (obj *_TasksetMgr) Reset() *_TasksetMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_TasksetMgr) Get() (result Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_TasksetMgr) Gets() (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TasksetMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Taskset{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTid tid获取 任务ID
func (obj *_TasksetMgr) WithTid(tid int) Option {
	return optionFunc(func(o *options) { o.query["tid"] = tid })
}

// WithName name获取 任务名称
func (obj *_TasksetMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithContent content获取 任务内容
func (obj *_TasksetMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithArea area获取 区域
func (obj *_TasksetMgr) WithArea(area string) Option {
	return optionFunc(func(o *options) { o.query["area"] = area })
}

// WithPoslo poslo获取 经度
func (obj *_TasksetMgr) WithPoslo(poslo float64) Option {
	return optionFunc(func(o *options) { o.query["poslo"] = poslo })
}

// WithPosli posli获取 纬度
func (obj *_TasksetMgr) WithPosli(posli float64) Option {
	return optionFunc(func(o *options) { o.query["posli"] = posli })
}

// WithCycle cycle获取 周期(每周完成多少次)
func (obj *_TasksetMgr) WithCycle(cycle int) Option {
	return optionFunc(func(o *options) { o.query["cycle"] = cycle })
}

// WithState state获取 任务状态
func (obj *_TasksetMgr) WithState(state int) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithDuration duration获取 任务时长
func (obj *_TasksetMgr) WithDuration(duration int) Option {
	return optionFunc(func(o *options) { o.query["duration"] = duration })
}


// GetByOption 功能选项模式获取
func (obj *_TasksetMgr) GetByOption(opts ...Option) (result Taskset, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TasksetMgr) GetByOptions(opts ...Option) (results []*Taskset, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_TasksetMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Taskset,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	
	resultPage.SetRecords(results)
	return
}


//////////////////////////enume case ////////////////////////////////////////////


// GetFromTid 通过tid获取内容 任务ID 
func (obj *_TasksetMgr)  GetFromTid(tid int) (result Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`tid` = ?", tid).First(&result).Error
	
	return
}

// GetBatchFromTid 批量查找 任务ID
func (obj *_TasksetMgr) GetBatchFromTid(tids []int) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`tid` IN (?)", tids).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容 任务名称 
func (obj *_TasksetMgr) GetFromName(name string) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`name` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 任务名称
func (obj *_TasksetMgr) GetBatchFromName(names []string) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`name` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromContent 通过content获取内容 任务内容 
func (obj *_TasksetMgr) GetFromContent(content string) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`content` = ?", content).Find(&results).Error
	
	return
}

// GetBatchFromContent 批量查找 任务内容
func (obj *_TasksetMgr) GetBatchFromContent(contents []string) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`content` IN (?)", contents).Find(&results).Error
	
	return
}
 
// GetFromArea 通过area获取内容 区域 
func (obj *_TasksetMgr) GetFromArea(area string) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`area` = ?", area).Find(&results).Error
	
	return
}

// GetBatchFromArea 批量查找 区域
func (obj *_TasksetMgr) GetBatchFromArea(areas []string) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`area` IN (?)", areas).Find(&results).Error
	
	return
}
 
// GetFromPoslo 通过poslo获取内容 经度 
func (obj *_TasksetMgr) GetFromPoslo(poslo float64) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`poslo` = ?", poslo).Find(&results).Error
	
	return
}

// GetBatchFromPoslo 批量查找 经度
func (obj *_TasksetMgr) GetBatchFromPoslo(poslos []float64) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`poslo` IN (?)", poslos).Find(&results).Error
	
	return
}
 
// GetFromPosli 通过posli获取内容 纬度 
func (obj *_TasksetMgr) GetFromPosli(posli float64) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`posli` = ?", posli).Find(&results).Error
	
	return
}

// GetBatchFromPosli 批量查找 纬度
func (obj *_TasksetMgr) GetBatchFromPosli(poslis []float64) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`posli` IN (?)", poslis).Find(&results).Error
	
	return
}
 
// GetFromCycle 通过cycle获取内容 周期(每周完成多少次) 
func (obj *_TasksetMgr) GetFromCycle(cycle int) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`cycle` = ?", cycle).Find(&results).Error
	
	return
}

// GetBatchFromCycle 批量查找 周期(每周完成多少次)
func (obj *_TasksetMgr) GetBatchFromCycle(cycles []int) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`cycle` IN (?)", cycles).Find(&results).Error
	
	return
}
 
// GetFromState 通过state获取内容 任务状态 
func (obj *_TasksetMgr) GetFromState(state int) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`state` = ?", state).Find(&results).Error
	
	return
}

// GetBatchFromState 批量查找 任务状态
func (obj *_TasksetMgr) GetBatchFromState(states []int) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`state` IN (?)", states).Find(&results).Error
	
	return
}
 
// GetFromDuration 通过duration获取内容 任务时长 
func (obj *_TasksetMgr) GetFromDuration(duration int) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`duration` = ?", duration).Find(&results).Error
	
	return
}

// GetBatchFromDuration 批量查找 任务时长
func (obj *_TasksetMgr) GetBatchFromDuration(durations []int) (results []*Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`duration` IN (?)", durations).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_TasksetMgr) FetchByPrimaryKey(tid int ) (result Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`tid` = ?", tid).First(&result).Error
	
	return
}
 
 // FetchUniqueByNewtableTidIDx primary or index 获取唯一内容
 func (obj *_TasksetMgr) FetchUniqueByNewtableTidIDx(tid int ) (result Taskset, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taskset{}).Where("`tid` = ?", tid).First(&result).Error
	
	return
}
