package controller

import (
	"errors"
	"net/http"
	"strconv"

	"rialfu/wallet/modules/item/dto"
	"rialfu/wallet/modules/item/service"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"
	"rialfu/wallet/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/do"
)

type (
	ItemController interface {
		GetAll(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		GetDropdownCategory(ctx *gin.Context)
		GetDropdownMerk(ctx *gin.Context)
		GetDropdownAttributeName(ctx *gin.Context)
		GetDropdownAttributeValue(ctx *gin.Context)
		GetInformationItem(ctx *gin.Context)
		GetInformationVariantItem(ctx *gin.Context)
	}

	itemController struct {
		service service.ItemService
		// userValidation *validation.UserValidation
		// db             *gorm.DB
	}
)

func NewItemController(injector *do.Injector, s service.ItemService) ItemController {
	return &itemController{
		service: s,
	}
}

func (c *itemController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_LIST_DATA, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_LIST_DATA, data)
	ctx.JSON(http.StatusOK, res)
}

func (c *itemController) Create(ctx *gin.Context) {
	var req dto.ItemCreateRequest
	var res utils.Response
	status := 200
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
	req.ID = nil

	result, err := c.service.CreateOrUpItem(ctx.Request.Context(), req)
	if err != nil {
		if errors.Is(err, dto.ErrCategoryNotFound) {
			res = utils.BuildResponseFailedValidation(map[string]string{"Category": constants.MESSAGE_FAILED_DATA_NOT_FOUND}, nil)
			status = http.StatusBadRequest
		} else if errors.Is(err, dto.ErrMerkNotFound) {
			res = utils.BuildResponseFailedValidation(map[string]string{"Merk": constants.MESSAGE_FAILED_DATA_NOT_FOUND}, nil)
			status = http.StatusBadRequest
		} else {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_CREATE_DATA, err.Error(), nil)
			status = http.StatusInternalServerError
		}
		ctx.AbortWithStatusJSON(status, res)
		return
	}

	res = utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_CREATE_DATA, result)
	ctx.JSON(status, res)
}
func (c *itemController) CreateVariant(ctx *gin.Context) {
	var req dto.ItemVariantCreateRequest
	var res utils.Response
	status := 200
	if err := ctx.ShouldBind(&req); err != nil {
		res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
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
	res = utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_CREATE_DATA, nil)
	ctx.JSON(status, res)
}
func (c *itemController) Update(ctx *gin.Context) {
	var req dto.ItemCreateRequest
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
	req.ID = &id
	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_UPDATE_DATA, req)
	ctx.JSON(http.StatusOK, res)
}
func (c *itemController) GetDropdownMerk(ctx *gin.Context) {
	search := ctx.Query("search")
	pageString := ctx.Query("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		page = 1
	}
	data, err := c.service.GetDropdownMerk(ctx, search, page)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_LIST_DATA, err.Error(), nil)
		ctx.AbortWithStatusJSON(500, res)
		return
	}
	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_LIST_DATA, data)
	ctx.JSON(200, res)
}
func (c *itemController) GetDropdownCategory(ctx *gin.Context) {
	search := ctx.Query("search")
	pageString := ctx.Query("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		page = 1
	}
	data, err := c.service.GetDropdownCategory(ctx, search, page)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_LIST_DATA, err.Error(), nil)
		ctx.AbortWithStatusJSON(500, res)
		return
	}
	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_LIST_DATA, data)
	ctx.JSON(200, res)
}

func (c *itemController) GetDropdownAttributeName(ctx *gin.Context) {
	search := ctx.Query("search")
	pageString := ctx.Query("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		page = 1
	}
	data, err := c.service.GetDropdownAttributeName(ctx, search, page)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_LIST_DATA, err.Error(), nil)
		ctx.AbortWithStatusJSON(500, res)
		return
	}
	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_LIST_DATA, data)
	ctx.JSON(200, res)
}

func (c *itemController) GetDropdownAttributeValue(ctx *gin.Context) {
	search := ctx.Query("search")
	pageString := ctx.Query("page")
	parent := ctx.Param("parent")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		page = 1
	}
	data, err := c.service.GetDropdownAttributeValueBasedParent(ctx, search, page, parent)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_LIST_DATA, err.Error(), nil)
		ctx.AbortWithStatusJSON(500, res)
		return
	}
	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_LIST_DATA, data)
	ctx.JSON(200, res)
}
func (c *itemController) GetInformationItem(ctx *gin.Context) {
	id := ctx.Param("id")
	var res utils.Response
	status := http.StatusOK
	if id == "" {
		res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
		status = 400
		ctx.AbortWithStatusJSON(status, res)
		return
	}
	data, err := c.service.GetItem(ctx, id)
	if err != nil {
		if errors.Is(err, constants.ErrDataNotFound) {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
			status = 404
		} else {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, err.Error(), nil)
			status = 500
		}
		ctx.AbortWithStatusJSON(status, res)
		return
	}
	res = utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_DATA, data)
	ctx.JSON(status, res)
}
func (c *itemController) GetInformationVariantItem(ctx *gin.Context) {
	id := ctx.Param("id")
	var res utils.Response
	status := http.StatusOK
	if id == "" {
		res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
		status = 400
		ctx.AbortWithStatusJSON(status, res)
		return
	}
	data, err := c.service.GetItemVariants(ctx, id)
	if err != nil {
		res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, err.Error(), nil)
		status = 500
		ctx.AbortWithStatusJSON(status, res)
		return
	}
	res = utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_DATA, data)
	ctx.JSON(status, res)
}
