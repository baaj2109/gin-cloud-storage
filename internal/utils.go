package internal

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"math/rand"
	"os"
	"strings"
)

func GetSHA256HashCode(file *os.File) string {
	// 建立一個sha256算法的hash.Hash接口對象
	hash := sha256.New()
	_, _ = io.Copy(hash, file)
	// 計算 hash值
	bytes := hash.Sum(nil)
	// 將字符算編碼為16進位格式 回傳字串
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

func EncodeMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetFileTypeInt(filePrefix string) int {
	filePrefix = strings.ToLower(filePrefix)
	if filePrefix == ".doc" || filePrefix == ".docx" || filePrefix == ".txt" || filePrefix == ".pdf" {
		return 1
	}
	if filePrefix == ".jpg" || filePrefix == ".png" || filePrefix == ".gif" || filePrefix == ".jpeg" {
		return 2
	}
	if filePrefix == ".mp4" || filePrefix == ".avi" || filePrefix == ".mov" || filePrefix == ".rmvb" || filePrefix == ".rm" {
		return 3
	}
	if filePrefix == ".mp3" || filePrefix == ".cda" || filePrefix == ".wav" || filePrefix == ".wma" || filePrefix == ".ogg" {
		return 4
	}
	return 5
}

func ConvertToMap(str string) map[string]string {
	var resultMap = make(map[string]string)
	values := strings.Split(str, "&")
	for _, value := range values {
		vs := strings.Split(value, "=")
		resultMap[vs[0]] = vs[1]
	}
	return resultMap
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
