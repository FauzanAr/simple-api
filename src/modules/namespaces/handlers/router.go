package namespacehandler

import (
	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/pkg/middleware"
)

func (nh *NamespaceHandler) NamespaceRoutes(router *gin.RouterGroup) {
	adminRoutes := router.Group("")

	multiRoleRoutes := router.Group("")
	
	adminRoutes.GET("/v1/namespaces", middleware.GinAuthAdminMiddleware(nh.log, []string{"admin"}), nh.GetAllNamespaces)
	adminRoutes.DELETE("/v1/namespaces/:id", middleware.GinAuthAdminMiddleware(nh.log, []string{"admin", "agent"}), nh.DeleteNamespace)
	adminRoutes.GET("/v1/namespaces/:id/status", middleware.GinAuthAdminMiddleware(nh.log, []string{"admin"}), nh.GetNameSpaceStatus)

	multiRoleRoutes.POST("v1/namespaces", middleware.GinMultiroleMiddleware(nh.log, []string{"admin", "agent"}), nh.CreateNamespace)
	multiRoleRoutes.GET("/v1/namespaces/:id", middleware.GinMultiroleMiddleware(nh.log, []string{"admin", "agent"}), nh.GetDetailNamespace)
}