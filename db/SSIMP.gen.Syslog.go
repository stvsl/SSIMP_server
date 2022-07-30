package	model	
import (	
"time"	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _SyslogMgr struct {
	*_BaseMgr
}

// SyslogMgr open func
func SyslogMgr(db *gorm.DB) *_SyslogMgr {
	if db == nil {
		panic(fmt.Errorf("SyslogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SyslogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Syslog"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SyslogMgr) GetTableName() string {
	return "Syslog"
}

// Reset 重置gorm会话
func (obj *_SyslogMgr) Reset() *_SyslogMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_SyslogMgr) Get() (result Syslog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SyslogMgr) Gets() (results []*Syslog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SyslogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Syslog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithInfo info获取 日志信息
func (obj *_SyslogMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithTime time获取 时间
func (obj *_SyslogMgr) WithTime(time time.Time) Option {
	return optionFunc(func(o *options) { o.query["time"] = time })
}

// WithText text获取 日志内容
func (obj *_SyslogMgr) WithText(text string) Option {
	return optionFunc(func(o *options) { o.query["text"] = text })
}

// WithExtra extra获取 附加信息
func (obj *_SyslogMgr) WithExtra(extra string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}


// GetByOption 功能选项模式获取
func (obj *_SyslogMgr) GetByOption(opts ...Option) (result Syslog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where(options.query).Find(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SyslogMgr) GetByOptions(opts ...Option) (results []*Syslog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_SyslogMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Syslog,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where(options.query)
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


// GetFromInfo 通过info获取内容 日志信息 
func (obj *_SyslogMgr) GetFromInfo(info string) (results []*Syslog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where("`info` = ?", info).Find(&results).Error
	
	return
}

// GetBatchFromInfo 批量查找 日志信息
func (obj *_SyslogMgr) GetBatchFromInfo(infos []string) (results []*Syslog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where("`info` IN (?)", infos).Find(&results).Error
	
	return
}
 
// GetFromTime 通过time获取内容 时间 
func (obj *_SyslogMgr) GetFromTime(time time.Time) (results []*Syslog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where("`time` = ?", time).Find(&results).Error
	
	return
}

// GetBatchFromTime 批量查找 时间
func (obj *_SyslogMgr) GetBatchFromTime(times []time.Time) (results []*Syslog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where("`time` IN (?)", times).Find(&results).Error
	
	return
}
 
// GetFromText 通过text获取内容 日志内容 
func (obj *_SyslogMgr) GetFromText(text string) (results []*Syslog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where("`text` = ?", text).Find(&results).Error
	
	return
}

// GetBatchFromText 批量查找 日志内容
func (obj *_SyslogMgr) GetBatchFromText(texts []string) (results []*Syslog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where("`text` IN (?)", texts).Find(&results).Error
	
	return
}
 
// GetFromExtra 通过extra获取内容 附加信息 
func (obj *_SyslogMgr) GetFromExtra(extra string) (results []*Syslog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where("`extra` = ?", extra).Find(&results).Error
	
	return
}

// GetBatchFromExtra 批量查找 附加信息
func (obj *_SyslogMgr) GetBatchFromExtra(extras []string) (results []*Syslog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Syslog{}).Where("`extra` IN (?)", extras).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

