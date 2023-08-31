package dao

import (
	"github.com/baaj2109/gin-cloud-storage/global"
	"github.com/baaj2109/gin-cloud-storage/model"
)

// get file store infor by user id
func GetFileStoreByUserId(userId int) (fileStore model.FileStore) {
	global.MySqlDB.Where("user_id = ?", userId).Find(&fileStore)
	return
}

// check user capacity is enough
func CheckUserCapacity(fileSize int64, fileStoreId int) bool {
	var fileStore model.FileStore
	global.MySqlDB.Where("id = ?", fileStoreId).Find(&fileStore)
	return fileStore.MaxSize-fileSize/1024 > 0
}
