package model

import (
	"time"
)

type File struct {
	Id             int
	FileName       string
	FileHash       string
	FileStoreId    int
	FilePath       string
	DownloadNumber int
	UploadTime     time.Time
	ParentFolderId int
	Size           int64
	SizeStr        string
	Type           int
	PostFix        string
}
