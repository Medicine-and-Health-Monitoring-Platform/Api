package api

import (
	"Api_Gateway/api/handler"
	"Api_Gateway/api/middleware"

	"Api_Gateway/config"

	_ "Api_Gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Personalized Medicine and Health Monitoring Platform
// @version 1.0
// @description API Gateway of Personalized Medicine and Health Monitoring Platform
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(cfg *config.Config) *gin.Engine {
	h := handler.NewHandler(cfg)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	api.Use(middleware.Check)

	admin := api.Group("/admin")
	{
		u := admin.Group("/user")
		{
			u.GET("/:id", h.GetUser)
			u.PUT("/:id", h.UpdateUser)
			u.DELETE("/:id", h.DeleteUser)
		}
		admin.GET("/users", h.FetchUsers)

	}

	user := api.Group("/customer")
	{

		p := user.Group("/profile")
		{
			p.GET("/", h.GetProfile)
			p.PUT("/", h.UpdateProfile)
		}

		medical := user.Group("/medical")
		{
			medical.POST("/", h.AddMedicalRecord)
			medical.GET("/:id", h.GetMedicalRecord)
			medical.PUT("/:id", h.UpdateMedicalRecord)
			medical.DELETE("/:id", h.DeleteMedicalRecord)
			medical.GET("/", h.ListMedicalRecords)
		}

		lifestyle := user.Group("/lifestyle")
		{
			lifestyle.POST("", h.AddLifestyleData)
			lifestyle.GET("/:id", h.GetLifestyleData)
			lifestyle.PUT("/:id", h.UpdateLifestyleData)
			lifestyle.DELETE("/:id", h.DeleteLifestyleData)
		}

		wearable := user.Group("/wearable")
		{
			wearable.POST("/", h.AddWearableData)
			wearable.GET("/:id", h.GetWearableData)
			wearable.PUT("/:id", h.UpdateWearableData)
			wearable.DELETE("/:id", h.DeleteWearableData)
		}

		monitoring := user.Group("/monitoring")
		{
			monitoring.POST("/recommendations", h.GenerateHealthRecommendations)
			monitoring.GET("/realtime", h.GetRealtimeHealthMonitoring)
			monitoring.GET("/daily", h.GetDailyHealthSummary)
			monitoring.GET("/weekly", h.GetWeeklyHealthSummary)
			// monitoring.POST("/")
		}
	}

	return router
}
