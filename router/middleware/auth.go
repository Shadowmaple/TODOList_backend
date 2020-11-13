package middleware

import (
	"todolist_backend/handler"
	"todolist_backend/pkg/errno"
	"todolist_backend/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set("userID", ctx.ID)

		c.Next()
	}
}
