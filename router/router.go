package router

import (
	"bytes"
	"fmt"
	"go-postgres/middleware"
	"io"
	"net/http"

	logkardo "github.com/DrifeTechnologies/drife-logging-service"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user", middleware.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newuser", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteuser/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")
	router.Use(Loggingmiddleware)
	return router
}

func Loggingmiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqbody, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(reqbody))
		fmt.Fprintln(w, "A message was received")

		logkardo.Requestbodymaker(r, reqbody)
		logkardo.Responsebodymaker(w)

		// do stuff before the handlers
		h.ServeHTTP(w, r)
	})
}
