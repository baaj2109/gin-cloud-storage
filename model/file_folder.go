package model

import "time"

type FileFolder struct {
	Id             int
	FileFolderName string
	ParentFolderId int
	FileStoreId    int
	Time           time.Time
}

// return table name
func (FileFolder) TableName() string {
	return "file_folder"
}
