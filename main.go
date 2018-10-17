package main

import (
	"net/http"
	"os"

	"github.com/LaughingBudda/codecloud-server/apihandlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/user", apihandlers.AllUsers).Methods("GET")
	r.Handle("/user/{id}", apihandlers.FindUser).Methods("GET")
	r.Handle("/createuser/", apihandlers.CreateUser).Methods("POST")
	r.Handle("/removeuser/", apihandlers.RemoveUser).Methods("DELETE")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}
