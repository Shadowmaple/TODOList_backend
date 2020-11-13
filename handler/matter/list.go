package matter

import "github.com/gin-gonic/gin"

type ListResponse struct {
	Count int           `json:"count"`
	List  []*MatterInfo `json:"list"`
}

func List(c *gin.Context) {

}
