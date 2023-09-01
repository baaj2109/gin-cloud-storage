package controller

import (
	"net/http"

	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	openId, _ := c.Get("openId")
	// get user info
	user := dao.GetUserInfo(openId.(string))
	// get user file store info
	fileStore := dao.GetFileStoreByUserId(user.FileStoreId)
	// get file count
	fileCount := dao.GetFileCount(user.FileStoreId)
	// get folder count
	folderCount := dao.GetFolderCount(user.FileStoreId)
	// get file user detail
	fileUseDetail := dao.GetFileUseDetail(user.FileStoreId)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":          user,
		"currentIndex":  "activate",
		"fileStore":     fileStore,
		"fileCount":     fileCount,
		"folderCount":   folderCount,
		"fileUseDetail": fileUseDetail,
	})
}
