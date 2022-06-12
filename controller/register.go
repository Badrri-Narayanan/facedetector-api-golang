package controller

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/model"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegister(ctx *gin.Context, db *sql.DB) {
	body := ctx.Request.Body
	values, err := ioutil.ReadAll(body)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid request. Please provide username and passowrd")
	}

	registerNewUser := model.Register{}
	json.Unmarshal([]byte(values), &registerNewUser)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerNewUser.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	trx, err := db.Begin()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Unable to connect to DB")
		panic(err)
	}

	_, err = trx.Exec("INSERT INTO login (email, hash) VALUES($1,$2);", registerNewUser.Email, string(hashedPassword))

	if err != nil {
		trx.Rollback()
		ctx.String(http.StatusInternalServerError, "Unable to create user")
		panic(err)
	}

	_, err = trx.Exec("INSERT INTO users(name, email, entries, joined, age, pet) VALUES($1, $2, 0, $3, 0, '');",
		registerNewUser.Name, registerNewUser.Email, time.Now())

	if err != nil {
		trx.Rollback()
		ctx.String(http.StatusInternalServerError, "Unable to create user")
		panic(err)
	}
	err = trx.Commit()

	if err != nil {
		trx.Rollback()
		ctx.String(http.StatusInternalServerError, "Unable to create user")
		panic(err)
	} else {
		ctx.String(http.StatusCreated, "Created User successfully")
	}
}
