package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/modules/users"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
)

type UserHandler struct {
	log logger.Logger
	us  users.Usecase
}

func NewUserHandlers(log logger.Logger, us users.Usecase) *UserHandler {
	return &UserHandler{
		log: log,
		us:  us,
	}
}

func (uh *UserHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	uh.log.Info(ctx, "User Login", nil)
	wrapper.SendSuccessResponse(c, "Success", nil, http.StatusOK)
}
