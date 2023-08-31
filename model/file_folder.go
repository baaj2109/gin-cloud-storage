package model

type FileFolder struct {
	Id             int
	FileFolderName string
	ParentFolderId int
	FileStoreId    int
	Time           string
}

// return table name
func (FileFolder) TableName() string {
	return "file_folder"
}
