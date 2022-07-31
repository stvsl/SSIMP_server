package db

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _EmployerMgr struct {
	*_BaseMgr
}

// EmployerMgr open func
func EmployerMgr(db *gorm.DB) *_EmployerMgr {
	if db == nil {
		panic(fmt.Errorf("EmployerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EmployerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Employer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EmployerMgr) GetTableName() string {
	return "Employer"
}

// Reset 重置gorm会话
func (obj *_EmployerMgr) Reset() *_EmployerMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EmployerMgr) Get() (result Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EmployerMgr) Gets() (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EmployerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Employer{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithEmployid employid获取 员工编号
func (obj *_EmployerMgr) WithEmployid(employid string) Option {
	return optionFunc(func(o *options) { o.query["employid"] = employid })
}

// WithPasswd passwd获取 登录密码
func (obj *_EmployerMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithName name获取 姓名
func (obj *_EmployerMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithBirthDay birth_day获取 出生日期
func (obj *_EmployerMgr) WithBirthDay(birthDay datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["birth_day"] = birthDay })
}

// WithEmployDay employ_day获取 入职日期
func (obj *_EmployerMgr) WithEmployDay(employDay datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["employ_day"] = employDay })
}

// WithIDcard idcard获取 身份证号
func (obj *_EmployerMgr) WithIDcard(idcard int) Option {
	return optionFunc(func(o *options) { o.query["idcard"] = idcard })
}

// WithAddress address获取 家庭地址
func (obj *_EmployerMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithTelephone telephone获取 联系电话
func (obj *_EmployerMgr) WithTelephone(telephone int) Option {
	return optionFunc(func(o *options) { o.query["telephone"] = telephone })
}

// WithBustPhoto bust_photo获取 半身照
func (obj *_EmployerMgr) WithBustPhoto(bustPhoto []byte) Option {
	return optionFunc(func(o *options) { o.query["bust_photo"] = bustPhoto })
}

// WithAvatar avatar获取 头像
func (obj *_EmployerMgr) WithAvatar(avatar []byte) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// GetByOption 功能选项模式获取
func (obj *_EmployerMgr) GetByOption(opts ...Option) (result Employer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EmployerMgr) GetByOptions(opts ...Option) (results []*Employer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EmployerMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Employer, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Employer{}).Where(options.query)
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
func (obj *_EmployerMgr) GetFromEmployid(employid string) (result Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`employid` = ?", employid).Find(&result).Error

	return
}

// GetBatchFromEmployid 批量查找 员工编号
func (obj *_EmployerMgr) GetBatchFromEmployid(employids []string) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`employid` IN (?)", employids).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容 登录密码
func (obj *_EmployerMgr) GetFromPasswd(passwd string) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`passwd` = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量查找 登录密码
func (obj *_EmployerMgr) GetBatchFromPasswd(passwds []string) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`passwd` IN (?)", passwds).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_EmployerMgr) GetFromName(name string) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 姓名
func (obj *_EmployerMgr) GetBatchFromName(names []string) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromBirthDay 通过birth_day获取内容 出生日期
func (obj *_EmployerMgr) GetFromBirthDay(birthDay datatypes.Date) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`birth_day` = ?", birthDay).Find(&results).Error

	return
}

// GetBatchFromBirthDay 批量查找 出生日期
func (obj *_EmployerMgr) GetBatchFromBirthDay(birthDays []datatypes.Date) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`birth_day` IN (?)", birthDays).Find(&results).Error

	return
}

// GetFromEmployDay 通过employ_day获取内容 入职日期
func (obj *_EmployerMgr) GetFromEmployDay(employDay datatypes.Date) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`employ_day` = ?", employDay).Find(&results).Error

	return
}

// GetBatchFromEmployDay 批量查找 入职日期
func (obj *_EmployerMgr) GetBatchFromEmployDay(employDays []datatypes.Date) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`employ_day` IN (?)", employDays).Find(&results).Error

	return
}

// GetFromIDcard 通过idcard获取内容 身份证号
func (obj *_EmployerMgr) GetFromIDcard(idcard int) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`idcard` = ?", idcard).Find(&results).Error

	return
}

// GetBatchFromIDcard 批量查找 身份证号
func (obj *_EmployerMgr) GetBatchFromIDcard(idcards []int) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`idcard` IN (?)", idcards).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 家庭地址
func (obj *_EmployerMgr) GetFromAddress(address string) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 家庭地址
func (obj *_EmployerMgr) GetBatchFromAddress(addresss []string) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromTelephone 通过telephone获取内容 联系电话
func (obj *_EmployerMgr) GetFromTelephone(telephone int) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`telephone` = ?", telephone).Find(&results).Error

	return
}

// GetBatchFromTelephone 批量查找 联系电话
func (obj *_EmployerMgr) GetBatchFromTelephone(telephones []int) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`telephone` IN (?)", telephones).Find(&results).Error

	return
}

// GetFromBustPhoto 通过bust_photo获取内容 半身照
func (obj *_EmployerMgr) GetFromBustPhoto(bustPhoto []byte) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`bust_photo` = ?", bustPhoto).Find(&results).Error

	return
}

// GetBatchFromBustPhoto 批量查找 半身照
func (obj *_EmployerMgr) GetBatchFromBustPhoto(bustPhotos [][]byte) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`bust_photo` IN (?)", bustPhotos).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像
func (obj *_EmployerMgr) GetFromAvatar(avatar []byte) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找 头像
func (obj *_EmployerMgr) GetBatchFromAvatar(avatars [][]byte) (results []*Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EmployerMgr) FetchByPrimaryKey(employid string) (result Employer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Employer{}).Where("`employid` = ?", employid).Find(&result).Error

	return
}
