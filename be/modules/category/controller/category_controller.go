package controller

import (
	"errors"
	"net/http"
	"rialfu/wallet/modules/category/dto"
	"rialfu/wallet/modules/category/service"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"
	"rialfu/wallet/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/do"
)

type (
	CategoryController interface {
		Update(ctx *gin.Context)
		Create(ctx *gin.Context)
		ReadAll(ctx *gin.Context)
		DropDownValue(ctx *gin.Context)
	}

	categoryController struct {
		service service.CategoryService
	}
)

func NewCategoryController(injector *do.Injector, s service.CategoryService) CategoryController {

	return &categoryController{
		service: s,
	}
}

func (c *categoryController) ReadAll(ctx *gin.Context) {
	val, exist := ctx.GetQuery("parent")
	var parentID *string
	if exist && val != "" {
		parentID = &val
	}
	data, path, err := c.service.GetAll(ctx, parentID, ctx.Request.URL.Query())
	if err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_LIST_DATA, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_LIST_DATA, map[string]any{
		"data":  data.Data,
		"limit": data.Limit,
		"page":  data.Page,
		"total": data.Total,
		"path":  path,
	})
	ctx.JSON(http.StatusOK, res)

}

func (c *categoryController) Update(ctx *gin.Context) {
	var req dto.CategoryUpdateRequest
	var res utils.Response
	status := 200
	id := ctx.Param("id")
	if id == "" {
		res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_UPDATE_DATA, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
		ctx.JSON(404, res)
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	// Validate request

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var validationMessages = map[string]map[string]string{}
		errs := helpers.TranslateValidationError(err, validationMessages)
		res := utils.BuildResponseFailedValidation(errs, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data, err := c.service.Update(ctx, req, id)
	if err != nil {
		if errors.Is(err, constants.ErrDataNotFound) {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_UPDATE_DATA, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
			status = 404
		} else if errors.Is(err, dto.ErrGetParentNotFound) {
			res = utils.BuildResponseFailedValidation(map[string]string{"parent": constants.MESSAGE_FAILED_DATA_NOT_FOUND}, nil)
			status = 400
		} else if errors.Is(err, dto.ErrDataCircular) {
			res = utils.BuildResponseFailedValidation(map[string]string{"parent": err.Error()}, nil)
			status = 400
		} else {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_UPDATE_DATA, err.Error(), nil)
			status = 500
		}

		ctx.JSON(status, res)
		return
	}
	res = utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_UPDATE_DATA, data)
	ctx.JSON(status, res)
}
func (c *categoryController) Create(ctx *gin.Context) {
	var req dto.CategoryCreateRequest
	var res utils.Response
	status := 200
	if err := ctx.ShouldBind(&req); err != nil {
		res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	// Validate request

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var validationMessages = map[string]map[string]string{}
		errs := helpers.TranslateValidationError(err, validationMessages)
		res := utils.BuildResponseFailedValidation(errs, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data, err := c.service.Create(ctx, req)
	if err != nil {
		if errors.Is(err, dto.ErrGetParentNotFound) {
			res = utils.BuildResponseFailedValidation(map[string]string{"parent": constants.MESSAGE_FAILED_DATA_NOT_FOUND}, nil)
			status = 400
		} else {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_UPDATE_DATA, err.Error(), nil)
			status = 500
		}

		ctx.JSON(status, res)
		return
	}
	res = utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_UPDATE_DATA, data)
	ctx.JSON(status, res)
}
func (c *categoryController) DropDownValue(ctx *gin.Context) {
	search := ctx.Query("search")
	pageString := ctx.Query("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		page = 1
	}
	data, err := c.service.GetDropdown(ctx, search, page)
	if err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_UPDATE_DATA, err.Error(), nil)
		ctx.AbortWithStatusJSON(500, res)
		return
	}
	res := utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_LIST_DATA, data)
	ctx.JSON(200, res)
}
