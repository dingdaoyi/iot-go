package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在处理请求之前执行的逻辑
		c.Next() // 继续处理请求
		// 在处理请求之后执行的逻辑
	}
}
