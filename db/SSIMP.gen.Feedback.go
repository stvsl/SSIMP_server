package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FeedbackMgr struct {
	*_BaseMgr
}

// FeedbackMgr open func
func FeedbackMgr(db *gorm.DB) *_FeedbackMgr {
	if db == nil {
		panic(fmt.Errorf("FeedbackMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FeedbackMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Feedback"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FeedbackMgr) GetTableName() string {
	return "Feedback"
}

// Reset 重置gorm会话
func (obj *_FeedbackMgr) Reset() *_FeedbackMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FeedbackMgr) Get() (result Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FeedbackMgr) Gets() (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FeedbackMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Feedback{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithQid qid获取 问题编号
func (obj *_FeedbackMgr) WithQid(qid int) Option {
	return optionFunc(func(o *options) { o.query["qid"] = qid })
}

// WithQuestion question获取 问题描述
func (obj *_FeedbackMgr) WithQuestion(question string) Option {
	return optionFunc(func(o *options) { o.query["question"] = question })
}

// WithDescription description获取 问题详细描述
func (obj *_FeedbackMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithPicture picture获取 问题图片
func (obj *_FeedbackMgr) WithPicture(picture []byte) Option {
	return optionFunc(func(o *options) { o.query["picture"] = picture })
}

// WithCreateDate create_date获取 创建日期
func (obj *_FeedbackMgr) WithCreateDate(createDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_date"] = createDate })
}

// WithSponsor sponsor获取 发起人
func (obj *_FeedbackMgr) WithSponsor(sponsor string) Option {
	return optionFunc(func(o *options) { o.query["sponsor"] = sponsor })
}

// WithTeleinfo teleinfo获取 发起人联系方式
func (obj *_FeedbackMgr) WithTeleinfo(teleinfo string) Option {
	return optionFunc(func(o *options) { o.query["teleinfo"] = teleinfo })
}

// WithPrincipal principal获取 委派负责人
func (obj *_FeedbackMgr) WithPrincipal(principal string) Option {
	return optionFunc(func(o *options) { o.query["principal"] = principal })
}

// WithStatus status获取 处理进度
func (obj *_FeedbackMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_FeedbackMgr) GetByOption(opts ...Option) (result Feedback, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FeedbackMgr) GetByOptions(opts ...Option) (results []*Feedback, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FeedbackMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Feedback, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where(options.query)
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

// GetFromQid 通过qid获取内容 问题编号
func (obj *_FeedbackMgr) GetFromQid(qid int) (result Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`qid` = ?", qid).Find(&result).Error

	return
}

// GetBatchFromQid 批量查找 问题编号
func (obj *_FeedbackMgr) GetBatchFromQid(qids []int) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`qid` IN (?)", qids).Find(&results).Error

	return
}

// GetFromQuestion 通过question获取内容 问题描述
func (obj *_FeedbackMgr) GetFromQuestion(question string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`question` = ?", question).Find(&results).Error

	return
}

// GetBatchFromQuestion 批量查找 问题描述
func (obj *_FeedbackMgr) GetBatchFromQuestion(questions []string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`question` IN (?)", questions).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 问题详细描述
func (obj *_FeedbackMgr) GetFromDescription(description string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 问题详细描述
func (obj *_FeedbackMgr) GetBatchFromDescription(descriptions []string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromPicture 通过picture获取内容 问题图片
func (obj *_FeedbackMgr) GetFromPicture(picture []byte) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`picture` = ?", picture).Find(&results).Error

	return
}

// GetBatchFromPicture 批量查找 问题图片
func (obj *_FeedbackMgr) GetBatchFromPicture(pictures [][]byte) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`picture` IN (?)", pictures).Find(&results).Error

	return
}

// GetFromCreateDate 通过create_date获取内容 创建日期
func (obj *_FeedbackMgr) GetFromCreateDate(createDate time.Time) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`create_date` = ?", createDate).Find(&results).Error

	return
}

// GetBatchFromCreateDate 批量查找 创建日期
func (obj *_FeedbackMgr) GetBatchFromCreateDate(createDates []time.Time) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`create_date` IN (?)", createDates).Find(&results).Error

	return
}

// GetFromSponsor 通过sponsor获取内容 发起人
func (obj *_FeedbackMgr) GetFromSponsor(sponsor string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`sponsor` = ?", sponsor).Find(&results).Error

	return
}

// GetBatchFromSponsor 批量查找 发起人
func (obj *_FeedbackMgr) GetBatchFromSponsor(sponsors []string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`sponsor` IN (?)", sponsors).Find(&results).Error

	return
}

// GetFromTeleinfo 通过teleinfo获取内容 发起人联系方式
func (obj *_FeedbackMgr) GetFromTeleinfo(teleinfo string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`teleinfo` = ?", teleinfo).Find(&results).Error

	return
}

// GetBatchFromTeleinfo 批量查找 发起人联系方式
func (obj *_FeedbackMgr) GetBatchFromTeleinfo(teleinfos []string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`teleinfo` IN (?)", teleinfos).Find(&results).Error

	return
}

// GetFromPrincipal 通过principal获取内容 委派负责人
func (obj *_FeedbackMgr) GetFromPrincipal(principal string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`principal` = ?", principal).Find(&results).Error

	return
}

// GetBatchFromPrincipal 批量查找 委派负责人
func (obj *_FeedbackMgr) GetBatchFromPrincipal(principals []string) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`principal` IN (?)", principals).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 处理进度
func (obj *_FeedbackMgr) GetFromStatus(status int) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 处理进度
func (obj *_FeedbackMgr) GetBatchFromStatus(statuss []int) (results []*Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FeedbackMgr) FetchByPrimaryKey(qid int) (result Feedback, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Feedback{}).Where("`qid` = ?", qid).Find(&result).Error

	return
}
