package matter

import (
	"strconv"

	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	MatterInfo
}

// @Summary 修改事项
// @Tags matter
// @Param Authorization header string true "token"
// @param id path string true "matter id"
// @Param data body matter.UpdateRequest true "body data"
// @Success 200 "OK"
// @Router /matter/{id} [put]
func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	matter, err := model.GetMatterByID(uint32(id))
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	userID := c.MustGet("userID").(uint32)
	if matter.UserID != userID {
		handler.SendResponse(c, errno.ErrPermissionDenied, nil)
		return
	}

	request := &CreateRequest{}
	if err := c.BindJSON(request); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	matter.Title = request.Title
	matter.Content = request.Content
	matter.Priority = request.Priority
	matter.State = request.State
	matter.Date = request.Date
	matter.Time = request.Time

	if err := matter.Update(); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, nil)
}
