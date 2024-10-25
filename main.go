package main

import (
	"log"
	"time"

	"github.com/ballteerajet/go-gorm/controllers"
	"github.com/ballteerajet/go-gorm/models"
	"github.com/ballteerajet/go-gorm/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := util.LoadEnv() // Call the LoadEnv function

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema (create tables, etc.)
	db.AutoMigrate(&models.User{}) // Add other models as needed

	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // Allow your Vue app's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Use the db instance in your routes
	r.POST("/login", controllers.Login(db))
	r.GET("/users/:id", controllers.GetUserByID(db))
	r.POST("/users", controllers.CreateUser(db))
	r.PUT("/users/:id", controllers.UpdateUser(db))
	r.DELETE("/users/:id", controllers.DeleteUser(db))

	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
