package model

type File struct {
	Name string  `json:"name"`
	Size int64 `json:"size"`
	Path string `json:"path"`
}

func (f *File) TableName() string {
	return "file"
}
