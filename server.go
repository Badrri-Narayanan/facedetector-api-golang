package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/model"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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

	router.GET("/signin", func(ctx *gin.Context) {
		body := ctx.Request.Body
		values, err := ioutil.ReadAll(body)
		if err != nil {
			panic("Unable to read body")
		}

		signIn := model.SignIn{}
		json.Unmarshal([]byte(values), &signIn)

		query := fmt.Sprintf("SELECT * from login where email = '%v';", signIn.Email)

		result, err := db.Query(query)
		if err != nil {
			panic(err)
		}
		var loginDbInfo model.Login

		for result.Next() {
			result.Scan(&loginDbInfo.Id, &loginDbInfo.Hash, &loginDbInfo.Email)
		}

		err = bcrypt.CompareHashAndPassword([]byte(loginDbInfo.Hash), []byte(signIn.Password))
		if err != nil {
			ctx.String(http.StatusOK, "Incorrect Password")
		} else {
			ctx.String(http.StatusOK, "Correct Password")
		}
	})
	router.Run(fmt.Sprintf(":%v", port))
}
