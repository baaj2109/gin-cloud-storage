package controller

import (
	"net/http"

	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/gin-gonic/gin"
)

func ImageFiles(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := dao.GetUserInfo(openId.(string))

	fileDetailUse := dao.GetFileUseDetail(user.FileStoreId)
	imgFiles := dao.GetTypeFile(2, user.FileStoreId)

	c.HTML(http.StatusOK, "image-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"imgFiles":      imgFiles,
		"imgCount":      len(imgFiles),
		"currImg":       "active",
		"currClass":     "active",
	})
}
