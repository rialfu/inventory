package main

import (
	"log"
	"os"
	"rialfu/wallet/middlewares"
	"rialfu/wallet/modules/attribute"
	"rialfu/wallet/modules/auth"
	"rialfu/wallet/modules/category"
	"rialfu/wallet/modules/merk"
	"rialfu/wallet/modules/user"
	"rialfu/wallet/providers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samber/do"
)

func args(injector *do.Injector) bool {
	// if len(os.Args) > 1 {
	// 	flag := script.Commands(injector)
	// 	return flag
	// }

	return true
}

func run(server *gin.Engine) {

	server.Static("/assets", "./assets")

	port := os.Getenv("GOLANG_PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "0.0.0.0:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	var (
		injector = do.New()
	)

	providers.RegisterDependencies(injector)

	if !args(injector) {
		return
	}

	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())
	// server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Register module routes
	attribute.RegisterRoutes(server, injector)
	user.RegisterRoutes(server, injector)
	auth.RegisterRoutes(server, injector)
	category.RegisterRoutes(server, injector)
	merk.RegisterRoutes(server, injector)

	run(server)
}
