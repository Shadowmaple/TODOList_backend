package router

import (
	"net/http"

	"todolist_backend/handler/matter"
	"todolist_backend/handler/sd"
	"todolist_backend/handler/user"
	"todolist_backend/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// api for authentication functionalities
	g.POST("/login", user.Login)

	u := g.Group("/api/v1/user")
	{
		u.POST("/login", user.Login)
		u.GET("", user.Get)
		u.PUT("", user.Update)
	}

	matterRouter := g.Group("api/v1/matter")
	{
		matterRouter.POST("", matter.Create)
		matterRouter.GET("", matter.List)
		matterRouter.GET("/:id", matter.Get)
		matterRouter.PUT("/:id", matter.Update)
		matterRouter.DELETE("/:id", matter.Delete)
		matterRouter.PUT("/:id/finish", matter.Finish)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
