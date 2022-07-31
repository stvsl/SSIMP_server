package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ArticleMgr struct {
	*_BaseMgr
}

// ArticleMgr open func
func ArticleMgr(db *gorm.DB) *_ArticleMgr {
	if db == nil {
		panic(fmt.Errorf("ArticleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArticleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Article"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArticleMgr) GetTableName() string {
	return "Article"
}

// Reset 重置gorm会话
func (obj *_ArticleMgr) Reset() *_ArticleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ArticleMgr) Get() (result Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArticleMgr) Gets() (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ArticleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Article{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAid aid获取 文章编号
func (obj *_ArticleMgr) WithAid(aid int) Option {
	return optionFunc(func(o *options) { o.query["aid"] = aid })
}

// WithTitle title获取 标题
func (obj *_ArticleMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithText text获取 正文
func (obj *_ArticleMgr) WithText(text string) Option {
	return optionFunc(func(o *options) { o.query["text"] = text })
}

// WithWritetime writetime获取 发表日期
func (obj *_ArticleMgr) WithWritetime(writetime time.Time) Option {
	return optionFunc(func(o *options) { o.query["writetime"] = writetime })
}

// WithUpdatetime updatetime获取 更新日期
func (obj *_ArticleMgr) WithUpdatetime(updatetime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updatetime"] = updatetime })
}

// WithAuthor author获取 作者
func (obj *_ArticleMgr) WithAuthor(author string) Option {
	return optionFunc(func(o *options) { o.query["author"] = author })
}

// WithPageviews pageviews获取 浏览量
func (obj *_ArticleMgr) WithPageviews(pageviews uint64) Option {
	return optionFunc(func(o *options) { o.query["pageviews"] = pageviews })
}

// WithStatus status获取 文章状态
func (obj *_ArticleMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_ArticleMgr) GetByOption(opts ...Option) (result Article, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ArticleMgr) GetByOptions(opts ...Option) (results []*Article, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ArticleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Article, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Article{}).Where(options.query)
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

// GetFromAid 通过aid获取内容 文章编号
func (obj *_ArticleMgr) GetFromAid(aid int) (result Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`aid` = ?", aid).Find(&result).Error

	return
}

// GetBatchFromAid 批量查找 文章编号
func (obj *_ArticleMgr) GetBatchFromAid(aids []int) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`aid` IN (?)", aids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_ArticleMgr) GetFromTitle(title string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_ArticleMgr) GetBatchFromTitle(titles []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromText 通过text获取内容 正文
func (obj *_ArticleMgr) GetFromText(text string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`text` = ?", text).Find(&results).Error

	return
}

// GetBatchFromText 批量查找 正文
func (obj *_ArticleMgr) GetBatchFromText(texts []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`text` IN (?)", texts).Find(&results).Error

	return
}

// GetFromWritetime 通过writetime获取内容 发表日期
func (obj *_ArticleMgr) GetFromWritetime(writetime time.Time) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`writetime` = ?", writetime).Find(&results).Error

	return
}

// GetBatchFromWritetime 批量查找 发表日期
func (obj *_ArticleMgr) GetBatchFromWritetime(writetimes []time.Time) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`writetime` IN (?)", writetimes).Find(&results).Error

	return
}

// GetFromUpdatetime 通过updatetime获取内容 更新日期
func (obj *_ArticleMgr) GetFromUpdatetime(updatetime time.Time) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`updatetime` = ?", updatetime).Find(&results).Error

	return
}

// GetBatchFromUpdatetime 批量查找 更新日期
func (obj *_ArticleMgr) GetBatchFromUpdatetime(updatetimes []time.Time) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`updatetime` IN (?)", updatetimes).Find(&results).Error

	return
}

// GetFromAuthor 通过author获取内容 作者
func (obj *_ArticleMgr) GetFromAuthor(author string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`author` = ?", author).Find(&results).Error

	return
}

// GetBatchFromAuthor 批量查找 作者
func (obj *_ArticleMgr) GetBatchFromAuthor(authors []string) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`author` IN (?)", authors).Find(&results).Error

	return
}

// GetFromPageviews 通过pageviews获取内容 浏览量
func (obj *_ArticleMgr) GetFromPageviews(pageviews uint64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`pageviews` = ?", pageviews).Find(&results).Error

	return
}

// GetBatchFromPageviews 批量查找 浏览量
func (obj *_ArticleMgr) GetBatchFromPageviews(pageviewss []uint64) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`pageviews` IN (?)", pageviewss).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 文章状态
func (obj *_ArticleMgr) GetFromStatus(status int) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 文章状态
func (obj *_ArticleMgr) GetBatchFromStatus(statuss []int) (results []*Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ArticleMgr) FetchByPrimaryKey(aid int) (result Article, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Article{}).Where("`aid` = ?", aid).Find(&result).Error

	return
}
