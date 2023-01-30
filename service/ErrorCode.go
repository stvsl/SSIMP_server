package service

import "github.com/gin-gonic/gin"

type Error func(c *gin.Context)

/*
*
  - 报错码范围区间：
  - SE100-SE199 信息报错响应
  - SE200-SE299 正常响应
  - SE300-SE399 暂时未定义
  - SE400-SE499 客户端错误
  - SE500-SE599	服务器内部错误
  - SE600-SE699	数据库错误
    SE001 通讯不允许，需检查密钥
    SE002 通讯密钥解密失败
    SE400 客户端发来的数据格式不正确
    SE401 客户端发来的数据内容不正确
    SE402 客户端数据无法解析
    SE403 客户端流量异常
    SE404 客户端请求的数据未找到
    SE405 客户端无权请求此数据
    SE406 客户端身份验证失败
    SE407 客户端请求未授权Token
    SE500 服务器内部报错
    SE501 服务器不支持请求的功能，无法完成请求
    SE505 服务器接收的数据版本为老版本，现已不支持
    SE506 服务器已将请求拒绝
    SE600 SQL数据库未知异常
    SE601 SQL数据库繁忙
    SE602 SQL数据库无法完成当前请求
    SE610 Redis数据库异常
*/
var Code *Error

// 通讯不允许，需检查密钥
func (Code *Error) SE001(c *gin.Context) {
	c.JSON(400, gin.H{
		"code": "SE001",
		"msg":  "通讯不允许,需要检查密钥",
	})
}

// 通讯密钥解析失败
func (Code *Error) SE002(c *gin.Context) {
	c.JSON(400, gin.H{
		"code": "SE002",
		"msg":  "通讯密钥解析失败",
	})
}

// 客户端发来的数据格式不正确
func (Code *Error) SE400(c *gin.Context) {
	c.JSON(400, gin.H{
		"code": "SE400",
		"msg":  "客户端发来的数据格式不正确",
	})
}

func (Code *Error) SE401(c *gin.Context) {
	c.JSON(401, gin.H{
		"code": "SE401",
		"msg":  "客户端发来的数据内容不正确",
	})
}

func (Code *Error) SE402(c *gin.Context) {
	c.JSON(402, gin.H{
		"code": "SE402",
		"msg":  "客户端数据无法解析",
	})
}

func (Code *Error) SE403(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE403",
		"msg":  "客户端流量异常",
	})
}

func (Code *Error) SE404(c *gin.Context) {
	c.JSON(404, gin.H{
		"code": "SE404",
		"msg":  "客户端请求的数据未找到",
	})
}

func (Code *Error) SE405(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE405",
		"msg":  "客户端无权请求此数据",
	})
}

func (Code *Error) SE406(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE406",
		"msg":  "客户端身份验证失败",
	})
}

func (Code *Error) SE407(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE407",
		"msg":  "客户端请求未授权Token或Token已过期或Token未携带",
	})
}

func (Code *Error) SE500(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE500",
		"msg":  "服务器内部报错",
	})
}

func (Code *Error) SE501(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE501",
		"msg":  "服务器不支持请求的功能，无法完成请求",
	})
}

func (Code *Error) SE505(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE505",
		"msg":  "服务器接收的数据版本为老版本，现已不支持",
	})
}

func (Code *Error) SE506(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE506",
		"msg":  "服务器已将请求拒绝",
	})
}

func (Code *Error) SE600(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE600",
		"msg":  "SQL数据库未知异常",
	})
}

func (Code *Error) SE601(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE601",
		"msg":  "SQL数据库繁忙",
	})
}

func (Code *Error) SE602(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE602",
		"msg":  "SQL数据库无法完成当前请求",
	})
}

func (Code *Error) SE610(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "SE610",
		"msg":  "Redis数据库异常",
	})
}
