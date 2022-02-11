package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/stonelike/CleanGo/src/api/handler"
	"github.com/stonelike/CleanGo/src/api/presenter"
	"github.com/stonelike/CleanGo/src/usecase/user"
)

func createUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var input struct {
				Name string `json:"name"`
			}

			err := json.NewDecoder(r.Body).Decode(&input)

			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			user, err := service.CreateUser(input.Name)

			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				handler.HttpErrorResponse(w, err)
				return
			}

			toJsonUser := &presenter.User{
				Id:   user.GetId(),
				Name: user.GetName(),
			}

			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(toJsonUser); err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		})
}

func getUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			vars := mux.Vars(r) ///user/{id}みたいなルートバイディングを扱うのに必要

			user, err := service.FindById(vars["id"])
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				handler.HttpErrorResponse(w, err)
				return
			}

			toJsonUser := &presenter.User{
				Id:   user.GetId(),
				Name: user.GetName(),
			}

			if err := json.NewEncoder(w).Encode(toJsonUser); err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		})
}

func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {

	r.Handle("/user", n.With(
		negroni.Wrap(createUser(service)),
	)).Methods("POST", "OPTIONS").Name("createUser")

	r.Handle("/user/{id}", n.With(
		negroni.Wrap(getUser(service)),
	)).Methods("GET", "OPTIONS").Name("getUser")
}
