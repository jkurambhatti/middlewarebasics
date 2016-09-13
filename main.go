package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

// func isLogged(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
// 	fmt.Fprint(rw, "checking logged in or not\n")
// 	if _, err := req.Cookie("session"); err != nil {
// 		fmt.Print("no cookie\n")
// 	}
// 	next(rw, req)
// 	fmt.Fprint(rw, "checked\n")
// }

// func valid(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
// 	fmt.Fprint(rw, "user is a valid\n")
// 	next(rw, req)
// 	fmt.Fprint(rw, "leaving validation\n")
// }

func isLogged(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "checking logged in or not\n")
	if _, err := req.Cookie("session"); err != nil {
		fmt.Print("no cookie\n")
	}
	// next(rw, req)
	fmt.Fprint(rw, "checked\n")
}

func valid(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "user is a valid\n")
	// next(rw, req)
	fmt.Fprint(rw, "leaving validation\n")
}

func fstatus(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "all is well\n")
}

func main() {
	r := mux.NewRouter()

	// n := negroni.New()
	// n.Use(negroni.HandlerFunc(isLogged))
	r.PathPrefix("/auth").Handler(negroni.New(
		// negroni.HandlerFunc(isLogged),
		// negroni.HandlerFunc(valid),
		negroni.Wrap(http.HandlerFunc(isLogged)),
		negroni.Wrap(http.HandlerFunc(valid)),
		negroni.Wrap(http.HandlerFunc(fstatus)),
	))

	r.HandleFunc("/", index)
	r.HandleFunc("/auth", auth)
	http.ListenAndServe(":3000", r)
}

func index(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "hello world of middleware")
}

func auth(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "restricted authorised access only")
}
