package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", homeHandler)

	n := negroni.New()
	recovery := negroni.NewRecovery()
	recovery.ErrorHandlerFunc = handleRecovery
	n.Use(recovery)
	n.Use(negroni.NewStatic(http.Dir("/public")))
	n.UseHandler(mux)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the home page again")
}

func handleRecovery(error interface{}) {
	// log error and try not to crash everything...
}
