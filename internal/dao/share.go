package dao

import (
	"strings"
	"time"

	"github.com/baaj2109/gin-cloud-storage/global"
	"github.com/baaj2109/gin-cloud-storage/internal"
	"github.com/baaj2109/gin-cloud-storage/model"
)

func CreateShare(code, username string, fid int) string {
	share := model.Share{
		Code:     strings.ToLower(code),
		Username: username,
		FileId:   fid,
		Hash:     internal.EncodeMd5(code + string(time.Now().Unix())),
	}
	global.MySqlDB.Create(&share)
	return share.Hash
}

func GetShareInfo(hash string) (share model.Share) {
	global.MySqlDB.Where("hash = ?", hash).First(&share)
	return
}

func VertifyShareCode(fid, code string) bool {
	var share model.Share
	global.MySqlDB.Where("file_id = ? and code = ?", fid, code).First(&share)
	return share.Id == 0
}
