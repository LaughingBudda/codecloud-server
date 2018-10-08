package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	. "github.com/LaughingBudda/codecloud-server/dao"
	. "github.com/LaughingBudda/codecloud-server/models"
	"codecloud/constants"
	"github.com/auth0-community/auth0"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	jose "gopkg.in/square/go-jose.v2"
)

var dao = DAO{}

const (
	mongoUrl = "localhost:27017"
	dbName   = "CodeCloud"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	dao.Server = mongoUrl
	dao.Database = dbName
	dao.Connect()
}

func main() {
	r := mux.NewRouter()

	r.Handle("/user", AllUsers).Methods("GET")
	r.Handle("/user/{id}", FindUser).Methods("GET")
	r.Handle("/createuser/", CreateUser).Methods("POST")
	r.Handle("/removeuser/", RemoveUser).Methods("DELETE")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte(constants.JWT_SECRET)
		secretProvider := auth0.NewKeyProvider(secret)
		audience := []string{constants.API_AUDIENCE}

		configuration := auth0.NewConfiguration(secretProvider, audience, constants.AUTH0_DOMAIN, jose.HS256)
		validator := auth0.NewValidator(configuration, nil)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

var FindUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	respondWithJson(w, http.StatusOK, user)
})

var AllUsers = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, user)
})

var CreateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.Uid = bson.NewObjectId()
	if err := dao.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
})

var RemoveUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
})
