package main

import (
    "github.com/spaaws/book-api/config"
    "github.com/spaaws/book-api/models"
    "github.com/spaaws/book-api/routes"
    "os"
)

func main() {
    config.Connect()
    config.DB.AutoMigrate(&models.Book{})

    // Pega a porta das variáveis de ambiente ou usa a porta 8080 como fallback
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r := routes.SetupRoutes()
    r.Run(":" + port) // Inicia o servidor na porta definida
}
