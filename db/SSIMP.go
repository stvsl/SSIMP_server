package	model	
import (	
"gorm.io/datatypes"	
"time"	
)	
/******sql******		
CREATE TABLE `Admin` (
  `adminID` char(10) NOT NULL COMMENT '管理员ID',
  `establish_day` date NOT NULL COMMENT '建立日期',
  `passwd` varchar(18) NOT NULL COMMENT '管理员密码',
  `name` varchar(10) NOT NULL COMMENT '管理员姓名',
  `birthday` date NOT NULL COMMENT '出生年月',
  `idcard` int(18) NOT NULL COMMENT '身份证号',
  `telephone` int(11) NOT NULL COMMENT '联系电话',
  `address` varchar(40) NOT NULL COMMENT '家庭住址',
  `bust_photo` blob DEFAULT NULL COMMENT '半身照片',
  `avatar` blob DEFAULT NULL COMMENT '头像',
  PRIMARY KEY (`adminID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员信息表'		
******sql******/		
// Admin 管理员信息表		
type	Admin	struct {		
AdminID	string	`gorm:"primaryKey;column:adminID;type:char(10);not null" json:"-"`	// 管理员ID			
EstablishDay	datatypes.Date	`gorm:"column:establish_day;type:date;not null" json:"establishDay"`	// 建立日期			
Passwd	string	`gorm:"column:passwd;type:varchar(18);not null" json:"passwd"`	// 管理员密码			
Name	string	`gorm:"column:name;type:varchar(10);not null" json:"name"`	// 管理员姓名			
Birthday	datatypes.Date	`gorm:"column:birthday;type:date;not null" json:"birthday"`	// 出生年月			
IDcard	int	`gorm:"column:idcard;type:int(18);not null" json:"idcard"`	// 身份证号			
Telephone	int	`gorm:"column:telephone;type:int(11);not null" json:"telephone"`	// 联系电话			
Address	string	`gorm:"column:address;type:varchar(40);not null" json:"address"`	// 家庭住址			
BustPhoto	[]byte	`gorm:"column:bust_photo;type:blob" json:"bustPhoto"`	// 半身照片			
Avatar	[]byte	`gorm:"column:avatar;type:blob" json:"avatar"`	// 头像			
}		

// TableName get sql table name.获取数据库表名
func (m *Admin) TableName() string {
	return "Admin"
}
	
/******sql******		
CREATE TABLE `Article` (
  `aid` int(20) NOT NULL COMMENT '文章编号',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `text` longtext NOT NULL COMMENT '正文',
  `writetime` datetime NOT NULL COMMENT '发表日期',
  `updatetime` datetime NOT NULL COMMENT '更新日期',
  `author` varchar(10) NOT NULL COMMENT '作者',
  `pageviews` bigint(20) unsigned NOT NULL COMMENT '浏览量',
  `status` int(1) NOT NULL COMMENT '文章状态',
  PRIMARY KEY (`aid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='网站文章相关数据'		
******sql******/		
// Article 网站文章相关数据		
type	Article	struct {		
Aid	int	`gorm:"primaryKey;column:aid;type:int(20);not null" json:"-"`	// 文章编号			
Title	string	`gorm:"column:title;type:varchar(50);not null" json:"title"`	// 标题			
Text	string	`gorm:"column:text;type:longtext;not null" json:"text"`	// 正文			
Writetime	time.Time	`gorm:"column:writetime;type:datetime;not null" json:"writetime"`	// 发表日期			
Updatetime	time.Time	`gorm:"column:updatetime;type:datetime;not null" json:"updatetime"`	// 更新日期			
Author	string	`gorm:"column:author;type:varchar(10);not null" json:"author"`	// 作者			
Pageviews	uint64	`gorm:"column:pageviews;type:bigint(20) unsigned;not null" json:"pageviews"`	// 浏览量			
Status	int	`gorm:"column:status;type:int(1);not null" json:"status"`	// 文章状态			
}		

// TableName get sql table name.获取数据库表名
func (m *Article) TableName() string {
	return "Article"
}
	
