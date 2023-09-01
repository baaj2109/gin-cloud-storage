package controller

import (
	"net/http"

	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/gin-gonic/gin"
)

func VideoFiles(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := dao.GetUserInfo(openId.(string))

	fileDetailUse := dao.GetFileUseDetail(user.FileStoreId)
	videoFiles := dao.GetTypeFile(3, user.FileStoreId)

	c.HTML(http.StatusOK, "video-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"videoFiles":    videoFiles,
		"videoCount":    len(videoFiles),
		"currVideo":     "active",
		"currClass":     "active",
	})
}
