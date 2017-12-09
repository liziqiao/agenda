package service

import (
	"net/http"

	"github.com/HinanawiTenshi/agenda/service/entities"
	"github.com/unrolled/render"
)

func listAllUsersHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		key := getKeyFromRequest(req)
		if verifyKey(key) {
			userList, err := entities.UserService.FindAll()
			for i := range userList {
				userList[i].Key = "******"
				userList[i].Password = "******"
			}
			panicIfErr(err)
			formatter.JSON(w, http.StatusOK, userList)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}

func deleteCurrentUserHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		key := getKeyFromRequest(req)
		if verifyKey(key) {
			err := entities.UserService.DeleteByKey(key)
			panicIfErr(err)
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}
