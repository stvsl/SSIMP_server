package db

import (	
"fmt"	
"context"	
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
	return &_TaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Task"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_TaskMgr) Debug() *_TaskMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
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
func (obj *_TaskMgr) Get() (result Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_TaskMgr) Gets() (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Task{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithEmployid employid获取 员工编号
func (obj *_TaskMgr) WithEmployid(employid string) Option {
	return optionFunc(func(o *options) { o.query["employid"] = employid })
}

// WithTid tid获取 任务
func (obj *_TaskMgr) WithTid(tid int) Option {
	return optionFunc(func(o *options) { o.query["tid"] = tid })
}


// GetByOption 功能选项模式获取
func (obj *_TaskMgr) GetByOption(opts ...Option) (result Task, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where(options.query).First(&result).Error
	
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
func (obj *_TaskMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Task,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Task{}).Where(options.query)
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


// GetFromEmployid 通过employid获取内容 员工编号 
func (obj *_TaskMgr) GetFromEmployid(employid string) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`employid` = ?", employid).Find(&results).Error
	
	return
}

// GetBatchFromEmployid 批量查找 员工编号
func (obj *_TaskMgr) GetBatchFromEmployid(employids []string) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`employid` IN (?)", employids).Find(&results).Error
	
	return
}
 
// GetFromTid 通过tid获取内容 任务 
func (obj *_TaskMgr) GetFromTid(tid int) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`tid` = ?", tid).Find(&results).Error
	
	return
}

// GetBatchFromTid 批量查找 任务
func (obj *_TaskMgr) GetBatchFromTid(tids []int) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`tid` IN (?)", tids).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 
 // FetchIndexByTaskFk  获取多个内容
 func (obj *_TaskMgr) FetchIndexByTaskFk(employid string ) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`employid` = ?", employid).Find(&results).Error
	
	return
}
 
 // FetchIndexByTaskFk1  获取多个内容
 func (obj *_TaskMgr) FetchIndexByTaskFk1(tid int ) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`tid` = ?", tid).Find(&results).Error
	
	return
}
