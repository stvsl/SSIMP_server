package db

import (	
"time"	
"gorm.io/datatypes"	
)	
/******sql******		
CREATE TABLE `Admin` (
  `adminID` char(10) NOT NULL COMMENT '管理员ID',
  `establish_day` date NOT NULL COMMENT '建立日期',
  `passwd` varchar(64) NOT NULL COMMENT '管理员密码',
  `name` varchar(10) NOT NULL COMMENT '管理员姓名',
  `birthday` date NOT NULL COMMENT '出生年月',
  `idcard` char(18) NOT NULL COMMENT '身份证号',
  `telephone` varchar(14) NOT NULL COMMENT '联系电话',
  `address` varchar(40) NOT NULL COMMENT '家庭住址',
  `avatar` varchar(100) NOT NULL COMMENT '头像',
  PRIMARY KEY (`adminID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员信息表'		
******sql******/		
// Admin 管理员信息表		
type	Admin	struct {		
Adminid	string	`gorm:"primaryKey;column:adminID;type:char(10);not null;comment:'管理员ID'" json:"-"`	// 管理员ID			
EstablishDay	datatypes.Date	`gorm:"column:establish_day;type:date;not null;comment:'建立日期'" json:"establishDay"`	// 建立日期			
Passwd	string	`gorm:"column:passwd;type:varchar(64);not null;comment:'管理员密码'" json:"passwd"`	// 管理员密码			
Name	string	`gorm:"column:name;type:varchar(10);not null;comment:'管理员姓名'" json:"name"`	// 管理员姓名			
Birthday	datatypes.Date	`gorm:"column:birthday;type:date;not null;comment:'出生年月'" json:"birthday"`	// 出生年月			
IDcard	string	`gorm:"column:idcard;type:char(18);not null;comment:'身份证号'" json:"idcard"`	// 身份证号			
Telephone	string	`gorm:"column:telephone;type:varchar(14);not null;comment:'联系电话'" json:"telephone"`	// 联系电话			
Address	string	`gorm:"column:address;type:varchar(40);not null;comment:'家庭住址'" json:"address"`	// 家庭住址			
Avatar	string	`gorm:"column:avatar;type:varchar(100);not null;comment:'头像'" json:"avatar"`	// 头像			
}		

// TableName get sql table name.获取数据库表名
func (m *Admin) TableName() string {
	return "Admin"
}
	
/******sql******		
CREATE TABLE `Article` (
  `aid` int(12) NOT NULL AUTO_INCREMENT COMMENT '文章编号',
  `coverimg` varchar(150) NOT NULL COMMENT '封面图片',
  `contentimg` varchar(150) NOT NULL COMMENT '内容大图',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `introduction` varchar(200) NOT NULL COMMENT '简介',
  `text` longtext NOT NULL COMMENT '正文',
  `writetime` datetime NOT NULL COMMENT '发表日期',
  `updatetime` datetime NOT NULL COMMENT '更新日期',
  `author` varchar(10) NOT NULL COMMENT '作者',
  `pageviews` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '浏览量',
  `status` int(1) NOT NULL COMMENT '文章状态',
  PRIMARY KEY (`aid`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000019 DEFAULT CHARSET=utf8mb4 COMMENT='网站文章相关数据'		
******sql******/		
// Article 网站文章相关数据		
type	Article	struct {		
Aid	int	`gorm:"autoIncrement:true;primaryKey;column:aid;type:int(12);not null;comment:'文章编号'" json:"-"`	// 文章编号			
Coverimg	string	`gorm:"column:coverimg;type:varchar(150);not null;comment:'封面图片'" json:"coverimg"`	// 封面图片			
Contentimg	string	`gorm:"column:contentimg;type:varchar(150);not null;comment:'内容大图'" json:"contentimg"`	// 内容大图			
Title	string	`gorm:"column:title;type:varchar(50);not null;comment:'标题'" json:"title"`	// 标题			
Introduction	string	`gorm:"column:introduction;type:varchar(200);not null;comment:'简介'" json:"introduction"`	// 简介			
Text	string	`gorm:"column:text;type:longtext;not null;comment:'正文'" json:"text"`	// 正文			
Writetime	time.Time	`gorm:"column:writetime;type:datetime;not null;comment:'发表日期'" json:"writetime"`	// 发表日期			
Updatetime	time.Time	`gorm:"column:updatetime;type:datetime;not null;comment:'更新日期'" json:"updatetime"`	// 更新日期			
Author	string	`gorm:"column:author;type:varchar(10);not null;comment:'作者'" json:"author"`	// 作者			
Pageviews	uint64	`gorm:"column:pageviews;type:bigint(20) unsigned;not null;default:0;comment:'浏览量'" json:"pageviews"`	// 浏览量			
Status	int	`gorm:"column:status;type:int(1);not null;comment:'文章状态'" json:"status"`	// 文章状态			
}		

