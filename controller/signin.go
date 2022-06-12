package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleSignIn(ctx *gin.Context, db *sql.DB) {
	body := ctx.Request.Body
	values, err := ioutil.ReadAll(body)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid request. Please provide username and passowrd")
	}

	signIn := model.SignIn{}
	json.Unmarshal([]byte(values), &signIn)

	query := fmt.Sprintf("SELECT * from login where email = '%v';", signIn.Email)

	result, err := db.Query(query)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Unable to connect to DB")
	}

	var loginDbInfo model.Login

	for result.Next() {
		result.Scan(&loginDbInfo.Id, &loginDbInfo.Hash, &loginDbInfo.Email)
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginDbInfo.Hash), []byte(signIn.Password))
	if err != nil {
		ctx.String(http.StatusForbidden, "Invalid Credentials")
	} else {
		ctx.String(http.StatusOK, "OK")
	}
}
