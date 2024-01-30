package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

   "go-northwind-server/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	db, err := sql.Open("sqlite3", "/usr/src/app/northwind.db")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error in opening database")
	}
	defer db.Close()

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error in pinging database")
		return
	}

	fmt.Println("Successfully connected to database!")

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.BaseHandler).Methods("GET")

	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCustomersHandler(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllOrdersHandler(w, r, db)
	}).Methods("GET")

   router.HandleFunc("/customer_orders", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCustomerOrdersHandler(w, r, db)
	}).Methods("GET")

   router.HandleFunc("/territories", func(w http.ResponseWriter, r *http.Request) {
      handlers.GetTerritoriesHandler(w, r, db)
   }).Methods("GET")

   router.HandleFunc("/customer_order_count", func(w http.ResponseWriter, r *http.Request) {
      handlers.GetCustomerOrderCountHandler(w, r, db)
   }).Methods("GET")

	fmt.Println("Starting server on the port 8080...")
	http.ListenAndServe(":8080", router)
}
