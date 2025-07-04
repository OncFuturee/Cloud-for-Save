package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 管理端登录校验中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("admin_token")
		if err != nil || token != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
			return
		}
		c.Next()
	}
}
