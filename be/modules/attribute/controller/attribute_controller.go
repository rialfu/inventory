package controller

import (
	"errors"
	"net/http"

	"rialfu/wallet/modules/attribute/dto"
	"rialfu/wallet/modules/attribute/service"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"
	"rialfu/wallet/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/do"
)

type (
	AttributeController interface {
		GetAllDataAttributeName(ctx *gin.Context)
		CreateAttributeName(ctx *gin.Context)
		UpdateAttributeName(ctx *gin.Context)
		CreateAttributeValue(ctx *gin.Context)
		UpdateAttributeValue(ctx *gin.Context)
	}

	attributeController struct {
		service service.AttributeService
		// userValidation *validation.UserValidation
		// db             *gorm.DB
	}
)

func NewAttributeController(injector *do.Injector, s service.AttributeService) AttributeController {
	// db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	// userValidation := validation.NewUserValidation()
	return &attributeController{
		service: s,
		// userValidation: userValidation,
	}
}

func (c *attributeController) GetAllDataAttributeName(ctx *gin.Context) {
	data, err := c.service.GetAllAttributeName(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_LIST_DATA, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_LIST_DATA, data)
	ctx.JSON(http.StatusOK, res)
}

func (c *attributeController) CreateAttributeName(ctx *gin.Context) {
	var req dto.AttributeNameCreateRequest
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
	result, err := c.service.CreateAttributeName(ctx.Request.Context(), req)
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

func (c *attributeController) CreateAttributeValue(ctx *gin.Context) {
	var req dto.AttributeValueCreateRequest
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
	result, err := c.service.CreateAttributeValue(ctx.Request.Context(), req)
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

func (c *attributeController) UpdateAttributeName(ctx *gin.Context) {
	var req dto.AttributeNameUpdateRequest
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
	result, err := c.service.UpdateAttributeName(ctx.Request.Context(), req, id)
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
func (c *attributeController) UpdateAttributeValue(ctx *gin.Context) {
	var req dto.AttributeValueUpdateRequest
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
	result, err := c.service.UpdateAttributeValue(ctx.Request.Context(), req, id)
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
