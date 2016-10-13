package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/vitorleandroloureiro/ACL/models"
)

func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//resource := models.Resource{"/acl", true};
	//resources := []models.Resource{resource};
	//role := models.Role{"admin", resources, resources};
	//roles := []models.Role{role};
	user := models.User{
		r.FormValue("name"),
		r.FormValue("username"),
		r.FormValue("email"),
		nil}

	user.Save();
}