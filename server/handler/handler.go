package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Path string `json:"path"`
	Size int64 `json:"size"`
}

func dirwalk(dir string) (files []File, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		path = strings.Replace(path, "images/", "http://localhost:8888/", 1)
		size := info.Size()
		f := File {
			Path: path,
			Size: size,
		}
		files = append(files, f)
		return nil
	})
	if err != nil {
		return
	}
	files = files[1:]
	return
}

func List(c *gin.Context) {
	files, err := dirwalk("./images")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message":err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	id := uuid.New()
	sid := id.String()

	for _, file := range files {

		pos := strings.LastIndex(file.Filename, ".")
		path := fmt.Sprintf("%s%s", sid, file.Filename[pos:])

		// TODO: 拡張子を取得する
		err := c.SaveUploadedFile(file, "images/" + path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func Delete(c *gin.Context) {
	path := c.Param("path")

	err := os.Remove(fmt.Sprintf("images/%s", path))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()} )
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id: %s is deleted", path)})
}
