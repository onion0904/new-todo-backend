package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
	"os"
	"log"
	"TodoApp/models"
	"TodoApp/controllers"
	"github.com/joho/godotenv"
)

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