package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/yurakawa/go-image-uploader/server/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	r.Use(static.Serve("/", static.LocalFile("./images", true)))
	r.GET("/images", handler.List)
	r.POST("/images", handler.Upload)
	r.DELETE("/images/:uuid", handler.Delete)
	r.Run(":8888")
}