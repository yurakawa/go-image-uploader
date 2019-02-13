package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/jinzhu/gorm"
	"github.com/yurakawa/go-image-uploader/server/config"
	"github.com/yurakawa/go-image-uploader/server/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


var err error

func main() {

	config.DB, err = gorm.Open("mysql", "uploader:uploader@tcp(127.0.0.1:3306)/uploader")
	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()

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