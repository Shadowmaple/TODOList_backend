package matter

import "github.com/gin-gonic/gin"

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

}
