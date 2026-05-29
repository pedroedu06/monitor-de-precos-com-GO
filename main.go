package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: .env nao encontrado, usando variaveis de ambiente do sistema")
	}

	router := gin.Default()

	group := router.Group("/")
	service.InitRoutes(group)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}