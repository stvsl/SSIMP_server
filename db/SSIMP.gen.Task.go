package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
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
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Find(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("Employer").Where("employid = ?", result.Employid).Find(&result.Employer).Error; err != nil { // 员工信息表 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// Gets 获取批量结果
func (obj *_TaskMgr) Gets() (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("Employer").Where("employid = ?", results[i].Employid).Find(&results[i].Employer).Error; err != nil { // 员工信息表 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
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

// WithTask task获取 任务
func (obj *_TaskMgr) WithTask(task string) Option {
	return optionFunc(func(o *options) { o.query["task"] = task })
}


// GetByOption 功能选项模式获取
func (obj *_TaskMgr) GetByOption(opts ...Option) (result Task, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("Employer").Where("employid = ?", result.Employid).Find(&result.Employer).Error; err != nil { // 员工信息表 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

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
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("Employer").Where("employid = ?", results[i].Employid).Find(&results[i].Employer).Error; err != nil { // 员工信息表 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
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
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("Employer").Where("employid = ?", results[i].Employid).Find(&results[i].Employer).Error; err != nil { // 员工信息表 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	resultPage.SetRecords(results)
	return
}


//////////////////////////enume case ////////////////////////////////////////////


// GetFromEmployid 通过employid获取内容 员工编号 
func (obj *_TaskMgr)  GetFromEmployid(employid string) (result Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`employid` = ?", employid).Find(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("Employer").Where("employid = ?", result.Employid).Find(&result.Employer).Error; err != nil { // 员工信息表 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}

// GetBatchFromEmployid 批量查找 员工编号
func (obj *_TaskMgr) GetBatchFromEmployid(employids []string) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`employid` IN (?)", employids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("Employer").Where("employid = ?", results[i].Employid).Find(&results[i].Employer).Error; err != nil { // 员工信息表 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}
 
// GetFromTask 通过task获取内容 任务 
func (obj *_TaskMgr) GetFromTask(task string) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`task` = ?", task).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("Employer").Where("employid = ?", results[i].Employid).Find(&results[i].Employer).Error; err != nil { // 员工信息表 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}

// GetBatchFromTask 批量查找 任务
func (obj *_TaskMgr) GetBatchFromTask(tasks []string) (results []*Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`task` IN (?)", tasks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {  
		if err = obj.NewDB().Table("Employer").Where("employid = ?", results[i].Employid).Find(&results[i].Employer).Error; err != nil { // 员工信息表 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}  
	}
}
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_TaskMgr) FetchByPrimaryKey(employid string ) (result Task, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Task{}).Where("`employid` = ?", employid).Find(&result).Error
	if err == nil && obj.isRelated {  
		if err = obj.NewDB().Table("Employer").Where("employid = ?", result.Employid).Find(&result.Employer).Error; err != nil { // 员工信息表 
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			} }

	return
}
 

 

	

