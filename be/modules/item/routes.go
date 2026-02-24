package item

import (
	"rialfu/wallet/modules/item/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	controller := do.MustInvoke[controller.ItemController](injector)
	// jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	routes := server.Group("/api/item")
	{
		routes.GET("/", controller.GetAll)
		routes.GET("/option/merk", controller.GetDropdownMerk)
		routes.GET("/option/category", controller.GetDropdownCategory)
		routes.GET("/option/attribute-name", controller.GetDropdownAttributeName)
		routes.GET("/option/attribute-value/:parent", controller.GetDropdownAttributeValue)
		routes.POST("/", controller.Create)
		routes.PUT("/", controller.Update)
		routes.GET("/:id/item", controller.GetInformationItem)
		routes.GET("/:id/variant", controller.GetInformationVariantItem)
	}
}
