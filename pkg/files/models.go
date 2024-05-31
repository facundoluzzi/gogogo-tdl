package files

import "gorm.io/gorm"

type FileRecord struct {
	gorm.Model
	Filename string
	Content  string
}
