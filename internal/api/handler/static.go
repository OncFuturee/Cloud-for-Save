package handler

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// ServeAdminPage 处理管理页面首页
func ServeAdminPage(c *gin.Context) {
	path := filepath.Join("internal", "web", "templates", "index.html")
	c.File(path)
}
