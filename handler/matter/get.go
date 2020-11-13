package matter

import "github.com/gin-gonic/gin"

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

}