// TableName get sql table name.获取数据库表名
func (m *Article) TableName() string {
	return "Article"
}
	
/******sql******		
CREATE TABLE `Attendance` (
  `employid` char(10) NOT NULL COMMENT '员工编号',
  `tid` int(11) NOT NULL COMMENT '任务编号',
  `startTime` datetime NOT NULL COMMENT '开始时间',
  `endTime` datetime DEFAULT NULL COMMENT '结束时间',
  `task_completion` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '任务完成情况',
  `inspection_track` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '巡查轨迹',
  PRIMARY KEY (`employid`,`tid`,`startTime`),
  KEY `Attendance_FK` (`employid`),
  KEY `Attendance_FK_1` (`tid`),
  CONSTRAINT `Attendance_FK` FOREIGN KEY (`employid`) REFERENCES `Employer` (`employid`) ON UPDATE CASCADE,
  CONSTRAINT `Attendance_FK_1` FOREIGN KEY (`tid`) REFERENCES `TaskSet` (`tid`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='考勤表'		
******sql******/		
// Attendance 考勤表		
type	Attendance	struct {		
Employid	string	`gorm:"primaryKey;index:Attendance_FK;column:employid;type:char(10);not null;comment:'员工编号'" json:"-"`	// 员工编号			
Tid	int	`gorm:"primaryKey;index:Attendance_FK_1;column:tid;type:int(11);not null;comment:'任务编号'" json:"-"`	// 任务编号			
Starttime	time.Time	`gorm:"primaryKey;column:startTime;type:datetime;not null;comment:'开始时间'" json:"-"`	// 开始时间			
Endtime	time.Time	`gorm:"column:endTime;type:datetime;default:null;comment:'结束时间'" json:"endTime"`	// 结束时间			
TaskCompletion	string	`gorm:"column:task_completion;type:varchar(10);default:null;comment:'任务完成情况'" json:"taskCompletion"`	// 任务完成情况			
InspectionTrack	string	`gorm:"column:inspection_track;type:longtext;default:null;comment:'巡查轨迹'" json:"inspectionTrack"`	// 巡查轨迹			
}		

// TableName get sql table name.获取数据库表名
func (m *Attendance) TableName() string {
	return "Attendance"
}
	
/******sql******		
CREATE TABLE `Camera` (
  `id` uuid NOT NULL COMMENT 'id',
  `longitude` float NOT NULL DEFAULT 39.9 COMMENT '经度',
  `latitude` float NOT NULL DEFAULT 116.4 COMMENT '纬度',
  `comment` varchar(150) DEFAULT NULL COMMENT '备注',
  `ip` char(17) DEFAULT NULL COMMENT 'ip',
  `port` int(11) NOT NULL DEFAULT 1203 COMMENT '端口号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='摄像头检测人流量'		
******sql******/		
// Camera 摄像头检测人流量		
type	Camera	struct {		
ID	string	`gorm:"primaryKey;column:id;type:uuid;not null;comment:'id'" json:"-"`	// id			
Longitude	float32	`gorm:"column:longitude;type:float;not null;default:39.9;comment:'经度'" json:"longitude"`	// 经度			
Latitude	float32	`gorm:"column:latitude;type:float;not null;default:116.4;comment:'纬度'" json:"latitude"`	// 纬度			
Comment	string	`gorm:"column:comment;type:varchar(150);default:null;comment:'备注'" json:"comment"`	// 备注			
IP	string	`gorm:"column:ip;type:char(17);default:null;comment:'ip'" json:"ip"`	// ip			
Port	int	`gorm:"column:port;type:int(11);not null;default:1203;comment:'端口号'" json:"port"`	// 端口号			
}		

// TableName get sql table name.获取数据库表名
func (m *Camera) TableName() string {
	return "Camera"
}
	
