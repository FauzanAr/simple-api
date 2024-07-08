package adminhandler

import (
	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/modules/admins"
	"simple-api.com/m/src/pkg/logger"
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

}