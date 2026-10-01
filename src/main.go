package main

import (
	Routes "delivery-go-api/src/infra"
	"delivery-go-api/src/infra/connections/db"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main()	{
  if err := godotenv.Load(); err != nil {
    log.Fatal("Erro ao carregar as variaveis de ambiente do arquivo .env")
  }

	database, err := db.ConnectToDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	if err := db.Migrate(database); err != nil {
		log.Fatal("Erro ao migrar o banco de dados:", err)
	}

	router := gin.Default()
	Routes.SetupRoutes(router, database)
	router.Run(":8080")
}