/******sql******		
CREATE TABLE `Employer` (
  `employid` char(10) NOT NULL COMMENT '员工编号',
  `passwd` varchar(64) NOT NULL COMMENT '登录密码',
  `name` varchar(10) NOT NULL COMMENT '姓名',
  `birth_day` date NOT NULL COMMENT '出生日期',
  `employ_day` date NOT NULL COMMENT '入职日期',
  `idcard` varchar(18) NOT NULL COMMENT '身份证号',
  `address` varchar(40) NOT NULL COMMENT '家庭地址',
  `telephone` varchar(11) NOT NULL COMMENT '联系电话',
  `bust_photo` varchar(100) NOT NULL COMMENT '半身照',
  `avatar` varchar(100) NOT NULL COMMENT '头像',
  PRIMARY KEY (`employid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='员工信息表'		
******sql******/		
// Employer 员工信息表		
type	Employer	struct {		
Employid	string	`gorm:"primaryKey;column:employid;type:char(10);not null;comment:'员工编号'" json:"-"`	// 员工编号			
Passwd	string	`gorm:"column:passwd;type:varchar(64);not null;comment:'登录密码'" json:"passwd"`	// 登录密码			
Name	string	`gorm:"column:name;type:varchar(10);not null;comment:'姓名'" json:"name"`	// 姓名			
BirthDay	datatypes.Date	`gorm:"column:birth_day;type:date;not null;comment:'出生日期'" json:"birthDay"`	// 出生日期			
EmployDay	datatypes.Date	`gorm:"column:employ_day;type:date;not null;comment:'入职日期'" json:"employDay"`	// 入职日期			
IDcard	string	`gorm:"column:idcard;type:varchar(18);not null;comment:'身份证号'" json:"idcard"`	// 身份证号			
Address	string	`gorm:"column:address;type:varchar(40);not null;comment:'家庭地址'" json:"address"`	// 家庭地址			
Telephone	string	`gorm:"column:telephone;type:varchar(11);not null;comment:'联系电话'" json:"telephone"`	// 联系电话			
BustPhoto	string	`gorm:"column:bust_photo;type:varchar(100);not null;comment:'半身照'" json:"bustPhoto"`	// 半身照			
Avatar	string	`gorm:"column:avatar;type:varchar(100);not null;comment:'头像'" json:"avatar"`	// 头像			
}		

// TableName get sql table name.获取数据库表名
func (m *Employer) TableName() string {
	return "Employer"
}
	
/******sql******		
CREATE TABLE `Feedback` (
  `qid` int(10) NOT NULL AUTO_INCREMENT COMMENT '问题编号',
  `question` varchar(50) NOT NULL COMMENT '问题描述',
  `description` varchar(100) DEFAULT NULL COMMENT '问题详细描述',
  `picture` varchar(150) DEFAULT NULL COMMENT '问题图片',
  `create_date` datetime NOT NULL COMMENT '创建日期',
  `sponsor` varchar(100) NOT NULL COMMENT '发起人',
  `teleinfo` varchar(20) NOT NULL COMMENT '发起人联系方式',
  `principal` varchar(20) DEFAULT NULL COMMENT '委派负责人',
  `status` int(3) NOT NULL DEFAULT 1 COMMENT '处理进度',
  PRIMARY KEY (`qid`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COMMENT='问题及反馈表'		
******sql******/		
// Feedback 问题及反馈表		
type	Feedback	struct {		
Qid	int	`gorm:"autoIncrement:true;primaryKey;column:qid;type:int(10);not null;comment:'问题编号'" json:"-"`	// 问题编号			
Question	string	`gorm:"column:question;type:varchar(50);not null;comment:'问题描述'" json:"question"`	// 问题描述			
Description	string	`gorm:"column:description;type:varchar(100);default:null;comment:'问题详细描述'" json:"description"`	// 问题详细描述			
Picture	string	`gorm:"column:picture;type:varchar(150);default:null;comment:'问题图片'" json:"picture"`	// 问题图片			
CreateDate	time.Time	`gorm:"column:create_date;type:datetime;not null;comment:'创建日期'" json:"createDate"`	// 创建日期			
Sponsor	string	`gorm:"column:sponsor;type:varchar(100);not null;comment:'发起人'" json:"sponsor"`	// 发起人			
Teleinfo	string	`gorm:"column:teleinfo;type:varchar(20);not null;comment:'发起人联系方式'" json:"teleinfo"`	// 发起人联系方式			
Principal	string	`gorm:"column:principal;type:varchar(20);default:null;comment:'委派负责人'" json:"principal"`	// 委派负责人			
Status	int	`gorm:"column:status;type:int(3);not null;default:1;comment:'处理进度'" json:"status"`	// 处理进度			
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
Info	string	`gorm:"column:info;type:varchar(100);not null;comment:'日志信息'" json:"info"`	// 日志信息			
Time	time.Time	`gorm:"column:time;type:datetime;not null;comment:'时间'" json:"time"`	// 时间			
Text	string	`gorm:"column:text;type:text;default:null;comment:'日志内容'" json:"text"`	// 日志内容			
Extra	string	`gorm:"column:extra;type:longtext;default:null;comment:'附加信息'" json:"extra"`	// 附加信息			
}		

