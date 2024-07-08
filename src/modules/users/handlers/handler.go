package userhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"simple-api.com/m/src/modules/users"
	usermodel "simple-api.com/m/src/modules/users/model"
	"simple-api.com/m/src/pkg/helper"
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
	var request usermodel.UserLoginRequest
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&request); err != nil {
		uh.log.Error(ctx, "Error while binding the request", err, nil)
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	result, err := uh.us.Login(ctx, request)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", result, http.StatusOK)
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var request usermodel.UserUpdateRequest
	ctx := c.Request.Context()
	user, ok := ctx.Value("user").(*helper.AccessClaims)
	if !ok {
		errMsg := wrapper.InternalServerError("Error while converting request")
		wrapper.SendErrorResponse(c, errMsg, nil, http.StatusInternalServerError)
		return
	}

	request.OriginalUsername = user.Username

	if err := c.ShouldBindJSON(&request); err != nil {
		uh.log.Error(ctx, "Error while binding the request", err, nil)
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	result, err := uh.us.UpdateUser(ctx, request)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", result, http.StatusOK)
}

func (uh *UserHandler) GetUserDetail(c *gin.Context) {
	var request usermodel.UserDetailRequest
	ctx := c.Request.Context()
	user, ok := ctx.Value("user").(*helper.AccessClaims)
	if !ok {
		errMsg := wrapper.InternalServerError("Error while converting request")
		wrapper.SendErrorResponse(c, errMsg, nil, http.StatusInternalServerError)
		return
	}

	request.Username = user.Claims.Username

	result, err := uh.us.GetUserDetail(ctx, request)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", result, http.StatusOK)
}

func (uh *UserHandler) Register(c *gin.Context) {
	var request usermodel.UserRegisterRequest
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&request); err != nil {
		uh.log.Error(ctx, "Error while binding the request", err, nil)
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	err := uh.us.RegisterUser(ctx, request)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", nil, http.StatusCreated)
}

func (uh *UserHandler) GetAllUser(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := uh.us.GetAllUserDetail(ctx)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", result, http.StatusOK)
}

func (uh *UserHandler) GetUserDetailAdmin(c *gin.Context) {
	var req usermodel.UserDetailRequest
	ctx := c.Request.Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		wrapper.SendErrorResponse(c, wrapper.BadRequestError("Invalid userId"), nil, http.StatusBadRequest)
		return
	}

	req.Id = userId

	result, err := uh.us.GetUserDetail(ctx, req)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", result, http.StatusOK)
}

func (uh *UserHandler) UpdateUserByAdmin(c *gin.Context) {
	var request usermodel.UserUpdateRequest
	ctx := c.Request.Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		wrapper.SendErrorResponse(c, wrapper.BadRequestError("Invalid userId"), nil, http.StatusBadRequest)
		return
	}

	request.Id = userId

	if err := c.ShouldBindJSON(&request); err != nil {
		uh.log.Error(ctx, "Error while binding the request", err, nil)
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	result, err := uh.us.UpdateUser(ctx, request)
	if err != nil {
		wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
		return
	}

	wrapper.SendSuccessResponse(c, "Success", result, http.StatusOK)
}
