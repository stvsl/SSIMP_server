package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TaskMgr struct {
	*_BaseMgr
}

// TaskMgr open func
func TaskMgr(db *gorm.DB) *_TaskMgr {
	if db == nil {
		panic(fmt.Errorf("TaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TaskMgr) GetTableName() string {
	return "Task"
}

// Reset 重置gorm会话
func (obj *_TaskMgr) Reset() *_TaskMgr {
	obj.New()
	return obj
}

// Get 获取
// func (obj *_TaskMgr) Get() (result Task, err error) {

// Gets 获取批量结果
func (obj *_TaskMgr) Gets() (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Task{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithEmployid employid获取 员工编号
func (obj *_TaskMgr) WithEmployid(employid string) Option {
	return optionFunc(func(o *options) { o.query["employid"] = employid })
}

// WithTask task获取 任务
func (obj *_TaskMgr) WithTask(task string) Option {
	return optionFunc(func(o *options) { o.query["task"] = task })
}

// GetByOption 功能选项模式获取
func (obj *_TaskMgr) GetByOption(opts ...Option) (result Task, err error) {
	var results []*Task
	if results, err = obj.GetByOptions(opts...); err != nil {
		return
	}
	if len(results) == 0 {
		err = gorm.ErrRecordNotFound
	} else {
		result = *results[0]
	}
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TaskMgr) GetByOptions(opts ...Option) (results []*Task, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where(options.query).Find(&results).Error
	return
}

// SelectPage 分页查询
func (obj *_TaskMgr) SelectPage(pageSize, pageIndex int, opts ...Option) (total int64, results []*Task, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	db := obj.DB.WithContext(obj.ctx).Model(Task{}).Where(options.query)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&results).Error
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromEmployid 通过employid获取内容 员工编号
func (obj *_TaskMgr) GetFromEmployid(employid string) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("employid = ?", employid).Find(&results).Error
	return
}

// GetBatchFromEmployid 批量唯一主键查找 员工编号
func (obj *_TaskMgr) GetBatchFromEmployid(employids []string) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("employid IN (?)", employids).Find(&results).Error
	return
}
