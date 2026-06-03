package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/repository"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/scrapper"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/service"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/worker"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env not found, using system environment variables")
	}

	router := gin.Default()

	group := router.Group("/")
	service.InitRoutes(group)

	productRepo := repository.NewProdutoRepository()
	priceHistoryRepo := repository.NewPriceHistoryRepository()
	priceScrapper := scrapper.NewScrapper()

	w := worker.NewWorker(productRepo, priceHistoryRepo, priceScrapper)

	go w.StartWorker()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}