package matter

import (
	"strconv"
	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"
	"todolist_backend/service"

	"github.com/gin-gonic/gin"
)

type GetResponse struct {
	MatterInfo
}

// @Summary 查看事项详情
// @Tags matter
// @Param Authorization header string true "token"
// @param id path string true "matter id"
// @Success 200 {object} matter.GetResponse
// @Router /matter/{id} [get]
func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrParam, nil, err.Error())
		return
	}

	matter, err := model.GetMatterByID(uint32(id))
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	state := matter.State
	if state == 0 && service.JudgeOutDate(matter.Date, matter.Time) {
		state = 2
	}

	response := &GetResponse{
		MatterInfo{
			Title:    matter.Title,
			Content:  matter.Content,
			Priority: matter.Priority,
			State:    state,
			Date:     matter.Date,
			Time:     matter.Time,
		},
	}

	handler.SendResponse(c, nil, response)
}
