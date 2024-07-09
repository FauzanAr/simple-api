package namespacehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/modules/namespaces"
	namespacemodel "simple-api.com/m/src/modules/namespaces/models"
	"simple-api.com/m/src/pkg/helper"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
)

type NamespaceHandler struct {
	log logger.Logger
	nu  namespaces.Usecase
}

func NewNamespaceHandler(log logger.Logger, nu namespaces.Usecase) *NamespaceHandler {
	return &NamespaceHandler{
		log: log,
		nu:  nu,
	}
}

func (nh *NamespaceHandler) CreateNamespace(c *gin.Context) {
	var request namespacemodel.NamespaceCreateRequest
	ctx := c.Request.Context()
	role := c.Request.Context().Value("role").(string)
	if role == "USER" {
		user, ok := ctx.Value("user").(*helper.AccessClaims)
		if !ok {
			errMsg := wrapper.InternalServerError("Error while converting request")
			wrapper.SendErrorResponse(c, errMsg, nil, http.StatusInternalServerError)
			return
		}

		request.UserID = int(user.Claims.Id)
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		nh.log.Error(ctx, "Error while binding the request", err, nil)
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	err := nh.nu.CreateNamespace(ctx, request)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", nil, http.StatusOK)
}

func (nh *NamespaceHandler) DeleteNamespace(c *gin.Context) {
	var req namespacemodel.NamespaceDeleteRequest
	ctx := c.Request.Context()

	namespaceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		wrapper.SendErrorResponse(c, wrapper.BadRequestError("Invalid namespaceId"), nil, http.StatusBadRequest)
		return
	}

	req.Id = namespaceId

	err = nh.nu.DeleteNamespace(ctx, req)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", nil, http.StatusOK)
}

func (nh *NamespaceHandler) GetAllNamespaces(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := nh.nu.GetAllNamespaces(ctx)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", res, http.StatusOK)
}

func (nh *NamespaceHandler) GetNameSpaceStatus(c *gin.Context) {
	var req namespacemodel.NamespaceGetStatusRequest
	ctx := c.Request.Context()

	namespaceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		wrapper.SendErrorResponse(c, wrapper.BadRequestError("Invalid namespaceId"), nil, http.StatusBadRequest)
		return
	}

	req.Id = namespaceId

	res, err := nh.nu.GetStatusNamespace(ctx, req)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", res, http.StatusOK)
}

func (nh *NamespaceHandler) GetDetailNamespace(c *gin.Context) {
	var req namespacemodel.NamespaceGetDetailRequest
	ctx := c.Request.Context()

	namespaceId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		wrapper.SendErrorResponse(c, wrapper.BadRequestError("Invalid namespaceId"), nil, http.StatusBadRequest)
		return
	}

	user, ok := ctx.Value("user").(*helper.AccessClaims)
	if !ok {
		errMsg := wrapper.InternalServerError("Error while converting request")
		wrapper.SendErrorResponse(c, errMsg, nil, http.StatusInternalServerError)
		return
	}

	req.Id = namespaceId
	req.Role = c.Request.Context().Value("role").(string)
	req.UserId = int(user.Claims.Id)

	res, err := nh.nu.GetDetailNamespace(ctx, req)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", res, http.StatusOK)
}
