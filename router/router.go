package router

import (
	"net/http"

	_ "todolist_backend/docs"
	"todolist_backend/handler/matter"
	"todolist_backend/handler/sd"
	"todolist_backend/handler/user"
	"todolist_backend/router/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	g.POST("/api/v1/login", user.Login)

	u := g.Group("/api/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("", user.Get)
		u.PUT("", user.Update)
	}

	matterRouter := g.Group("api/v1/matter")
	matterRouter.Use(middleware.AuthMiddleware())
	{
		matterRouter.POST("", matter.Create)
		matterRouter.GET("", matter.List)
		matterRouter.GET("/:id", matter.Get)
		matterRouter.PUT("/:id", matter.Update)
		matterRouter.DELETE("/:id", matter.Delete)
		matterRouter.PUT("/:id/finish", matter.Finish)
	}

	// swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
