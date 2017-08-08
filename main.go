// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {

// 	defer database.Close()

// 	//init the server
// 	router := mux.NewRouter()

// 	people = append(people, Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
// 	people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})
// 	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
// 	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
// 	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
// 	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

// 	/Initialize routes
// 	router := InitRouter()

// 	//starting the server
// 	log.Println("The RestAPI Api has been started on port", api.ConfigKeys.Port)
// 	log.Fatal(http.ListenAndServe(":"+api.ConfigKeys.Port, router))
// }

// package main

// import (
// 	"fmt"
// 	"html"
// 	"log"
// 	"net/http"
// )

// func main() {

// 	// Basic web server that responds to any requests by simply outputting the request url
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	})

// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
