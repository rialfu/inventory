package category

import (
	"rialfu/wallet/modules/category/controller"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	controller := do.MustInvoke[controller.CategoryController](injector)
	// jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	routes := server.Group("/api/category")
	{
		routes.GET("/", controller.ReadAll)
		routes.POST("/", controller.Create)
		routes.PUT("/:id", controller.Update)
	}
}
