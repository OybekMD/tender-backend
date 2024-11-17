package api

import (
	_ "tender/api/docs" // swag
	v1 "tender/api/handlers/v1"
	"tender/api/middleware"

	"tender/api/tokens"
	"tender/config"
	"tender/storage"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf     *config.Config
	Storage  storage.StorageI
	Enforcer *casbin.Enforcer
}

// @Title       	Asrlan-Monolithic
// @securityDefinitions.apikey BearerAuth
// @In          	header
// @Name        	Authorization
func New(option *Option) *gin.Engine {
	gin.SetMode(option.Conf.GinMode)
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	jwtHandler := tokens.JWTHandler{
		SigningKey: option.Conf.SigningKey,
	}

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:        option.Conf,
		Storage:    option.Storage,
		JWTHandler: jwtHandler,
		Enforcer:   option.Enforcer,
	})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.NewAuthorizer(option.Enforcer, jwtHandler, *option.Conf))

	api := router.Group("")

	// Auth
	api.POST("/login", handlerV1.Login)       // Login user
	api.POST("/register", handlerV1.Register) // Register a new user

	// Tender
	api.POST("/api/client/tenders", handlerV1.CreateTender)          // Create a new tender
	api.GET("/api/client/tenders", handlerV1.ListTenders)            // List all tenders
	api.PUT("/api/client/tenders/:id", handlerV1.UpdateTenderStatus) // Update a tender's status by ID
	api.DELETE("/api/client/tenders/:id", handlerV1.DeleteTender)    // Delete a tender by ID

	// Bid
	router.POST("/api/contractor/tenders/:tender_id/bid", handlerV1.SubmitBid)        // Submit a bid by contractor
	router.GET("/api/contractor/tenders/:tender_id/bid", handlerV1.ViewSubmittedBids) // View all submitted bids by contractor
	router.POST("/api/client/tender/:tender_id/award/:bid_id", handlerV1.AwardBid)    // Awarding a bid by client
	router.POST("/api/contractor/bids/:id", handlerV1.DeleteBid)                      // Delete a bid by contractor

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
