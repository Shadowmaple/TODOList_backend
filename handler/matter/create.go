package matter

import "github.com/gin-gonic/gin"

type CreateRequest struct {
	MatterInfo
}

// @Summary 创建事项
// @Tags matter
// @Param Authorization header string true "token"
// @param data body matter.CreateRequest true "body data"
// @Success 200 "OK"
// @Router /matter [post]
func Create(c *gin.Context) {

}
