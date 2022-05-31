package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/milinches/contacts-app-backend/api/models"
	"github.com/milinches/contacts-app-backend/api/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Create new user"))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a single user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
