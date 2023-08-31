package model

type FileStore struct {
	Id          int
	UserId      int
	CurrentSize int64
	MaxSize     int64
}

// table name
func (FileStore) TableName() string {
	return "file_store"
}
