package service

import (
	"net/http"

	"github.com/HinanawiTenshi/agenda/service/entities"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// NewServer returns a negroni server that has already initialized
// routes
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{IndentJSON: true})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	// # User
	// ## User-key
	// ### Get user key
	mx.HandleFunc("/v1/key", getUserKeyHandler(formatter)).Methods("GET")

	// ### Create a new user
	mx.HandleFunc("/v1/key", createNewUserHandler(formatter)).Methods("POST")

	// # Users
	// ## Users Collection
	// ### List all Users
	mx.HandleFunc("/v1/users", listAllUsersHandler(formatter)).Methods("GET")

	// ### Delete current user
	mx.HandleFunc("/v1/users", deleteCurrentUserHandler(formatter)).Methods("DELETE")

	// # Meetings
	// ## Meetings Collection
	// ### List all meetings
	mx.HandleFunc("/v1/meetings", listAllMeetingsHandler(formatter)).Methods("GET")

	// ### Create a new meeting
	mx.HandleFunc("/v1/meetings", createNewMeetingHandler(formatter)).Methods("POST")

	// ### Clear all meetings
	mx.HandleFunc("/v1/meetings", clearMeetingsHandler(formatter)).Methods("DELETE")
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getKeyFromRequest(req *http.Request) string {
	req.ParseForm()
	return req.FormValue("key")
}

func verifyKey(key string) bool {
	user, err := entities.UserService.FindBy("key", key)
	panicIfErr(err)
	return user != (entities.User{})
}
