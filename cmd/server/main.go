package main

import (
	"log"
	"github.com/NicolasPetruci/Figest-ComprasService/internal/config"
	"github.com/NicolasPetruci/Figest-ComprasService/internal/database"
	"github.com/NicolasPetruci/Figest-ComprasService/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.ConnectDB()

	r := gin.Default()
	routes.RegisterRoutes(r)

	log.Println("Starting server on port 3003...")
	r.Run(":3003")
}
