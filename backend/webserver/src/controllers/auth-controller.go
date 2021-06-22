package controllers

import (
	"fmt"
	"go-sandbox/src/models"
	"go-sandbox/src/utils"
	"log"
	"net/http"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	var loginInfo = &models.Login{}
	log.Println("Parsing login data")
	utils.ParseBody(r, loginInfo)

	if loginInfo.User == "admin" && loginInfo.Password == "admin" {
		fmt.Fprint(w, "Success")
		return
	}

	fmt.Fprint(w, "Login error")
}
