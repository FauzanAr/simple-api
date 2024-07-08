package adminhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/modules/admins"
	adminmodel "simple-api.com/m/src/modules/admins/models"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
)

type AdminHandler struct {
	log logger.Logger
	au admins.Usecase
}

func NewAdminHandlers(log logger.Logger, au admins.Usecase) *AdminHandler {
	return &AdminHandler{
		log: log,
		au: au,
	}
}

func (ah *AdminHandler) Login(c *gin.Context) {
	var request adminmodel.AdminLoginRequest
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&request); err != nil {
		ah.log.Error(ctx, "Error while binding the request", err, nil)
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	result, err := ah.au.Login(ctx, request)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", result, http.StatusOK)
}