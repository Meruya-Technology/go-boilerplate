// package main

// import (
// 	"github.com/evermos/evm-product/configs"
// 	evmbe "github.com/evermos/evm-product/external/evermos-be"
// 	"github.com/evermos/evm-product/infras"
// 	catalogrepository "github.com/evermos/evm-product/internal/domain/catalog/repository"
// 	catalogservice "github.com/evermos/evm-product/internal/domain/catalog/service"
// 	productrepo "github.com/evermos/evm-product/internal/domain/product/repository"
// 	productsrv "github.com/evermos/evm-product/internal/domain/product/service"
// 	productmodel "github.com/evermos/evm-product/internal/domain/product_model"
// 	"github.com/evermos/evm-product/internal/handlers"
// 	"github.com/evermos/evm-product/transport/http"
// 	"github.com/evermos/evm-product/transport/http/middleware"
// 	"github.com/evermos/evm-product/transport/http/router"
// 	"github.com/google/wire"
// )

// // Wiring for configurations.
// var configurations = wire.NewSet(configs.Get)

// // Wiring for persistences.
// var persistences = wire.NewSet(infras.ProvideMySQLConn,
// 	infras.ProvideEVMDBConn,
// 	infras.ESNewClient,
// 	infras.RedisNewClient)

// // Wiring external dependency
// var productModelExt = wire.NewSet(evmbe.ProvideService, wire.Bind(
// 	new(evmbe.Service),
// 	new(*evmbe.ServiceImpl),
// ))

// // Wiring model repository
// var modelRepository = wire.NewSet(productmodel.ProvideRepository, wire.Bind(
// 	new(productmodel.Repository),
// 	new(*productmodel.RepositoryImpl),
// ))

// // Wiring catalog repository
// var catalogRepository = wire.NewSet(catalogrepository.ProvideRepository, wire.Bind(
// 	new(catalogrepository.Repository),
// 	new(*catalogrepository.RepositoryImpl),
// ))

// // Wiring product repository
// var productRepository = wire.NewSet(productrepo.ProvideRepository, wire.Bind(
// 	new(productrepo.Repository),
// 	new(*productrepo.RepositoryImpl),
// ))

// // Wiring for domain services.
// var domainServices = wire.NewSet(
// 	productmodel.ProvideService,
// 	catalogservice.ProvideService,
// 	productsrv.ProvideService,
// )

// // Wiring for auth middleware.
// var authMiddleware = wire.NewSet(middleware.ProvideAuthentication)

// // Wiring for handlers.
// var handl = wire.NewSet(handlers.Provide)

// // Wiring for HTTP routing.
// var routing = wire.NewSet(router.ProvideRouter)

// func InitializeService() *http.HTTP {
// 	wire.Build(
// 		configurations,
// 		persistences,
// 		productModelExt,
// 		modelRepository,
// 		catalogRepository,
// 		productRepository,
// 		authMiddleware,
// 		domainServices,
// 		handl,
// 		routing,
// 		http.ProvideHTTP)
// 	return &http.HTTP{}
// }