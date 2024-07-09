package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"simple-api.com/m/src/pkg/helper"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
)

func GinRequestTrace(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		traceID := uuid.New().String()
		ctx := context.WithValue(c.Request.Context(), "trace_id", traceID)
		ctx = context.WithValue(ctx, "url", c.Request.RequestURI)
		ctx = context.WithValue(ctx, "method", c.Request.Method)
		ctx = context.WithValue(ctx, "remote_ip", c.ClientIP())
		c.Request = c.Request.WithContext(ctx)

		log.Info(ctx, "Request started", nil)
		c.Header("X-Trace-ID", traceID)

		c.Next()

		log.Info(ctx, "Request completed", logger.MetaData{
			"status":       c.Writer.Status(),
			"elapsed_time": time.Since(start).Seconds(),
		})
	}
}

func GinAuthMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("auth header missing"), nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := helper.VerifyToken(c.Request.Context(), tokenString, "USER")
		if err != nil {
			wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("invalid token"), nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), "user", claims)
		ctx = context.WithValue(ctx, "role", "USER")
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinAuthAdminMiddleware(log logger.Logger, allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("auth header missing"), nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := helper.VerifyToken(c.Request.Context(), tokenString, "ADMIN")
		if err != nil {
			wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("invalid token"), nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		role := claims.Role
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				ctx := context.WithValue(c.Request.Context(), "user", claims)
				ctx = context.WithValue(ctx, "role", "ADMIN")
				c.Request = c.Request.WithContext(ctx)
				c.Next()
				return
			}
		}

		wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("forbidden access"), nil, http.StatusForbidden)
		c.Abort()
	}
}

func GinMultiroleMiddleware(log logger.Logger, allowedRolesForAdmin []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("auth header missing"), nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claimsUser, err := helper.VerifyToken(c.Request.Context(), tokenString, "USER")
		if err == nil {
			ctx := context.WithValue(c.Request.Context(), "user", claimsUser)
			ctx = context.WithValue(ctx, "role", "USER")
			c.Request = c.Request.WithContext(ctx)
			c.Next()
			return
		}

		claimsAdmin, err := helper.VerifyToken(c.Request.Context(), tokenString, "ADMIN")
		if err == nil {
			role := claimsAdmin.Role
			for _, allowedRole := range allowedRolesForAdmin {
				if role == allowedRole {
					ctx := context.WithValue(c.Request.Context(), "user", claimsAdmin)
					ctx = context.WithValue(ctx, "role", "ADMIN")
					c.Request = c.Request.WithContext(ctx)
					c.Next()
					return
				}
			}

			wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("forbidden access"), nil, http.StatusForbidden)
			c.Abort()
			return
		}

		wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("invalid token"), nil, http.StatusUnauthorized)
		c.Abort()
		return
	}
}
