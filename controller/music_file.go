package controller

import (
	"net/http"

	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/gin-gonic/gin"
)

func MusicFiles(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := dao.GetUserInfo(openId.(string))

	fileDetailUse := dao.GetFileUseDetail(user.FileStoreId)

	musicFiles := dao.GetTypeFile(4, user.FileStoreId)

	c.HTML(http.StatusOK, "music-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"musicFiles":    musicFiles,
		"musicCount":    len(musicFiles),
		"currMusic":     "active",
		"currClass":     "active",
	})
}
