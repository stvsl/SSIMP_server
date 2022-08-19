package service

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		Code.SE001(c)
	})
	router.Run(":8080")
}
