package user

import (
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Code string `json:"code"` // 小程序 auth code
}

type LoginResponse struct {
	Token string `json:"token"`
}

// @Summary 登录
// @Tags user
// @Param data body user.LoginRequest true "body data"
// @Success 200 {object} user.LoginResponse
// @Router /user/login [post]
func Login(c *gin.Context) {

}
