package matter

import (
	"strconv"

	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
		handler.SendBadRequest(c, errno.ErrParam, nil, err.Error())
		return
	}

	// 一个 bug，当 id=0 时，DB不存在但是 gorm 不会报错
	// 但是 >0 且不存在时，gorm 会报错 record not found
	if id <= 0 {
		handler.SendBadRequest(c, errno.ErrParam, nil, "id is wrong")
		return
	}

	matter, err := model.GetMatterByID(uint32(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			handler.SendBadRequest(c, errno.ErrParam, nil, "id is wrong; the mater does not exist")
			return
		}
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
