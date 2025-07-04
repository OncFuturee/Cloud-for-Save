package main

import (
	"cloud-for-save/internal/api"
	"cloud-for-save/internal/config"
	"cloud-for-save/pkg/logger"
	"fmt"
	"os"
)

func main() {
	_ = logger.Init("cloud-for-save.log")
	defer logger.Close()

	if err := config.LoadConfig(); err != nil {
		fmt.Println("配置加载失败:", err)
		logger.Error("配置加载失败:", err)
		os.Exit(1)
	}

	r := api.SetupRouter()
	port := config.GetConfig().Server.Port

	fmt.Println("Cloud for Save 服务启动...")
	fmt.Printf("http://localhost:%d\n", port)

	r.Run(fmt.Sprintf(":%d", port))
}
