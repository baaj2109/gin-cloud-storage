package middleware

import (
	"log"
	"net/http"

	"github.com/baaj2109/gin-cloud-storage/global"
	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/gin-gonic/gin"
)

func CheckLogin(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		log.Printf("failed get token from cookie: %v", err)
		c.Redirect(http.StatusNotFound, "/")
		c.Abort()
	}

	// get id from redis cache
	id, err := global.RedisDB.GetKey(token)
	if err != nil {
		log.Printf("failed get id from redis: %v", err)
		c.Redirect(http.StatusNotFound, "/")
		c.Abort()
	}

	// get user from database
	user := dao.GetUserInfo(id)
	if user.Id == 0 {
		c.Redirect(http.StatusNotFound, "/")
	}
	c.Set("openId", id)
	c.Next()
}
