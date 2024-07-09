package userhandler

import (
	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/pkg/middleware"
)

func (uh *UserHandler) UserRoutes(router *gin.RouterGroup) {
	router.POST("/v1/login", uh.Login)

	// Protected
	protectedRoutes := router.Group("")

	protectedRoutes.Use(middleware.GinAuthMiddleware(uh.log))

	protectedRoutes.GET("/v1/profile", uh.GetUserDetail)
	protectedRoutes.PUT("/v1/profile", uh.UpdateUser)

	// Admin
	adminRoutes := router.Group("")

	adminRoutes.POST("/v1/register", middleware.GinAuthAdminMiddleware(uh.log, []string{"admin", "agent"}), uh.Register)
	adminRoutes.GET("/v1/users/", middleware.GinAuthAdminMiddleware(uh.log, []string{"admin"}), uh.GetAllUser)
	adminRoutes.GET("/v1/users/:id", middleware.GinAuthAdminMiddleware(uh.log, []string{"admin", "agent"}), uh.GetUserDetailAdmin)
	adminRoutes.PUT("/v1/users/:id", middleware.GinAuthAdminMiddleware(uh.log, []string{"admin", "agent"}), uh.UpdateUserByAdmin)
}
