package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/model"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	port := 8000
	router := gin.Default()
	router.SetTrustedProxies([]string{"http://localhost:3000/"})

	connStr := "postgres://badrri:secret123@localhost:5432/test_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Error", err)
		panic("Error connecting to DB ")
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("The server is up and running at port %v", port))
	})

	router.GET("/list_of_users", func(ctx *gin.Context) {
		result, err := db.Query("SELECT * FROM users;")
		if err != nil {
			fmt.Println("Error querying users", err)
			panic("Error querying users table")
		}
		users := []model.User{}

		for result.Next() {
			user := model.User{}
			result.Scan(&user.Id, &user.Name, &user.Email, &user.Entries, &user.Joined, &user.Age, &user.Pet)
			users = append(users, user)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	})
	router.Run(fmt.Sprintf(":%v", port))
}
