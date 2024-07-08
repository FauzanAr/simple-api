package adminhandler

import (
	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/pkg/middleware"
)

func (ah *AdminHandler) AdminRoutes(router *gin.RouterGroup) {
	protectedRoutes := router.Group("")
	
	protectedRoutes.Use(middleware.GinAuthAdminMiddleware(ah.log))
	protectedRoutes.POST("/v1/admin", ah.Login)
}