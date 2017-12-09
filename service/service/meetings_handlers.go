package service

import (
	"encoding/json"
	"net/http"

	"github.com/HinanawiTenshi/agenda/service/entities"
	"github.com/unrolled/render"
)

func listAllMeetingsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		key := getKeyFromRequest(req)
		if verifyKey(key) {
			meetingList, err := entities.MeetingService.FindAll()
			panicIfErr(err)
			formatter.JSON(w, http.StatusOK, meetingList)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}

func createNewMeetingHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		key := getKeyFromRequest(req)
		if verifyKey(key) {
			decoder := json.NewDecoder(req.Body)
			var meeting entities.Meeting
			err := decoder.Decode(&meeting)
			panicIfErr(err)
			host, err := entities.UserService.FindBy("key", key)
			panicIfErr(err)
			meeting.Host = host.Username
			err = entities.MeetingService.Insert(&meeting)
			panicIfErr(err)
			formatter.JSON(w, http.StatusCreated, meeting)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}

func clearMeetingsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		key := getKeyFromRequest(req)
		if verifyKey(key) {
			user, err := entities.UserService.FindBy("key", key)
			panicIfErr(err)
			err = entities.MeetingService.DeleteMeetingsHostedByUser(user.Username)
			panicIfErr(err)
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}
