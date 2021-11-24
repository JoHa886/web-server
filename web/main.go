package main

import (
	"web-server/web/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/view", "./view")
	router.GET("/api/session", controller.GetSession)
	// router.GET("/", func(c *gin.Context) {
	// 	c.Writer.WriteString("项目开始了。。。")
	// })
	router.Run(":8090")
}