// TableName get sql table name.获取数据库表名
func (m *Syslog) TableName() string {
	return "Syslog"
}
	
/******sql******		
CREATE TABLE `Task` (
  `employid` char(10) NOT NULL COMMENT '员工编号',
  `tid` int(11) NOT NULL COMMENT '任务',
  KEY `Task_FK_1` (`tid`),
  KEY `Task_FK` (`employid`),
  CONSTRAINT `Task_FK` FOREIGN KEY (`employid`) REFERENCES `Employer` (`employid`),
  CONSTRAINT `Task_FK_1` FOREIGN KEY (`tid`) REFERENCES `TaskSet` (`tid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='任务表'		
******sql******/		
// Task 任务表		
type	Task	struct {		
Employid	string	`gorm:"index:Task_FK;column:employid;type:char(10);not null;comment:'员工编号'" json:"employid"`	// 员工编号			
Tid	int	`gorm:"index:Task_FK_1;column:tid;type:int(11);not null;comment:'任务'" json:"tid"`	// 任务			
}		

// TableName get sql table name.获取数据库表名
func (m *Task) TableName() string {
	return "Task"
}
	
/******sql******		
CREATE TABLE `TaskSet` (
  `tid` int(11) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `name` varchar(100) NOT NULL COMMENT '任务名称',
  `content` varchar(100) NOT NULL COMMENT '任务内容',
  `area` varchar(100) NOT NULL COMMENT '区域',
  `poslo` double NOT NULL COMMENT '经度',
  `posli` double NOT NULL COMMENT '纬度',
  `cycle` int(11) NOT NULL DEFAULT 7 COMMENT '周期(每周完成多少次)',
  `state` int(11) NOT NULL DEFAULT 0 COMMENT '任务状态',
  `duration` int(11) NOT NULL DEFAULT 1 COMMENT '任务时长',
  PRIMARY KEY (`tid`),
  UNIQUE KEY `NewTable_tid_IDX` (`tid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1015 DEFAULT CHARSET=utf8mb4 COMMENT='任务配置表'		
******sql******/		
// Taskset 任务配置表		
type	Taskset	struct {		
Tid	int	`gorm:"autoIncrement:true;primaryKey;unique;column:tid;type:int(11);not null;comment:'任务ID'" json:"-"`	// 任务ID			
Name	string	`gorm:"column:name;type:varchar(100);not null;comment:'任务名称'" json:"name"`	// 任务名称			
Content	string	`gorm:"column:content;type:varchar(100);not null;comment:'任务内容'" json:"content"`	// 任务内容			
Area	string	`gorm:"column:area;type:varchar(100);not null;comment:'区域'" json:"area"`	// 区域			
Poslo	float64	`gorm:"column:poslo;type:double;not null;comment:'经度'" json:"poslo"`	// 经度			
Posli	float64	`gorm:"column:posli;type:double;not null;comment:'纬度'" json:"posli"`	// 纬度			
Cycle	int	`gorm:"column:cycle;type:int(11);not null;default:7;comment:'周期(每周完成多少次)'" json:"cycle"`	// 周期(每周完成多少次)			
State	int	`gorm:"column:state;type:int(11);not null;default:0;comment:'任务状态'" json:"state"`	// 任务状态			
Duration	int	`gorm:"column:duration;type:int(11);not null;default:1;comment:'任务时长'" json:"duration"`	// 任务时长			
}		

// TableName get sql table name.获取数据库表名
func (m *Taskset) TableName() string {
	return "TaskSet"
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
Wxinfo	string	`gorm:"column:wxinfo;type:longtext;default:null;comment:'微信登录信息'" json:"wxinfo"`	// 微信登录信息			
ID	string	`gorm:"column:id;type:varchar(20);default:null;comment:'旅客id'" json:"id"`	// 旅客id			
Passwd	string	`gorm:"column:passwd;type:varchar(20);default:null;comment:'登录密码'" json:"passwd"`	// 登录密码			
History	string	`gorm:"column:history;type:longtext;default:null;comment:'旅游记录'" json:"history"`	// 旅游记录			
}		

// TableName get sql table name.获取数据库表名
func (m *Taveler) TableName() string {
	return "Taveler"
}
