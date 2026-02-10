package attribute

import (
	"rialfu/wallet/modules/attribute/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	controller := do.MustInvoke[controller.AttributeController](injector)
	// jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	routes := server.Group("/api/attribute")
	{
		routes.GET("/", controller.GetAllDataAttributeName)
		routes.POST("/", controller.CreateAttributeName)
		routes.PUT("/:parent", controller.UpdateAttributeName)
		routes.GET("/:parent", controller.GetAttributeName)
		routes.POST("/:parent/attribute-value", controller.CreateAttributeValue)
		routes.PUT("/:parent/attribute-value/", controller.UpdateAttributeValue)
	}
}
