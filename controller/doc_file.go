package controller

import (
	"net/http"

	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/gin-gonic/gin"
)

func DocFiles(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := dao.GetUserInfo(openId.(string))

	fileDetailUse := dao.GetFileUseDetail(user.FileStoreId)

	docFiles := dao.GetTypeFile(1, user.FileStoreId)

	c.HTML(http.StatusOK, "doc-files.html", gin.H{
		"user":          user,
		"fileDetailUse": fileDetailUse,
		"docFiles":      docFiles,
		"docCount":      len(docFiles),
		"currDoc":       "active",
		"currClass":     "active",
	})
}
