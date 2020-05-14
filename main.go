package main

import (
	"gin/config"
	"gin/controllers"
	"gin/models"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)
func main() {
	dir ,err := os.Getwd();
	if err !=nil {
		log.Fatal("can not get wd:", err.Error())
	}
	//加载配置
	config.Load(dir);
	cfg := config.GetConfig()
	setup(cfg)
	//数据迁移
	models.GetDb().AutoMigrate(&models.User{},&models.Content{})
	r := gin.Default()
	//加载路由
	controllers.Init(r)

	if cfg.Server.Listen!="" {
		panic(r.Run(cfg.Server.Listen))
		return
	}
	panic(r.Run())
}
func setup(cfg *config.Config){
	models.InitDb(cfg.Mysql,cfg.Server.Debug)
}