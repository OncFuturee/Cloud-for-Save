package handler

import (
	"cloud-for-save/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录请求
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if req.Username == "admin" && req.Password == "admin" {
		c.SetCookie("admin_token", "admin", 3600, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
}

// 获取当前配置
func GetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, config.GetConfig())
}

// 保存配置
func SaveConfig(c *gin.Context) {
	var cfg config.Config
	if err := c.ShouldBindJSON(&cfg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if err := config.SaveConfig(cfg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "保存成功"})
}

// 恢复默认配置
func ResetConfig(c *gin.Context) {
	if err := config.ResetToDefault(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "恢复失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "已恢复默认设置"})
}
