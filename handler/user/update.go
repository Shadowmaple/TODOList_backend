package user

import (
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Nickname string `json:"nickname"`
}

// @Summery 修改用户信息
// @Tags user
// @param data body user.UpdateRequest true "body data"
// @Success 200 string "OK"
// @Router /user [get]
func Update(c *gin.Context) {

}
