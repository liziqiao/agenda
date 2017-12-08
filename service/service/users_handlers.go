package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func listAllUsersHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		// TODO
	}

}

func deleteCurrentUserHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		// TODO
	}

}
