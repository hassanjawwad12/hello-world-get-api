package main

import (
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", HelloWorld)

	http.ListenAndServe(":3001", router)
}
