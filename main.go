package main

import (
	"Gin/Job"
	"Gin/models"
	"Gin/route"
	"Gin/until"
	"github.com/spf13/viper"
)

//var PDO *gorm.DB


func main() {
	//gin.SetMode(gin.ReleaseMode)
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath("./conf")
	viper.ReadInConfig()

	models.Conn()
	until.LogObj()
	Job.InitJobs()
	r := route.Route()
	//默认为监听8080端口
	r.Run(":8000")
}
