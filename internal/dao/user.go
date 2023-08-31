package dao

import (
	"time"

	"github.com/baaj2109/gin-cloud-storage/global"
	"github.com/baaj2109/gin-cloud-storage/model"
)

func CreateUser(openId, username, imageUrl string) {
	user := model.User{
		OpenId:       openId,
		Username:     username,
		FileStoreId:  0,
		ImagePath:    imageUrl,
		RegisterTime: time.Now(),
	}
	// global.MySqlDB.Create(&user)
	fileStore := model.FileStore{
		UserId:      user.Id,
		CurrentSize: 0,
		MaxSize:     1048576,
	}
	global.MySqlDB.Create(&fileStore)
	user.FileStoreId = fileStore.Id
	global.MySqlDB.Create(&user)
}

func GetUserInfo(openId string) (user model.User) {
	global.MySqlDB.Find(&user, "open_id = ?", openId)
	return
}

func IsUserExist(openId string) bool {
	var user model.User
	global.MySqlDB.Find(&user, "open_id = ?", openId)
	return user.Id != 0
}
