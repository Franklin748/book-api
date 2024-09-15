package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/spaaws/book-api/controllers"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    r.GET("/books", controllers.GetBooks)
    r.GET("/books/:id", controllers.GetBook)
    r.POST("/books", controllers.CreateBook)
    r.PUT("/books/:id", controllers.UpdateBook)
    r.DELETE("/books/:id", controllers.DeleteBook)

    return r
}
