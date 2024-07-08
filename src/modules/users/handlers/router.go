package userhandler

import (
	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/pkg/middleware"
)

func (uh *UserHandler) UserRoutes(router *gin.RouterGroup) {
	router.POST("/v1/login", uh.Login)
	router.POST("/v1/register", uh.Register)

	// Protected
	protectedRoutes := router.Group("")

	protectedRoutes.Use(middleware.GinAuthMiddleware(uh.log))

	protectedRoutes.GET("/v1/profile", uh.GetUserDetail)
	protectedRoutes.PUT("/v1/profile", uh.UpdateUser)
}