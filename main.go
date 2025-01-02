package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	fmt.Println("Simple hello-world program using go-chi")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// the route and the handler func are passed to the router
	router.Get("/", HelloWorld)
	http.ListenAndServe(":3000", router)

}
