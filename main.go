package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Alumni struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	KnownFor string `json:"knownFor"`
}

var db map[int]Alumni
var id int

func main() {
	db = make(map[int]Alumni)
	id = 1

	// Initialize Router
	router := mux.NewRouter()

	// Register Routes
	router.Path("/").Methods(http.MethodGet).HandlerFunc(healthCheck)
	router.Path("/alumni").Methods(http.MethodGet).HandlerFunc(getAlumni)
	router.Path("/alumni").Methods(http.MethodPost).HandlerFunc(postAlumni)
	router.Path("/alumni/{id}").Methods(http.MethodDelete).HandlerFunc(deleteAlumni)

	// Serve
	log.Println("Serving on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "healthy af")
}

func getAlumni(w http.ResponseWriter, r *http.Request) {
	alumnis := []Alumni{}
	for _, alumni := range db {
		alumnis = append(alumnis, alumni)
	}

	jsonAlumnis, err := json.Marshal(alumnis)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonAlumnis))
}

func postAlumni(w http.ResponseWriter, r *http.Request) {
	var alumni Alumni
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&alumni)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	// Set ID
	alumni.ID = id
	// Add to database
	db[id] = alumni
	// Increment ID for next entry
	id++

	response, err := json.Marshal(alumni)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))

}

func deleteAlumni(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		fmt.Fprintln(w, "request not formatted properly")
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}

	_, found := db[intID]
	if !found {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "id does not exist")
		return
	}

	delete(db, intID)

	fmt.Fprintf(w, "deleted alumni with id: %s", id)
}
