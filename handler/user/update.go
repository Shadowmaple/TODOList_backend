package user

import (
	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Nickname string `json:"nickname"`
}

// @Summary 修改用户信息
// @Tags user
// @Param Authorization header string true "token"
// @param data body user.UpdateRequest true "body data"
// @Success 200 "OK"
// @Router /user [get]
func Update(c *gin.Context) {
	userID := c.MustGet("userID").(uint32)

	user, err := model.GetUserById(userID)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	request := UpdateRequest{}
	if err := c.BindJSON(&request); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	user.Nickname = request.Nickname
	if err := user.Update(); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, nil)
}
