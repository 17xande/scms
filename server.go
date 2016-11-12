package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)

	n := negroni.Classic() // includes default middlewares
	n.UseHandler(router)

	n.Run(":8080")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the home page again")
}
