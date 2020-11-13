package matter

import "github.com/gin-gonic/gin"

type FinishPayload struct {
	State bool `json:"state"` // 是否完成
}

func Finish(c *gin.Context) {

}
