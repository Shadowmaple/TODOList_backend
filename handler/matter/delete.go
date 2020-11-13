package matter

import (
	"strconv"

	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

// @Summary 删除事项
// @Tags matter
// @Param Authorization header string true "token"
// @param id path string true "matter id"
// @Success 200 "OK"
// @Router /matter/{id} [delete]
func Delete(c *gin.Context) {
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

	if err := model.DeleteMatter(uint32(id)); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, nil)
}
