package matter

import "github.com/gin-gonic/gin"

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

}
