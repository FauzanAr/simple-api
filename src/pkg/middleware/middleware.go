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
		claims, err := helper.VerifyToken(c.Request.Context(), tokenString)
		if err != nil {
			wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("invalid token"), nil, http.StatusUnauthorized)
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), "user", claims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
