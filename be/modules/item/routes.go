package item

import (
	"rialfu/wallet/modules/merk/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	controller := do.MustInvoke[controller.MerkController](injector)
	// jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	routes := server.Group("/api/merk")
	{
		routes.POST("/", controller.Create)
		routes.PUT("/:id", controller.Update)
	}
}
