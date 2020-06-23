package main

import (
	"consulapi/controller"
	_ "consulapi/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

//var ConsulServer = flag.String("ip", "127.0.0.1", "Consul Server IP")
//var ConsulPort = flag.Int("port", 8500, "Consul Server Port")

// @Title Consul Service Api
// @Version 1.0
// @Description Consul API For Register and Deregister
func main() {
	//flag.Parse()
	gin.SetMode("release")
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(Cors())
	r.Use(Server())
	r.POST("/consul", controller.Register)
	r.DELETE("/consul/:id", controller.Deregister)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Server", "Gin")
		context.Next()
	}
}

func Server() gin.HandlerFunc {
	return func(context *gin.Context) {
		SERVER := os.Getenv("CONSULHOST")
		context.Set("server", SERVER)
		context.Set("port", "8500")
		context.Next()
	}
}
