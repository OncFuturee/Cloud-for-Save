package api

import (
	"cloud-for-save/internal/api/handler"
	"cloud-for-save/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 静态文件和管理页面
	r.Static("/assets", "./internal/web/assets")
	r.GET("/", handler.ServeAdminPage)

	r.POST("/api/login", handler.Login)

	admin := r.Group("/api/admin", middleware.AdminAuth())
	{
		admin.GET("/config", handler.GetConfig)
		admin.POST("/config", handler.SaveConfig)
		admin.POST("/config/reset", handler.ResetConfig)
	}

	return r
}
