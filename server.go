package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"server/controller"
	"server/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	router := gin.Default()
	router.Use(middleware.HandleCrossOriginRequest())

	connStr := os.Getenv("DATABASE_URL")

	if len(connStr) == 0 {
		connStr = "postgres://badrri:secret123@localhost:5432/test_db?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Error", err)
		panic("Error connecting to DB ")
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("The server is up and running at port %v", port))
	})

	router.POST("/signin", func(ctx *gin.Context) { controller.HandleSignIn(ctx, db) })
	router.POST("/register", func(ctx *gin.Context) { controller.HandleRegister(ctx, db) })
	router.GET("/restaurant_menu", controller.HandeListOfFoodItemInMenu)

	router.Run(fmt.Sprintf(":%v", port))
}
