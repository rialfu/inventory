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

		routes.POST("/attribute-name", controller.CreateAttributeName)
		routes.PUT("/attribute-name/:id", controller.UpdateAttributeName)
		routes.POST("/attribute-value", controller.CreateAttributeValue)
		routes.PUT("/attribute-value/:id", controller.UpdateAttributeValue)
	}
}
