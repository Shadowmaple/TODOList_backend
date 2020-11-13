package matter

import "github.com/gin-gonic/gin"

// @Summary 删除事项
// @Tags matter
// @Param Authorization header string true "token"
// @param id path string true "matter id"
// @Success 200 "OK"
// @Router /matter/{id} [delete]
func Delete(c *gin.Context) {

}
