package controller

import (
	"net/http"
	"strconv"

	"github.com/baaj2109/gin-cloud-storage/internal"
	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/gin-gonic/gin"
)

var codeSize = 4

// share file
func ShareFile(c *gin.Context) {
	openId, _ := c.Get("openId")

	// get user from openId
	user := dao.GetUserInfo(openId.(string))

	fId := c.Query("id")
	url := c.Query("url")

	code := internal.RandStringRunes(codeSize)
	filedId, _ := strconv.Atoi(fId)
	hash := dao.CreateShare(code, user.Username, filedId)
	c.JSON(http.StatusOK, gin.H{
		"url":  url + "?=" + hash,
		"code": code,
	})
}

// share pass
func SharePass(c *gin.Context) {
	f := c.Query("f")
	// get share info
	info := dao.GetShareInfo(f)
	// get file info
	file := dao.GetFileInfo(strconv.Itoa(info.FileId))
	c.HTML(http.StatusOK, "share.html", gin.H{
		"id":       info.FileId,
		"username": info.Username,
		"fileType": file.Type,
		"filename": file.FileName + "." + file.PostFix,
		"hash":     info.Hash,
	})
}

// download shard file
func DownloadShareFile(c *gin.Context) {
	// fileId := c.Query("id")
	// code := c.Query("code")
	// hash := c.Query("hash")
	// info = dao.GetFileInfo(fileId)

	// if ok := dao.VertifyShareCode(fileId, code); !ok {
	// 	c.Redirect(http.StatusMovedPermanently, "/file/share?f="+hash)
	// 	return
	// }

	// // get file from oss
	// // fileData := dao.
	// c.SaveUploadedFile()
}
