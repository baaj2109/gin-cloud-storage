package dao

import (
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/baaj2109/gin-cloud-storage/global"
	"github.com/baaj2109/gin-cloud-storage/internal"
	"github.com/baaj2109/gin-cloud-storage/model"
)

func CreateFile(filename, fileHash string, fileSize int64, fId string, fileStoreId int) {

	fileSuffix := path.Ext(filename)
	filePreFix := filename[0 : len(filename)-len(fileSuffix)]
	fid, _ := strconv.Atoi(fId)

	var sizeStr string
	if fileSize < 1048576 {
		sizeStr = strconv.FormatInt(fileSize/1024, 10) + "KB"
	} else {
		sizeStr = strconv.FormatInt(fileSize/102400, 10) + "MB"
	}

	file := model.File{
		FileName:       filename,
		FileHash:       fileHash,
		FileStoreId:    fileStoreId,
		FilePath:       filePreFix + "." + fileSuffix,
		DownloadNumber: 0,
		UploadTime:     time.Now(),
		ParentFolderId: fid,
		Size:           fileSize / 1024,
		SizeStr:        sizeStr,
		Type:           internal.GetFileTypeInt(filePreFix),
		PostFix:        strings.ToLower(fileSuffix),
	}
	global.MySqlDB.Create(&file)
}

func GetUserFile(parentId string, storeId int) (files []model.File) {
	global.MySqlDB.Where("parent_folder_id = ?", parentId).Where("file_store_id = ?", storeId).Find(&files)
	return
}

// 檔案上傳後減去相應容量
func SubtractSize(size int64, fileStoreId int) {
	var fileStore model.FileStore
	global.MySqlDB.Where("id = ?", fileStoreId).First(&fileStore)
	fileStore.CurrentSize = fileStore.CurrentSize + size/1024
	fileStore.MaxSize = fileStore.MaxSize - size/1024
	global.MySqlDB.Save(&fileStore)
}

// get user file number
func GetUserFileCount(fileStoreId int) (fileCount int64) {
	var files []model.File
	global.MySqlDB.Where("file_store_id = ?", fileStoreId).Find(&files).Count(&fileCount)
	return
}

// get user file use detail
func GetFileUseDetail(fileStoreId int) map[string]int64 {
	var files []model.File
	var (
		docCount   int64
		imageCount int64
		videoCount int64
		musicCount int64
		otherCount int64
	)
	detailMap := make(map[string]int64, 0)
	global.MySqlDB.Where("file_store_id = ? AND type = ?", fileStoreId, 0).Find(&files).Count(&docCount)
	global.MySqlDB.Where("file_store_id = ? AND type = ?", fileStoreId, 1).Find(&files).Count(&imageCount)
	global.MySqlDB.Where("file_store_id = ? AND type = ?", fileStoreId, 2).Find(&files).Count(&videoCount)
	global.MySqlDB.Where("file_store_id = ? AND type = ?", fileStoreId, 3).Find(&files).Count(&musicCount)
	global.MySqlDB.Where("file_store_id = ? AND type = ?", fileStoreId, 4).Find(&files).Count(&otherCount)
	detailMap["doc"] = docCount
	detailMap["image"] = imageCount
	detailMap["video"] = videoCount
	detailMap["music"] = musicCount
	detailMap["other"] = otherCount
	return detailMap
}

// get file with certain type
func GetTypeFile(fileType, fileStoreId int) (files []model.File) {
	global.MySqlDB.Where("file_store_id = ? and type = ?", fileStoreId, fileType).Find(&files)
	return
}

// check is file exist
func IsFileExist(fid, filename string) bool {
	fileSuffix := strings.ToLower(path.Ext(filename))
	filePreFix := filename[0 : len(filename)-len(fileSuffix)]
	var file model.File
	global.MySqlDB.Where("parent_folder_id = ? and file_name = ? and postfix = ?", fid, filePreFix, fileSuffix).Find(&file)
	return file.Size >= 0
}

// check if file is already upload by hash
func FileOssExists(fileHash string) bool {
	var file model.File
	global.MySqlDB.Where("file_hash = ?", fileHash).Find(&file)
	return file.FileHash != ""
}

// get file info by file id
func GetFileInfo(fid string) (file model.File) {
	global.MySqlDB.Where("id = ?", fid).Find(&file)
	return
}

// file download number + 1
func AddDownloadNumber(fid string) {
	var file model.File
	global.MySqlDB.Where("id = ?", fid).Find(&file)
	file.DownloadNumber = file.DownloadNumber + 1
	global.MySqlDB.Save(&file)
}

// delete file

func DeleteUserFile(fId, folderId string, storeId int) {
	global.MySqlDB.Where("id = ? and parent_folder_id = ? and file_store_id = ?", fId, folderId, storeId).Delete(&model.File{})
}
