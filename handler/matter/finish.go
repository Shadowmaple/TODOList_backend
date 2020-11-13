package matter

import (
	"strconv"

	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

type FinishPayload struct {
	State bool `json:"state"` // 是否完成
}

// @Summary 完成/取消完成事项
// @Tags matter
// @Param Authorization header string true "token"
// @param id path string true "matter id"
// @Param data body matter.FinishPayload true "body data"
// @Success 200 {object} matter.FinishPayload
// @Router /matter/{id}/finish [put]
func Finish(c *gin.Context) {
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

	request := &FinishPayload{}
	if err := c.BindJSON(request); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if request.State && matter.State != 1 {
		handler.SendResponse(c, nil, "has not finished yet.")
		return
	} else if !request.State && matter.State == 1 {
		handler.SendResponse(c, nil, "has finished already.")
		return
	}

	matter.State = 1 - matter.State
	if err := matter.Update(); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, FinishPayload{State: !request.State})
}
