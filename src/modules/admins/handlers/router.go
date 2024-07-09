package adminhandler

import (
	"github.com/gin-gonic/gin"
)

func (ah *AdminHandler) AdminRoutes(router *gin.RouterGroup) {
	router.POST("/v1/admin", ah.Login)
}