package model

import "time"

type User struct {
	Id           int       `json:"id"`
	OpenId       string    `json:"open_id"`
	FileStoreId  int       `json:"file_store_id"`
	Username     string    `json:"username"`
	RegisterTime time.Time `json:"register_time"`
	ImagePath    string    `json:"image_path"`
}

// return table name
func (User) TableName() string {
	return "user"
}
