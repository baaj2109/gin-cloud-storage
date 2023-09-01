package controller

import (
	"net/http"

	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/gin-gonic/gin"
)

func OtherFiles(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := dao.GetUserInfo(openId.(string))

	fileDetailUse := dao.GetFileUseDetail(user.FileStoreId)
	otherFiles := dao.GetTypeFile(5, user.FileStoreId)

	c.HTML(http.StatusOK, "other-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"otherFiles":    otherFiles,
		"otherCount":    len(otherFiles),
		"currOther":     "active",
		"currClass":     "active",
	})
}
