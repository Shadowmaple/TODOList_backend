package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Incorrect API route")
	})

	router := g.Group("api/v1")
	{
		router.GET("", func(c *gin.Context) {})
	}

	return g
}
