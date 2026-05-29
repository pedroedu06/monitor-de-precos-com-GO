package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/health", controller.Ping)
	r.POST("/auth", controller.LoginUser)
	r.POST("/usuarios", controller.CreateUser)

	api := r.Group("/api")
	{
		api.POST("/produtos", controller.CreateProduct)
		api.GET("/produtos", controller.ListProduct)
    	api.DELETE("/produtos/:id", controller.DeleteProduct)
    	api.PATCH("/produtos/:id/pausar", controller.PauseProduct)
    	api.GET("/produtos/:id/historico", controller.HistoricPrices)
    	api.GET("/notificacoes", controller.ListNotifications)
	}
}