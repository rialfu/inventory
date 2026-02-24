package providers

import (
	"rialfu/wallet/config"
	attrController "rialfu/wallet/modules/attribute/controller"
	attrRepo "rialfu/wallet/modules/attribute/repository"
	attrService "rialfu/wallet/modules/attribute/service"
	authController "rialfu/wallet/modules/auth/controller"
	authService "rialfu/wallet/modules/auth/service"
	cateController "rialfu/wallet/modules/category/controller"
	cateRepo "rialfu/wallet/modules/category/repository"
	cateService "rialfu/wallet/modules/category/service"
	merkController "rialfu/wallet/modules/merk/controller"
	merkRepo "rialfu/wallet/modules/merk/repository"
	merkService "rialfu/wallet/modules/merk/service"
	userController "rialfu/wallet/modules/user/controller"
	userRepo "rialfu/wallet/modules/user/repository"
	userService "rialfu/wallet/modules/user/service"

	itemController "rialfu/wallet/modules/item/controller"
	itemRepo "rialfu/wallet/modules/item/repository"
	itemService "rialfu/wallet/modules/item/service"
	"rialfu/wallet/pkg/constants"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func(i *do.Injector) (*gorm.DB, error) {
		return config.SetUpTestDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (authService.JWTService, error) {
		return authService.NewJWTService(), nil
	})

	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[authService.JWTService](injector, constants.JWTService)

	userRepository := userRepo.NewUserRepository(db)
	catRepo := cateRepo.NewCategoryRepository(db)
	mr := merkRepo.NewMerkRepository(db)
	anr := attrRepo.NewAttributeNameRepository(db)
	avr := attrRepo.NewAttributeValueRepository(db)
	ir := itemRepo.NewItemRepository(db)
	ivr := itemRepo.NewItemVariantRepository(db)

	userService := userService.NewUserService(userRepository, db)
	authService := authService.NewAuthService(userRepository, jwtService, db)
	catService := cateService.NewCategoryService(catRepo, db)
	ms := merkService.NewMerkService(mr, db)
	as := attrService.NewAttributeService(anr, avr, db)
	is := itemService.NewItemService(ir, ivr, mr, catRepo, avr, anr, db)
	do.Provide(injector, func(i *do.Injector) (userController.UserController, error) {
		return userController.NewUserController(i, userService), nil
	})

	do.Provide(injector, func(i *do.Injector) (authController.AuthController, error) {
		return authController.NewAuthController(i, authService), nil
	})

	do.Provide(injector, func(i *do.Injector) (cateController.CategoryController, error) {
		return cateController.NewCategoryController(i, catService), nil
	})
	do.Provide(injector, func(i *do.Injector) (attrController.AttributeController, error) {
		return attrController.NewAttributeController(i, as), nil
	})
	do.Provide(injector, func(i *do.Injector) (merkController.MerkController, error) {
		return merkController.NewMerkController(i, ms), nil
	})
	do.Provide(injector, func(i *do.Injector) (itemController.ItemController, error) {
		return itemController.NewItemController(i, is), nil
	})
}
