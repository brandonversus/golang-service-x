package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

var (
	name string
	port string
	url  string
)

func main() {

	// Read environment variables
	name = os.Getenv("NAME")
	port = os.Getenv("PORT")
	url = os.Getenv("URL")

	// Config router
	router := chi.NewRouter()
	router.Get("/", handler)
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {})

	// Start up server
	fmt.Printf("Starting up server on port %s ...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func handler(w http.ResponseWriter, r *http.Request) {

	// If no route is defined, return list with service name
	if len(url) == 0 {
		fmt.Println("Creating trace ...")
		if err := json.NewEncoder(w).Encode(&trace{Hops: []string{name}}); err != nil {
			http.Error(w, "Could not encode hop", 500)
		} else {
			fmt.Println("Created trace")
		}
		return
	}

	// Call route and return error if fails
	fmt.Printf("Calling service %s ...\n", url)
	resp, err := http.DefaultClient.Get(url)
	if err != nil || resp.StatusCode != 200 {
		var message string
		if err != nil {
			message = fmt.Sprintf("Could not call route: %s: %v", url, err)
		} else {
			message = fmt.Sprintf("Could not call route: %s: %d", url, resp.StatusCode)
		}
		fmt.Println(message)
		http.Error(w, message, 500)
		return
	}

	// Add service name to list
	fmt.Println("Appending service name to trace ...")
	var t trace
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&t)
	t.Hops = append(t.Hops, name)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		http.Error(w, "Could not encode hop after appending service name", 500)
	} else {
		fmt.Println("Appended service name to trace")
	}
}

type trace struct {
	Hops []string `json:"hops"`
}
