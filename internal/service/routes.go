package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/controller"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/model/middleware"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/health", controller.Ping)
	r.POST("/auth", controller.AuthUser)

	api := r.Group("/api")
	api.Use(middleware.Auth())
	{
		api.POST("/produtos", controller.CreateProduct)
		api.GET("/produtos", controller.ListProduct)
    	api.DELETE("/produtos/:id", controller.DeleteProduct)
    	api.PATCH("/produtos/:id/pause", controller.PauseProduct)
    	api.PATCH("/produtos/:id/resume", controller.ResumeProduct)
    	api.GET("/produtos/:id/historico", controller.HistoricPrices)
    	api.GET("/notificacoes", controller.ListNotifications)
	}
}