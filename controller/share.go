package controller

import (
	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/gin-gonic/gin"
)

// share file
func ShareFile(c *gin.Context) {
	openId, _ := c.Get("openId")

	// get user from openId
	user := dao.GetUserInfo(openId.(string))

	fId := c.Query("id")
	url := c.Query("url")

	// code =
}
