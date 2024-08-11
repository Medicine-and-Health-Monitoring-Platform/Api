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
			u.GET("/:id", h.GetProfile)
			u.PUT("/:id", h.UpdateProfile)
			u.DELETE("/:id", h.DeleteUser)
		}
		admin.GET("/users", h.FetchUsers)

		//p := admin.Group("/product")
		//{
		//	p.POST("", h.PostProduct)
		//	p.GET("/:id", h.GetProduct)
		//	p.PUT("/:id", h.PutProduct)
		//	p.DELETE("/:id", h.DeleteProduct)
		//	p.GET("/all", h.FetchProducts)
		//}

	}

	user := api.Group("/customer")
	{
		//a := user.Group("/order")
		//{
		//	a.POST("/", h.MakeOrder)
		//	a.GET("/:id", h.GetOrder)
		//}
		//c := user.Group("/category")
		//{
		//	c.GET("/all", h.GetCategories)
		//	c.POST("/", h.CreateCategory)
		//	c.PUT("/:id", h.UpdateCategory)
		//	c.DELETE("/:id", h.DeleteCategory)
		//}
		p := user.Group("/profile")
		{
			p.GET("/", h.GetProfile)
			p.PUT("/", h.UpdateProfile)
		}

		//prod := user.Group("/product")
		//prod.GET("/all", h.FetchProducts)
		//
		//b := user.Group("/basket")
		//{
		//	b.POST("", h.AddToBasket)
		//	b.GET("/:id", h.GetBasket)
		//	b.PUT("/:id", h.UpdateBasket)
		//	b.DELETE("/:id", h.RemoveFromBasket)
		//}
		//r := user.Group("/review")
		//{
		//	r.POST("/", h.CreateReview)
		//	r.PUT("/:id", h.UpdateReview)
		//	r.DELETE("/:id", h.DeleteReview)
		//	r.GET("/all", h.FetchReviews)
		//}
	}

	return router
}
