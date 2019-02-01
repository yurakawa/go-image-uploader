package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	uuid := c.PostForm("uuid")


	for _, file := range files {
		err := c.SaveUploadedFile(file, "images/" + uuid + ".jpg")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
