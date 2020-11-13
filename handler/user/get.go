package user

import (
	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"

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
	userID := c.MustGet("userID").(uint32)

	user, err := model.GetUserById(userID)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, GetResponse{
		Username: user.Username,
		Nickname: user.Nickname,
		QQ:       user.QQ,
	})
}
