package userhandler

import "github.com/gin-gonic/gin"

func (uh *UserHandler) UserRoutes(router *gin.RouterGroup) {
	router.POST("/v1/login", uh.Login)
}