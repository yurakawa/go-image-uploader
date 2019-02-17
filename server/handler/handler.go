package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yurakawa/go-image-uploader/server/model"
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
	var files []model.File
	err := model.GetAllFile(&files)

	var f []model.File
	for _, file := range files {
		file.Path =  "http://localhost:8888/" + file.Path
		f = append(f, file)
	}

	// files, err := dirwalk("./images")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message":err.Error()})
		return
	}
	c.JSON(http.StatusOK, f)
}

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	id := uuid.New()
	sid := id.String()

	for _, file := range files {


		pos := strings.LastIndex(file.Filename, ".")
		path := fmt.Sprintf("%s%s", sid, file.Filename[pos:])

		err := c.SaveUploadedFile(file, "images/" + path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		}

		f := model.File{
			Name: file.Filename,
			Path : path,
			Size: file.Size,
		}
		err = model.AddNewFile(&f)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func Delete(c *gin.Context) {
	path := c.Param("uuid")

	err := model.DeleteFile(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()} )
		return
	}

	err = os.Remove(fmt.Sprintf("images/%s", path))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()} )
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id: %s is deleted", path)})
}