/******sql******		
CREATE TABLE `Attendance` (
  `employid` char(10) NOT NULL COMMENT '员工编号',
  `startTime` datetime NOT NULL COMMENT '开始时间',
  `endTime` datetime DEFAULT NULL COMMENT '结束时间',
  `task_completion` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '任务完成情况',
  `inspection_track` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '巡查轨迹' CHECK (json_valid(`inspection_track`)),
  PRIMARY KEY (`employid`),
  CONSTRAINT `Attendance_FK` FOREIGN KEY (`employid`) REFERENCES `Employer` (`employid`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='考勤表'		
******sql******/		
// Attendance 考勤表		
type	Attendance	struct {		
Employid	string	`gorm:"primaryKey;column:employid;type:char(10);not null" json:"-"`	// 员工编号			
Employer	Employer	`gorm:"joinForeignKey:employid;foreignKey:employid" json:"employerList"`	// 员工信息表			
StartTime	time.Time	`gorm:"column:startTime;type:datetime;not null" json:"startTime"`	// 开始时间			
EndTime	time.Time	`gorm:"column:endTime;type:datetime" json:"endTime"`	// 结束时间			
TaskCompletion	string	`gorm:"column:task_completion;type:longtext" json:"taskCompletion"`	// 任务完成情况			
InspectionTrack	string	`gorm:"column:inspection_track;type:longtext" json:"inspectionTrack"`	// 巡查轨迹			
}		

// TableName get sql table name.获取数据库表名
func (m *Attendance) TableName() string {
	return "Attendance"
}
	
/******sql******		
CREATE TABLE `Employer` (
  `employid` char(10) NOT NULL COMMENT '员工编号',
  `passwd` varchar(18) NOT NULL COMMENT '登录密码',
  `name` varchar(10) NOT NULL COMMENT '姓名',
  `birth_day` date NOT NULL COMMENT '出生日期',
  `employ_day` date NOT NULL COMMENT '入职日期',
  `idcard` int(18) NOT NULL COMMENT '身份证号',
  `address` varchar(40) NOT NULL COMMENT '家庭地址',
  `telephone` int(11) NOT NULL COMMENT '联系电话',
  `bust_photo` blob DEFAULT NULL COMMENT '半身照',
  `avatar` blob DEFAULT NULL COMMENT '头像',
  PRIMARY KEY (`employid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='员工信息表'		
******sql******/		
// Employer 员工信息表		
type	Employer	struct {		
Employid	string	`gorm:"primaryKey;column:employid;type:char(10);not null" json:"-"`	// 员工编号			
Passwd	string	`gorm:"column:passwd;type:varchar(18);not null" json:"passwd"`	// 登录密码			
Name	string	`gorm:"column:name;type:varchar(10);not null" json:"name"`	// 姓名			
BirthDay	datatypes.Date	`gorm:"column:birth_day;type:date;not null" json:"birthDay"`	// 出生日期			
EmployDay	datatypes.Date	`gorm:"column:employ_day;type:date;not null" json:"employDay"`	// 入职日期			
IDcard	int	`gorm:"column:idcard;type:int(18);not null" json:"idcard"`	// 身份证号			
Address	string	`gorm:"column:address;type:varchar(40);not null" json:"address"`	// 家庭地址			
Telephone	int	`gorm:"column:telephone;type:int(11);not null" json:"telephone"`	// 联系电话			
BustPhoto	[]byte	`gorm:"column:bust_photo;type:blob" json:"bustPhoto"`	// 半身照			
Avatar	[]byte	`gorm:"column:avatar;type:blob" json:"avatar"`	// 头像			
}		

// TableName get sql table name.获取数据库表名
func (m *Employer) TableName() string {
	return "Employer"
}
	
