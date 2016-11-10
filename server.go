package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "welcome to the home page!")
	})

	n := negroni.Classic() // includes default middlewares
	n.UseHandler(mux)

	http.ListenAndServe(":8080", n)
}
