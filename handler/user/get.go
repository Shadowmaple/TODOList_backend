package user

import (
	"github.com/gin-gonic/gin"
)

type GetResponse struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	QQ       string `json:"qq"`
}

// Get gets an user by the user identifier.
// @Summary 获取用户信息
// @Tags user
// @Param Authorization header string true "token"
// @Success 200 {object} user.GetResponse
// @Router /user [get]
func Get(c *gin.Context) {

}
