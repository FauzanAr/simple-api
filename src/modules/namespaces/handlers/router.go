package namespacehandler

import (
	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/pkg/middleware"
)

func (nh *NamespaceHandler) NamespaceRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("")
	userRoutes.Use(middleware.GinAuthMiddleware(nh.log))

	adminRoutes := router.Group("")
	adminRoutes.Use(middleware.GinAuthAdminMiddleware(nh.log))

	userRoutes.POST("/v1/namespaces", nh.CreateNamespace)

	adminRoutes.DELETE("/v1/namespaces/:id", nh.DeleteNamespace)
}