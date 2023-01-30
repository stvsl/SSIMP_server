package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"stvsljl.com/SSIMP/security"
)

const TokenExpireDuration = time.Hour * 2 // token过期时间

type Claims struct {
	AES string `json:"aes"`
	jwt.StandardClaims
}

// 生成jwt
func GenToken(aes string) (string, error) {
	// 创建声明
	c := Claims{
		aes,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			NotBefore: time.Now().Unix() - 10,
			Issuer:    "CSystemServer",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	priv, _ := jwt.ParseRSAPrivateKeyFromPEM(security.SERVER_RSA.PRIVATE_KEY)
	return token.SignedString(priv)
}

// 解析jwt
func ParseToken(tokenstr string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenstr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 这里的密钥要和生成的密钥一致
		return security.SERVER_RSA.PUBLIC_KEY, nil
	})
	if err != nil {
		fmt.Println("解析token失败")
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		fmt.Println("token无效")
		return nil, errors.New("token无效")
	}
}

// token校验
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			Code.SE407(c)
			c.Abort()
			return
		}
		// 解析token部分
		message, err := ParseToken(authHeader)
		if err != nil {
			Code.SE407(c)
			c.Abort()
			return
		}
		// 将当前请求的信息保存到请求的上下文c上
		c.Set("aes", message.AES)
		c.Next()
	}
}
