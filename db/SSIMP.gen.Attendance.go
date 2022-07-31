package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AttendanceMgr struct {
	*_BaseMgr
}

// AttendanceMgr open func
func AttendanceMgr(db *gorm.DB) *_AttendanceMgr {
	if db == nil {
		panic(fmt.Errorf("AttendanceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AttendanceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Attendance"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AttendanceMgr) GetTableName() string {
	return "Attendance"
}

// Reset 重置gorm会话
func (obj *_AttendanceMgr) Reset() *_AttendanceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AttendanceMgr) Get() (result Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Employer").Where("employid = ?", result.Employid).Find(&result.Employer).Error; err != nil { // 员工信息表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_AttendanceMgr) Gets() (results []*Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Find(&results).Error
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
func (obj *_AttendanceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Attendance{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithEmployid employid获取 员工编号
func (obj *_AttendanceMgr) WithEmployid(employid string) Option {
	return optionFunc(func(o *options) { o.query["employid"] = employid })
}

// WithStartTime startTime获取 开始时间
func (obj *_AttendanceMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["startTime"] = startTime })
}

// WithEndTime endTime获取 结束时间
func (obj *_AttendanceMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["endTime"] = endTime })
}

// WithTaskCompletion task_completion获取 任务完成情况
func (obj *_AttendanceMgr) WithTaskCompletion(taskCompletion string) Option {
	return optionFunc(func(o *options) { o.query["task_completion"] = taskCompletion })
}

// WithInspectionTrack inspection_track获取 巡查轨迹
func (obj *_AttendanceMgr) WithInspectionTrack(inspectionTrack string) Option {
	return optionFunc(func(o *options) { o.query["inspection_track"] = inspectionTrack })
}

// GetByOption 功能选项模式获取
func (obj *_AttendanceMgr) GetByOption(opts ...Option) (result Attendance, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Employer").Where("employid = ?", result.Employid).Find(&result.Employer).Error; err != nil { // 员工信息表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AttendanceMgr) GetByOptions(opts ...Option) (results []*Attendance, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where(options.query).Find(&results).Error
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
func (obj *_AttendanceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Attendance, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where(options.query)
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
func (obj *_AttendanceMgr) GetFromEmployid(employid string) (result Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`employid` = ?", employid).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Employer").Where("employid = ?", result.Employid).Find(&result.Employer).Error; err != nil { // 员工信息表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromEmployid 批量查找 员工编号
func (obj *_AttendanceMgr) GetBatchFromEmployid(employids []string) (results []*Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`employid` IN (?)", employids).Find(&results).Error
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

// GetFromStartTime 通过startTime获取内容 开始时间
func (obj *_AttendanceMgr) GetFromStartTime(startTime time.Time) (results []*Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`startTime` = ?", startTime).Find(&results).Error
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

// GetBatchFromStartTime 批量查找 开始时间
func (obj *_AttendanceMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`startTime` IN (?)", startTimes).Find(&results).Error
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

// GetFromEndTime 通过endTime获取内容 结束时间
func (obj *_AttendanceMgr) GetFromEndTime(endTime time.Time) (results []*Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`endTime` = ?", endTime).Find(&results).Error
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

// GetBatchFromEndTime 批量查找 结束时间
func (obj *_AttendanceMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`endTime` IN (?)", endTimes).Find(&results).Error
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

// GetFromTaskCompletion 通过task_completion获取内容 任务完成情况
func (obj *_AttendanceMgr) GetFromTaskCompletion(taskCompletion string) (results []*Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`task_completion` = ?", taskCompletion).Find(&results).Error
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

// GetBatchFromTaskCompletion 批量查找 任务完成情况
func (obj *_AttendanceMgr) GetBatchFromTaskCompletion(taskCompletions []string) (results []*Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`task_completion` IN (?)", taskCompletions).Find(&results).Error
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

// GetFromInspectionTrack 通过inspection_track获取内容 巡查轨迹
func (obj *_AttendanceMgr) GetFromInspectionTrack(inspectionTrack string) (results []*Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`inspection_track` = ?", inspectionTrack).Find(&results).Error
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

// GetBatchFromInspectionTrack 批量查找 巡查轨迹
func (obj *_AttendanceMgr) GetBatchFromInspectionTrack(inspectionTracks []string) (results []*Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`inspection_track` IN (?)", inspectionTracks).Find(&results).Error
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
func (obj *_AttendanceMgr) FetchByPrimaryKey(employid string) (result Attendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Attendance{}).Where("`employid` = ?", employid).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("Employer").Where("employid = ?", result.Employid).Find(&result.Employer).Error; err != nil { // 员工信息表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
