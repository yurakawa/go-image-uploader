package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	uuid := c.PostForm("uuid")


	for _, file := range files {
		// TODO: 拡張子を取得する
		err := c.SaveUploadedFile(file, "images/" + uuid + ".png")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	err := os.Remove(fmt.Sprintf("images/%s.png", uuid))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()} )
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id: %s is deleted", uuid)})
}
