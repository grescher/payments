package main

import (
	"fmt"
	"net/http"
	"payments/config"
	"reflect"
	"runtime"
)

func main() {
	server := http.Server{
		Addr: config.ServerAddress() + ":" + config.ServerPort(),
	}
	http.HandleFunc("/hello", log(hello))
	http.HandleFunc("/world", log(world))

	server.ListenAndServe()
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!\n")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!\n")
}
