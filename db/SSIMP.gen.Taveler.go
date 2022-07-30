package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _TavelerMgr struct {
	*_BaseMgr
}

// TavelerMgr open func
func TavelerMgr(db *gorm.DB) *_TavelerMgr {
	if db == nil {
		panic(fmt.Errorf("TavelerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TavelerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Taveler"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TavelerMgr) GetTableName() string {
	return "Taveler"
}

// Reset 重置gorm会话
func (obj *_TavelerMgr) Reset() *_TavelerMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_TavelerMgr) Get() (result Taveler, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_TavelerMgr) Gets() (results []*Taveler, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TavelerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Taveler{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithWxinfo wxinfo获取 微信登录信息
func (obj *_TavelerMgr) WithWxinfo(wxinfo string) Option {
	return optionFunc(func(o *options) { o.query["wxinfo"] = wxinfo })
}

// WithID id获取 旅客id
func (obj *_TavelerMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPasswd passwd获取 登录密码
func (obj *_TavelerMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithHistory history获取 旅游记录
func (obj *_TavelerMgr) WithHistory(history string) Option {
	return optionFunc(func(o *options) { o.query["history"] = history })
}


// GetByOption 功能选项模式获取
func (obj *_TavelerMgr) GetByOption(opts ...Option) (result Taveler, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where(options.query).Find(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TavelerMgr) GetByOptions(opts ...Option) (results []*Taveler, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_TavelerMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Taveler,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where(options.query)
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


// GetFromWxinfo 通过wxinfo获取内容 微信登录信息 
func (obj *_TavelerMgr) GetFromWxinfo(wxinfo string) (results []*Taveler, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where("`wxinfo` = ?", wxinfo).Find(&results).Error
	
	return
}

// GetBatchFromWxinfo 批量查找 微信登录信息
func (obj *_TavelerMgr) GetBatchFromWxinfo(wxinfos []string) (results []*Taveler, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where("`wxinfo` IN (?)", wxinfos).Find(&results).Error
	
	return
}
 
// GetFromID 通过id获取内容 旅客id 
func (obj *_TavelerMgr) GetFromID(id string) (results []*Taveler, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where("`id` = ?", id).Find(&results).Error
	
	return
}

// GetBatchFromID 批量查找 旅客id
func (obj *_TavelerMgr) GetBatchFromID(ids []string) (results []*Taveler, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where("`id` IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromPasswd 通过passwd获取内容 登录密码 
func (obj *_TavelerMgr) GetFromPasswd(passwd string) (results []*Taveler, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where("`passwd` = ?", passwd).Find(&results).Error
	
	return
}

// GetBatchFromPasswd 批量查找 登录密码
func (obj *_TavelerMgr) GetBatchFromPasswd(passwds []string) (results []*Taveler, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where("`passwd` IN (?)", passwds).Find(&results).Error
	
	return
}
 
// GetFromHistory 通过history获取内容 旅游记录 
func (obj *_TavelerMgr) GetFromHistory(history string) (results []*Taveler, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where("`history` = ?", history).Find(&results).Error
	
	return
}

// GetBatchFromHistory 批量查找 旅游记录
func (obj *_TavelerMgr) GetBatchFromHistory(historys []string) (results []*Taveler, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Taveler{}).Where("`history` IN (?)", historys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

