package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize Router
	router := mux.NewRouter()

	// Register Routes
	router.Path("/").Methods(http.MethodGet).HandlerFunc(healthCheck)

	// Serve
	log.Panicln("Serving on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "healthy af")
}
