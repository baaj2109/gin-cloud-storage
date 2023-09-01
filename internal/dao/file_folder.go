package dao

import (
	"log"
	"strconv"
	"time"

	"github.com/baaj2109/gin-cloud-storage/global"
	"github.com/baaj2109/gin-cloud-storage/model"
)

// create folder by
func CreateFolder(folderName, parentId string, fileStoreId int) {
	parentIdInt, err := strconv.Atoi(parentId)
	if err != nil {
		log.Printf("error parentId: %v", err)
		return
	}
	newFolder := model.FileFolder{
		FileFolderName: folderName,
		ParentFolderId: parentIdInt,
		FileStoreId:    fileStoreId,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
	}
	global.MySqlDB.Create(&newFolder)
}

// get parent folder by fid
func GetParentFolder(fid string) (parentFolder model.FileFolder) {
	global.MySqlDB.Where("id = ?", fid).First(&parentFolder)
	return
}

// get all folder by parentid and file store id
func GetAllFolder(parentId string, fileStoreId int) (folders []model.FileFolder) {
	global.MySqlDB.Where("parent_folder_id = ? and file_store_id = ?",
		parentId, fileStoreId).Order("time desc").Find(&folders)
	return
}

// get current folder by fid
func GetCurrentFolder(fid string) (folder model.FileFolder) {
	global.MySqlDB.Where("id = ?", fid).First(&folder)
	return
}

// get all parent folder by current folder
func GetAllParentFolder(currentFolder model.FileFolder, folders []model.FileFolder) []model.FileFolder {
	var parentFolder model.FileFolder
	if currentFolder.ParentFolderId != 0 {
		global.MySqlDB.Where("id = ?", currentFolder.ParentFolderId).First(&parentFolder)
		folders = append(folders, parentFolder)
		return GetAllParentFolder(parentFolder, folders)
	}
	for i, j := 0, len(folders)-1; i < j; i, j = i+1, j-1 {
		folders[i], folders[j] = folders[j], folders[i]
	}
	return folders
}

// get folder count under current folder
func GetFolderCount(fileStoreId int) (folderCount int64) {
	global.MySqlDB.Where("file_store_id = ?", fileStoreId).Count(&folderCount)
	return
}

// update folder name
func UpdateFolderName(fid string, folderName string) {
	global.MySqlDB.Model(&model.FileFolder{}).Where("id = ?", fid).Update("file_folder_name", folderName)
}

// delete folder info
func DeleteFolder(fid string) bool {

	// delete folder
	global.MySqlDB.Where("id = ?", fid).Delete(&model.FileFolder{})

	// delete file under folder
	global.MySqlDB.Where("parent_folder_id = ?", fid).Delete(&model.File{})

	// // delete folder under folder
	// global.MySqlDB.Where("parent_folder_id = ?", fid).Delete(&model.FileFolder{})

	var childFolder []model.FileFolder
	global.MySqlDB.Where("parent_folder_id = ?", fid).Find(&childFolder)
	for _, folder := range childFolder {
		return DeleteFolder(strconv.Itoa(folder.Id))
	}
	return true
}
