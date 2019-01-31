package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		err := c.SaveUploadedFile(file, "images/" + file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "success!!"})
}