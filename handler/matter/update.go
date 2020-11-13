package matter

import "github.com/gin-gonic/gin"

type UpdateRequest struct {
	MatterInfo
}

// @Summery 修改事项
// @Tags matter
// @Param Authorization header string true "token"
// @param id path string true "matter id"
// @Param data body matter.UpdateRequest true "body data"
// @Success 200 "OK"
// @Router /matter/{id} [put]
func Update(c *gin.Context) {

}
