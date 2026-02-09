package controller

import (
	"errors"
	"net/http"

	"rialfu/wallet/modules/merk/dto"
	"rialfu/wallet/modules/merk/service"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"
	"rialfu/wallet/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/do"
)

type (
	MerkController interface {
		GetAllData(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Get(ctx *gin.Context)
	}

	merkController struct {
		service service.MerkService
		// userValidation *validation.UserValidation
		// db             *gorm.DB
	}
)

func NewMerkController(injector *do.Injector, s service.MerkService) MerkController {
	// db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	// userValidation := validation.NewUserValidation()
	return &merkController{
		service: s,
		// userValidation: userValidation,
	}
}

func (c *merkController) GetAllData(ctx *gin.Context) {
	data, err := c.service.GetAll(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_LIST_DATA, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_LIST_DATA, data)
	ctx.JSON(http.StatusOK, res)
}

func (c *merkController) Get(ctx *gin.Context) {
	var res utils.Response
	var status int
	id := ctx.Param("id")
	if id == "" {
		res = utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_DATA, dto.MerkResponse{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}
	data, err := c.service.Get(ctx, id)
	if err != nil {
		if errors.Is(err, constants.ErrDataNotFound) {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, constants.MESSAGE_FAILED_DATA_NOT_FOUND, data)
			status = 404
		} else {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, err.Error(), nil)
			status = 500
		}

		ctx.JSON(status, res)
		return
	}
	res = utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_DATA, data)
	ctx.JSON(http.StatusOK, res)
}

func (c *merkController) Create(ctx *gin.Context) {
	var req dto.MerkCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var validationMessages = map[string]map[string]string{}
		errs := helpers.TranslateValidationError(err, validationMessages)
		res := utils.BuildResponseFailedValidation(errs, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// id := ctx.Param("id")
	result, err := c.service.Create(ctx.Request.Context(), req)
	if err != nil {
		var res utils.Response
		var status int
		if errors.Is(err, constants.ErrValueNotUniq) {
			res = utils.BuildResponseFailedValidation(map[string]string{"name": constants.MESSAGE_FAILED_VALUE_NOT_UNIQUE}, nil)
			status = http.StatusBadRequest
		} else {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_CREATE_DATA, err.Error(), nil)
			status = http.StatusInternalServerError
		}

		ctx.AbortWithStatusJSON(status, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_CREATE_DATA, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *merkController) Update(ctx *gin.Context) {
	var req dto.MerkUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var validationMessages = map[string]map[string]string{}
		errs := helpers.TranslateValidationError(err, validationMessages)
		res := utils.BuildResponseFailedValidation(errs, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	id := ctx.Param("id")
	result, err := c.service.Update(ctx.Request.Context(), req, id)
	if err != nil {
		var res utils.Response
		var status int
		if errors.Is(err, constants.ErrDataNotFound) {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_UPDATE_DATA, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
			status = http.StatusNotFound
		} else if errors.Is(err, constants.ErrValueNotUniq) {
			res = utils.BuildResponseFailedValidation(map[string]string{"name": constants.MESSAGE_FAILED_VALUE_NOT_UNIQUE}, nil)
			status = http.StatusBadRequest
		} else {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_UPDATE_DATA, err.Error(), nil)
			status = http.StatusInternalServerError
		}

		ctx.AbortWithStatusJSON(status, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_UPDATE_DATA, result)
	ctx.JSON(http.StatusOK, res)
}
