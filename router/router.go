package router

import (
	"github.com/baaj2109/gin-cloud-storage/controller"
	"github.com/baaj2109/gin-cloud-storage/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()

	router.GET("/", controller.Login)
	// router ping pong
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// google login
	router.GET("/google_login", controller.HandleGoogleLogin)
	router.GET("/google_callback", controller.HandleGoogleCallBack)
	// router.GET("/oauth/GoogleCallBack", controller.HandleGoogleCallBack)

	router.GET("/file/share", controller.SharePass)
	// router.GET("/file/shareDownload", controller.DownloadShareFile)

	cloud := router.Group("cloud")
	cloud.Use(middleware.CheckLogin)
	{
		cloud.GET("/index", controller.Index)
		cloud.GET("/files", controller.Files)
		cloud.GET("/upload", controller.Upload)
		cloud.GET("/doc-files", controller.DocFiles)
		cloud.GET("/image-files", controller.ImageFiles)
		cloud.GET("/video-files", controller.VideoFiles)
		cloud.GET("/music-files", controller.MusicFiles)
		cloud.GET("/other-files", controller.OtherFiles)
		cloud.GET("/logout", controller.Logout)
		cloud.GET("/downloadFile", controller.DownloadFile)
		cloud.GET("/deleteFile", controller.DeleteFile)
		cloud.GET("/deleteFolder", controller.DeleteFileFolder)
		cloud.GET("/help", controller.Help)
	}

	{
		cloud.POST("/uploadFile", controller.HandleUpload)
		cloud.POST("/addFolder", controller.AddFolder)
		cloud.POST("/updateFolder", controller.UpdateFileFolder)
		cloud.POST("/getQrCode", controller.ShareFile)
	}

	return router
}
