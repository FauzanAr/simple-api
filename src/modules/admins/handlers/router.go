package adminhandler

import (
	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/pkg/middleware"
)

func (ah *AdminHandler) AdminRoutes(router *gin.RouterGroup) {
	router.POST("/v1/admin", ah.Login)

	protectedRoutes := router.Group("")
	
	protectedRoutes.Use(middleware.GinAuthAdminMiddleware(ah.log))
}