package main

import (
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/sunliang711/insurance/config"
	"github.com/sunliang711/insurance/handler"
	"github.com/sunliang711/insurance/model"
	"github.com/sunliang711/insurance/utils"

	log "github.com/sirupsen/logrus"
	_ "github.com/sunliang711/insurance/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host aliyun.eagle711.win:8081
// @BasePath /
// 2019/08/08 13:19:12
func main() {
	dsn := viper.GetString("mysql.dsn")
	logrus.Infof("mysql dsn: %v", dsn)
	model.InitMysql(dsn)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	log.SetLevel(utils.LogLevel(viper.GetString("server.loglevel")))

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	router.Use(cors.New(corConfig))
	// router.GET()
	router.GET("/location", handler.Location)
	router.GET("/reason", handler.Reason)
	router.GET("/type", handler.Typ)
	router.POST("/claim", handler.Claim)
	router.POST("/decrypt", handler.Decrypt)
	//swagger handler
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	logrus.Infof("Listen on %v", addr)
	router.Run(addr)
}
