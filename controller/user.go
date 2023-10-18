package controller

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/baaj2109/gin-cloud-storage/global"
	"github.com/baaj2109/gin-cloud-storage/internal"
	"github.com/baaj2109/gin-cloud-storage/internal/dao"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Logout(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		log.Printf("failed to get token from cookie: %v", err)
	}
	// remove from redis
	if err := global.RedisDB.DelKey(token); err != nil {
		log.Printf("failed to remove token from redis: %v", err)
	}
	c.SetCookie("token", "", -1, "/", "", false, false)
	c.Redirect(http.StatusOK, "/")
}

var endpoint = oauth2.Endpoint{
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://oauth2.googleapis.com/token",
}

var googleOauthConfig = &oauth2.Config{
	ClientID:     "368713028639-232vujcid04q68dk0umdhjk60s84v774.apps.googleusercontent.com",
	ClientSecret: "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",
	// RedirectURL:  "http://localhost:8081/oauth/GoogleCallBack",
	RedirectURL: "http://localhost:8080/google_callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
		"Openid",
	},
	Endpoint: endpoint,
}

const oauthStateString = "random"

func HandleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusMovedPermanently, url)
}

func HandleGoogleCallBack(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid session state: %s", oauthStateString))
		return
	}
	code := c.Query("code")
	var err error
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	/*
	   在 Google OAuth 2.0 官方文件有提到，在拿到回應之後，
	   最好把 access_token 和 refresh_token 都保存在一個長期有效的安全地方，
	   例如有人會存在 shared session。
	   refresh_token 這組 token 是當 access_token 過期時，
	   可以拿 refresh_token 再去交換一組有效的 access_token。
	*/
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer resp.Body.Close()
	usefInfoData, _ := io.ReadAll(resp.Body)
	log.Println("Resp body: ", string(usefInfoData))

	userJsonData, err := simplejson.NewJson(usefInfoData)
	if err != nil {
		panic(err.Error())
	}

	openId := userJsonData.Get("sub").MustString()

	hashToken := internal.EncodeMd5("token" + string(rune(time.Now().Unix())) + openId)

	// save to redis
	if err := global.RedisDB.SetKey(hashToken, openId, 3600); err != nil {
		log.Printf("failed to save token to redis: %v", err)
		return
	}
	// set cookie
	c.SetCookie("token", hashToken, 3600, "/", "", false, false)

	// check user exist
	if ok := dao.IsUserExist(openId); ok {
		c.Redirect(http.StatusMovedPermanently, "/cloud/index")
	} else {
		dao.CreateUser(openId, userJsonData.Get("name").MustString(), userJsonData.Get("picture").MustString())
		c.Redirect(http.StatusMovedPermanently, "/cloud/index")
	}

}