/******sql******		
CREATE TABLE `Feedback` (
  `qid` int(20) NOT NULL COMMENT '问题编号',
  `question` varchar(50) NOT NULL COMMENT '问题描述',
  `description` varchar(100) DEFAULT NULL COMMENT '问题详细描述',
  `picture` blob DEFAULT NULL COMMENT '问题图片',
  `create_date` datetime NOT NULL COMMENT '创建日期',
  `sponsor` varchar(100) NOT NULL COMMENT '发起人',
  `teleinfo` varchar(20) NOT NULL COMMENT '发起人联系方式',
  `principal` char(10) DEFAULT NULL COMMENT '委派负责人',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT '处理进度',
  PRIMARY KEY (`qid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='问题及反馈表'		
******sql******/		
// Feedback 问题及反馈表		
type	Feedback	struct {		
Qid	int	`gorm:"primaryKey;column:qid;type:int(20);not null" json:"-"`	// 问题编号			
Question	string	`gorm:"column:question;type:varchar(50);not null" json:"question"`	// 问题描述			
Description	string	`gorm:"column:description;type:varchar(100)" json:"description"`	// 问题详细描述			
Picture	[]byte	`gorm:"column:picture;type:blob" json:"picture"`	// 问题图片			
CreateDate	time.Time	`gorm:"column:create_date;type:datetime;not null" json:"createDate"`	// 创建日期			
Sponsor	string	`gorm:"column:sponsor;type:varchar(100);not null" json:"sponsor"`	// 发起人			
Teleinfo	string	`gorm:"column:teleinfo;type:varchar(20);not null" json:"teleinfo"`	// 发起人联系方式			
Principal	string	`gorm:"column:principal;type:char(10)" json:"principal"`	// 委派负责人			
Status	int	`gorm:"column:status;type:int(11);not null;default:1" json:"status"`	// 处理进度			
}		

// TableName get sql table name.获取数据库表名
func (m *Feedback) TableName() string {
	return "Feedback"
}
	
/******sql******		
CREATE TABLE `Syslog` (
  `info` varchar(100) NOT NULL COMMENT '日志信息',
  `time` datetime NOT NULL COMMENT '时间',
  `text` text DEFAULT NULL COMMENT '日志内容',
  `extra` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '附加信息' CHECK (json_valid(`extra`))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统日志'		
******sql******/		
// Syslog 系统日志		
type	Syslog	struct {		
Info	string	`gorm:"column:info;type:varchar(100);not null" json:"info"`	// 日志信息			
Time	time.Time	`gorm:"column:time;type:datetime;not null" json:"time"`	// 时间			
Text	string	`gorm:"column:text;type:text" json:"text"`	// 日志内容			
Extra	string	`gorm:"column:extra;type:longtext" json:"extra"`	// 附加信息			
}		

// TableName get sql table name.获取数据库表名
func (m *Syslog) TableName() string {
	return "Syslog"
}
	
/******sql******		
CREATE TABLE `Task` (
  `employid` char(10) NOT NULL COMMENT '员工编号',
  `task` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '任务' CHECK (json_valid(`task`)),
  PRIMARY KEY (`employid`),
  CONSTRAINT `Task_FK` FOREIGN KEY (`employid`) REFERENCES `Employer` (`employid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='任务表'		
******sql******/		
// Task 任务表		
type	Task	struct {		
Employid	string	`gorm:"primaryKey;column:employid;type:char(10);not null" json:"-"`	// 员工编号			
Employer	Employer	`gorm:"joinForeignKey:employid;foreignKey:employid" json:"employerList"`	// 员工信息表			
Task	string	`gorm:"column:task;type:longtext;not null" json:"task"`	// 任务			
}		

// TableName get sql table name.获取数据库表名
func (m *Task) TableName() string {
	return "Task"
}
	
/******sql******		
CREATE TABLE `Taveler` (
  `wxinfo` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '微信登录信息' CHECK (json_valid(`wxinfo`)),
  `id` varchar(20) DEFAULT NULL COMMENT '旅客id',
  `passwd` varchar(20) DEFAULT NULL COMMENT '登录密码',
  `history` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '旅游记录' CHECK (json_valid(`history`))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='旅客表'		
******sql******/		
// Taveler 旅客表		
type	Taveler	struct {		
Wxinfo	string	`gorm:"column:wxinfo;type:longtext" json:"wxinfo"`	// 微信登录信息			
ID	string	`gorm:"column:id;type:varchar(20)" json:"id"`	// 旅客id			
Passwd	string	`gorm:"column:passwd;type:varchar(20)" json:"passwd"`	// 登录密码			
History	string	`gorm:"column:history;type:longtext" json:"history"`	// 旅游记录			
}		

// TableName get sql table name.获取数据库表名
func (m *Taveler) TableName() string {
	return "Taveler"
}
	

