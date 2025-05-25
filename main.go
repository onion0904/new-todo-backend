package main

import (
	"TodoApp/controllers"
	"TodoApp/models"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CORS ミドルウェア
func CORSMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	dbUser := os.Getenv("USERNAME")
	dbPassword := os.Getenv("USERPASS")
	dbDatabase := os.Getenv("DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3308)/%s?parseTime=true", dbUser,dbPassword, dbDatabase)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	err = db.AutoMigrate(&models.Todo{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

	Todo := gin.Default()

	Todo.Use(CORSMiddleware())

	con := controllers.NewTodoController(db)

	method := Todo.Group("/Todo")
	{
		method.POST("/add",con.Add)
		method.GET("/list",con.List)
		method.GET("/list/sorted",con.SortedList)
		method.PUT("/update",con.Update)
		method.DELETE("/delete",con.Delete)
	}

    Todo.Run(":8080")
}