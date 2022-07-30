package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"gorm.io/datatypes"	
)	

type _AdminMgr struct {
	*_BaseMgr
}

// AdminMgr open func
func AdminMgr(db *gorm.DB) *_AdminMgr {
	if db == nil {
		panic(fmt.Errorf("AdminMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdminMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Admin"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AdminMgr) GetTableName() string {
	return "Admin"
}

// Reset 重置gorm会话
func (obj *_AdminMgr) Reset() *_AdminMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_AdminMgr) Get() (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_AdminMgr) Gets() (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AdminMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Admin{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAdminID adminID获取 管理员ID
func (obj *_AdminMgr) WithAdminID(adminID string) Option {
	return optionFunc(func(o *options) { o.query["adminID"] = adminID })
}

// WithEstablishDay establish_day获取 建立日期
func (obj *_AdminMgr) WithEstablishDay(establishDay datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["establish_day"] = establishDay })
}

// WithPasswd passwd获取 管理员密码
func (obj *_AdminMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithName name获取 管理员姓名
func (obj *_AdminMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithBirthday birthday获取 出生年月
func (obj *_AdminMgr) WithBirthday(birthday datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithIDcard idcard获取 身份证号
func (obj *_AdminMgr) WithIDcard(idcard int) Option {
	return optionFunc(func(o *options) { o.query["idcard"] = idcard })
}

// WithTelephone telephone获取 联系电话
func (obj *_AdminMgr) WithTelephone(telephone int) Option {
	return optionFunc(func(o *options) { o.query["telephone"] = telephone })
}

// WithAddress address获取 家庭住址
func (obj *_AdminMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithBustPhoto bust_photo获取 半身照片
func (obj *_AdminMgr) WithBustPhoto(bustPhoto []byte) Option {
	return optionFunc(func(o *options) { o.query["bust_photo"] = bustPhoto })
}

// WithAvatar avatar获取 头像
func (obj *_AdminMgr) WithAvatar(avatar []byte) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}


// GetByOption 功能选项模式获取
func (obj *_AdminMgr) GetByOption(opts ...Option) (result Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query).Find(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AdminMgr) GetByOptions(opts ...Option) (results []*Admin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_AdminMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Admin,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Admin{}).Where(options.query)
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


// GetFromAdminID 通过adminID获取内容 管理员ID 
func (obj *_AdminMgr)  GetFromAdminID(adminID string) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`adminID` = ?", adminID).Find(&result).Error
	
	return
}

// GetBatchFromAdminID 批量查找 管理员ID
func (obj *_AdminMgr) GetBatchFromAdminID(adminIDs []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`adminID` IN (?)", adminIDs).Find(&results).Error
	
	return
}
 
// GetFromEstablishDay 通过establish_day获取内容 建立日期 
func (obj *_AdminMgr) GetFromEstablishDay(establishDay datatypes.Date) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`establish_day` = ?", establishDay).Find(&results).Error
	
	return
}

// GetBatchFromEstablishDay 批量查找 建立日期
func (obj *_AdminMgr) GetBatchFromEstablishDay(establishDays []datatypes.Date) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`establish_day` IN (?)", establishDays).Find(&results).Error
	
	return
}
 
// GetFromPasswd 通过passwd获取内容 管理员密码 
func (obj *_AdminMgr) GetFromPasswd(passwd string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`passwd` = ?", passwd).Find(&results).Error
	
	return
}

// GetBatchFromPasswd 批量查找 管理员密码
func (obj *_AdminMgr) GetBatchFromPasswd(passwds []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`passwd` IN (?)", passwds).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容 管理员姓名 
func (obj *_AdminMgr) GetFromName(name string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`name` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 管理员姓名
func (obj *_AdminMgr) GetBatchFromName(names []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`name` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromBirthday 通过birthday获取内容 出生年月 
func (obj *_AdminMgr) GetFromBirthday(birthday datatypes.Date) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`birthday` = ?", birthday).Find(&results).Error
	
	return
}

// GetBatchFromBirthday 批量查找 出生年月
func (obj *_AdminMgr) GetBatchFromBirthday(birthdays []datatypes.Date) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`birthday` IN (?)", birthdays).Find(&results).Error
	
	return
}
 
// GetFromIDcard 通过idcard获取内容 身份证号 
func (obj *_AdminMgr) GetFromIDcard(idcard int) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`idcard` = ?", idcard).Find(&results).Error
	
	return
}

// GetBatchFromIDcard 批量查找 身份证号
func (obj *_AdminMgr) GetBatchFromIDcard(idcards []int) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`idcard` IN (?)", idcards).Find(&results).Error
	
	return
}
 
// GetFromTelephone 通过telephone获取内容 联系电话 
func (obj *_AdminMgr) GetFromTelephone(telephone int) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`telephone` = ?", telephone).Find(&results).Error
	
	return
}

// GetBatchFromTelephone 批量查找 联系电话
func (obj *_AdminMgr) GetBatchFromTelephone(telephones []int) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`telephone` IN (?)", telephones).Find(&results).Error
	
	return
}
 
// GetFromAddress 通过address获取内容 家庭住址 
func (obj *_AdminMgr) GetFromAddress(address string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`address` = ?", address).Find(&results).Error
	
	return
}

// GetBatchFromAddress 批量查找 家庭住址
func (obj *_AdminMgr) GetBatchFromAddress(addresss []string) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`address` IN (?)", addresss).Find(&results).Error
	
	return
}
 
// GetFromBustPhoto 通过bust_photo获取内容 半身照片 
func (obj *_AdminMgr) GetFromBustPhoto(bustPhoto []byte) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`bust_photo` = ?", bustPhoto).Find(&results).Error
	
	return
}

// GetBatchFromBustPhoto 批量查找 半身照片
func (obj *_AdminMgr) GetBatchFromBustPhoto(bustPhotos [][]byte) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`bust_photo` IN (?)", bustPhotos).Find(&results).Error
	
	return
}
 
// GetFromAvatar 通过avatar获取内容 头像 
func (obj *_AdminMgr) GetFromAvatar(avatar []byte) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`avatar` = ?", avatar).Find(&results).Error
	
	return
}

// GetBatchFromAvatar 批量查找 头像
func (obj *_AdminMgr) GetBatchFromAvatar(avatars [][]byte) (results []*Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`avatar` IN (?)", avatars).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_AdminMgr) FetchByPrimaryKey(adminID string ) (result Admin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Admin{}).Where("`adminID` = ?", adminID).Find(&result).Error
	
	return
}
 

 

	

