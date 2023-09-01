package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/baaj2109/gin-cloud-storage/internal"
	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/baaj2109/gin-cloud-storage/model"
	"github.com/gin-gonic/gin"
)

// upload file
func Upload(c *gin.Context) {
	openId, _ := c.Get("openId")
	fId := c.DefaultQuery("fId", "0")
	// get user info
	user := dao.GetUserInfo(openId.(string))
	// get current folder info
	currentFolder := dao.GetCurrentFolder(fId)
	// get all foulder under current folder
	folders := dao.GetAllFolder(fId, user.FileStoreId)
	// get parent folder info
	parentFolder := dao.GetParentFolder(fId)
	// get all parent folder
	allParentFolders := dao.GetAllParentFolder(currentFolder, make([]model.FileFolder, 0))
	// get user info detail
	fileUseDetail := dao.GetFileUseDetail(user.FileStoreId)

	c.HTML(http.StatusOK, "upload.html", gin.H{
		"user":             user,
		"currentUpload":    "active",
		"fId":              currentFolder.Id,
		"fName":            currentFolder.FileFolderName,
		"fileFolders":      folders,
		"allParentFolders": allParentFolders,
		"parentFolder":     parentFolder,
		"fileUseDetail":    fileUseDetail,
	})
}

// handle upload file
func HandleUpload(c *gin.Context) {
	openId, _ := c.Get("openId")
	// get user info
	user := dao.GetUserInfo(openId.(string))
	fId := c.GetHeader("id")
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Printf("failed to upload file")
		return
	}

	// is file already exist
	if ok := dao.IsFileExist(fId, head.Filename); !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 501,
		})
		return
	}

	// check if user have enough space
	if ok := dao.CheckUserCapacity(head.Size, user.FileStoreId); !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 503,
		})
	}
	defer file.Close()

	// file save to local
	wd, _ := os.Getwd()

	localLocation := path.Join(wd, head.Filename)
	newFile, err := os.Create(localLocation)
	if err != nil {
		fmt.Printf("failed to create file")
		return
	}
	defer newFile.Close()
	// copy file to new file
	fileSize, err := io.Copy(newFile, file)
	if err != nil {
		fmt.Printf("failed to copy file")
		return
	}
	// shift point to head
	_, _ = newFile.Seek(0, 0)
	fileHash := internal.GetSHA256HashCode(newFile)

	// // check hash to detemeter file already upload to oss or not
	// if ok := dao.FileOssExists(fileHash); ok {
	// 	// update load to cloud
	// }
	dao.CreateFile(head.Filename, fileHash, fileSize, fId, user.FileStoreId)
	dao.SubtractSize(fileSize/1024, user.FileStoreId)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
