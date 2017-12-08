package service

import (
	"encoding/json"
	"net/http"

	"github.com/HinanawiTenshi/agenda/service/entities"
	"github.com/unrolled/render"
)

func getUserKeyHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		user, err := entities.UserService.FindBy("username",
			req.FormValue("username"))
		panicIfErr(err)
		if user.Password == req.FormValue("password") {
			formatter.JSON(w, http.StatusOK, user)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}

func createNewUserHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var user entities.User
		err := decoder.Decode(&user)
		panicIfErr(err)
		check, err := entities.UserService.FindBy("username", user.Username)
		if check != (entities.User{}) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		panicIfErr(err)
		err = entities.UserService.Insert(&user)
		panicIfErr(err)
		formatter.JSON(w, http.StatusCreated, user)
	}

}
