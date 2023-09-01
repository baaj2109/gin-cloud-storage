package controller

import (
	"net/http"
	"strconv"

	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/baaj2109/gin-cloud-storage/model"
	"github.com/gin-gonic/gin"
)

func Files(c *gin.Context) {
	openId, _ := c.Get("openId")
	fId := c.DefaultQuery("fId", "0")
	// get user info
	user := dao.GetUserInfo(openId.(string))

	// get all files
	files := dao.GetUserFile(fId, user.FileStoreId)
	// get all foulder under current folder
	folders := dao.GetAllFolder(fId, user.FileStoreId)
	// get parent folder info
	parentFolder := dao.GetParentFolder(fId)

	// get current folder info
	currentFolder := dao.GetCurrentFolder(fId)
	// get all parent folder
	allParentFolders := dao.GetAllParentFolder(currentFolder, make([]model.FileFolder, 0))
	// get user info detail
	fileUseDetail := dao.GetFileUseDetail(user.FileStoreId)

	c.HTML(http.StatusOK, "files.html", gin.H{
		"user":             user,
		"currAll":          "active",
		"files":            files,
		"fId":              currentFolder.Id,
		"fName":            currentFolder.FileFolderName,
		"fileFolders":      folders,
		"allParentFolders": allParentFolders,
		"parentFolder":     parentFolder,
		"fileUseDetail":    fileUseDetail,
	})
}

// add folder
func AddFolder(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := dao.GetUserInfo(openId.(string))
	folderName := c.PostForm("fileFolderName")
	parentId := c.DefaultPostForm("parentFolderId", "0")
	dao.CreateFolder(folderName, parentId, user.FileStoreId)
	// get parent folder
	parentFolder := dao.GetParentFolder(parentId)
	c.Redirect(
		http.StatusMovedPermanently,
		"/cloud/files?fId="+parentId+"&fName="+parentFolder.FileFolderName,
	)
}

func DownloadFile(c *gin.Context) {
	fId := c.Query("fId")
	file := dao.GetFileInfo(fId)
	if file.FileHash == "" {
		return
	}
	// download from oss
	// file data

	dao.AddDownloadNumber(fId)
	c.Header(
		"Content-disposition",
		"attachment;filename=\""+file.FileName+file.PostFix+"\"",
	)

	c.Data(
		http.StatusOK,
		"application/octect-stream",
		make([]byte, 10), // fileData,
	)
}

func DeleteFile(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := dao.GetUserInfo(openId.(string))
	fId := c.DefaultQuery("fId", "0")
	folderId := c.Query("folder")
	if fId == "" {
		return
	}
	dao.DeleteUserFile(fId, folderId, user.FileStoreId)
	c.Redirect(http.StatusMovedPermanently, "/cloud/files?fid="+folderId)
}

func DeleteFileFolder(c *gin.Context) {
	fId := c.DefaultQuery("fId", "")
	if fId == "" {
		return
	}
	folderInfo := dao.GetCurrentFolder(fId)
	c.Redirect(
		http.StatusMovedPermanently,
		"/cloud/files?fId="+strconv.Itoa(folderInfo.ParentFolderId),
	)
}

func UpdateFileFolder(c *gin.Context) {
	fileFolderName := c.PostForm("fileFolderName")
	fileFolderId := c.PostForm("fileFolderId")
	fileFolder := dao.GetCurrentFolder(fileFolderId)
	dao.UpdateFolderName(fileFolderId, fileFolderName)
	c.Redirect(
		http.StatusMovedPermanently,
		"/cloud/files?fId="+strconv.Itoa(fileFolder.ParentFolderId),
	)
}
