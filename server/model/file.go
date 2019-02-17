package model

import (
	"fmt"
	"github.com/yurakawa/go-image-uploader/server/config"
)


func GetAllFile(f *[]File) (err error) {
	if err = config.DB.Find(f).Error; err != nil  {
		return err
	}
	return nil
}

func AddNewFile(f *File) (err error) {
	if err = config.DB.Create(f).Error; err != nil {
		return err
	}
	return nil
}

func GetOneFile(f *File, id int) (err error) {
	if err := config.DB.Where("id = ?", id).First(f).Error; err != nil {
		return err
	}
	return nil
}

func PutOneFile(f *File, id int) (err error) {
	fmt.Println(f)
	config.DB.Save(f)
	return nil
}

// func DeleteFile(f *File, id int)(err error) {
// 	Config.DB.Where("id = ?", id).Delete(f)
// 	return nil
// }

func DeleteFile(uuid string)(err error) {
	config.DB.Where("path = ?", uuid).Delete(File{})
	return nil
}