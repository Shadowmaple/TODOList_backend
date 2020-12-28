package matter

import (
	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	MatterInfo
}

// @Summary 创建事项
// @Tags matter
// @Param Authorization header string true "token"
// @param data body matter.CreateRequest true "body data"
// @Success 200 "OK"
// @Router /matter [post]
func Create(c *gin.Context) {
	request := &CreateRequest{}
	if err := c.BindJSON(request); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userID := c.MustGet("userID").(uint32)

	matter := &model.MatterModel{
		Title:    request.Title,
		Content:  request.Content,
		Priority: request.Priority,
		State:    0,
		Date:     request.Date,
		Time:     request.Time,
		UserID:   userID,
	}

	if err := matter.Create(); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, nil)
}
