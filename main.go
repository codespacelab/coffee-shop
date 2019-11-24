package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Coffee model
type Coffee struct {
	ID    string  `json:"id"`
	Size  string  `json:"size"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//Init values/DB
var coffees []Coffee

//Get All
func getCoffees(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(coffees)
}

//Get Coffee
func getCoffee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range coffees {
		if item.Name == params["name"] {
			json.NewEncoder(response).Encode(item)
			return
		}
	}
}

//Add Coffee
func addCoffee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var newCoffee Coffee
	json.NewDecoder(request.Body).Decode(&newCoffee)
	newCoffee.ID = strconv.Itoa(len(coffees) + 1)
	coffees = append(coffees, newCoffee)
	json.NewEncoder(response).Encode(newCoffee)
}

//Update Coffee
func updateCoffee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for i, item := range coffees {
		if item.ID == params["id"] {
			coffees = append(coffees[:i], coffees[i+1:]...)
			var newCoffee Coffee
			json.NewDecoder(request.Body).Decode(&newCoffee)
			newCoffee.ID = params["id"]
			coffees = append(coffees, newCoffee)
			json.NewEncoder(response).Encode(newCoffee)
			return
		}
	}
}

//Delete Coffee
func deleteCoffee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	for i, item := range coffees {
		if item.ID == params["id"] {
			coffees = append(coffees[:i], coffees[i+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(coffees)
}

func main() {

	coffees = append(coffees, Coffee{ID: "1", Size: "Medium", Name: "Latte", Price: 2.99})
	coffees = append(coffees, Coffee{ID: "2", Size: "Small", Name: "Latte", Price: 1.99})
	coffees = append(coffees, Coffee{ID: "3", Size: "Large", Name: "Americano", Price: 3.20})

	handler := mux.NewRouter()

	//endpoints
	handler.HandleFunc("/coffee", getCoffees).Methods("GET")
	handler.HandleFunc("/coffee/{name}", getCoffee).Methods("GET")
	handler.HandleFunc("/coffee", addCoffee).Methods("POST")
	handler.HandleFunc("/coffee/{id}", updateCoffee).Methods("PUT")
	handler.HandleFunc("/coffee/{id}", deleteCoffee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", handler))

}
