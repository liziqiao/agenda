package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func listAllMeetingsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		// TODO
	}

}

func createNewMeetingHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		// TODO
	}

}

func clearMeetingsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		// TODO
	}

}
