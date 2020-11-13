package matter

import (
	"todolist_backend/handler"
	"todolist_backend/model"
	"todolist_backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	Count int           `json:"count"`
	List  []*MatterInfo `json:"list"`
}

// @Summary 获取所有事项
// @Tags matter
// @Param Authorization header string true "token"
// @Success 200 {object} matter.ListResponse
// @Router /matter [get]
func List(c *gin.Context) {
	userID := c.MustGet("userID").(uint32)

	matters, err := model.ListMatters(userID)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	list := make([]*MatterInfo, 0)
	for _, matter := range matters {
		list = append(list, &MatterInfo{
			Title:    matter.Title,
			Content:  matter.Content,
			Priority: matter.Priority,
			State:    matter.State,
			Date:     matter.Date,
			Time:     matter.Time,
		})
	}

	handler.SendResponse(c, nil, ListResponse{
		Count: len(list),
		List:  list,
	})
}
