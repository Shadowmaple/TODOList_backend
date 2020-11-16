package util

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestID, ok := v.(string); ok {
		return requestID
	}
	return ""
}

func GetTokenExpiredTime() time.Duration {
	day := viper.GetInt("jwt.expired")
	return time.Hour * 24 * time.Duration(day)
}

func GetAppIDAndAppSecret() (string, string) {
	return viper.GetString("qq.app_id"), viper.GetString("qq.app_secret")
}
