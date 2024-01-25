package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type CustomerOrder struct {
	CustomerID  string `json:"CustomerID"`
	CompanyName string `json:"CompanyName"`
	ContactName string `json:"ContactName"`
	OrderID     int    `json:"OrderID"`
	OrderDate   string `json:"OrderDate"`
	ShipName    string `json:"ShipName"`
	ShipAddress string `json:"ShipAddress"`
	ShipCity    string `json:"ShipCity"`
	ShipRegion  string `json:"ShipRegion"`
}

type Customer struct {
	CustomerID  string `json:"CustomerID"`
	CompanyName string `json:"CompanyName"`
	ContactName string `json:"ContactName"`
}

type Order struct {
	OrderID        int            `json:"OrderID"`
	CustomerID     string         `json:"CustomerID"`
	EmployeeID     int            `json:"EmployeeID"`
	OrderDate      sql.NullString `json:"OrderDate"`
	RequiredDate   sql.NullString `json:"RequiredDate"`
	ShippedDate    sql.NullString `json:"ShippedDate"`
	ShipVia        int            `json:"ShipVia"`
	Freight        float64        `json:"Freight"`
	ShipName       sql.NullString `json:"ShipName"`
	ShipAddress    sql.NullString `json:"ShipAddress"`
	ShipCity       sql.NullString `json:"ShipCity"`
	ShipRegion     sql.NullString `json:"ShipRegion"`
	ShipPostalCode sql.NullString `json:"ShipPostalCode"`
	ShipCountry    sql.NullString `json:"ShipCountry"`
}

type OrderDetails struct {
   OrderID int `json:"OrderId"`
   ProductID int `json:"ProductId"`
   UnitPrice float64 `json:"UnitPrice"`
   Quantity int `json:"Quantity"`
   Discount float64 `json:"Discount"` 
}

type Product struct {
   ProductID int `json:"ProductId"`
   ProductName string `json:"ProductName"`
   SupplierID int `json:"SupplierId"`
   CategoryID int `json:"CategoryId"`
   QuantityPerUnit string `json:"QuantityPerUnit"`
   UnitPrice float64 `json:"UnitPrice"`
   UnitsInStock int `json:"UnitsInStock"`
   UnitsOnOrder int `json:"UnitsOnOrder"`
   ReorderLevel int `json:"ReorderLevel"`
   Discontinued int `json:"Discontinued"`
}

type Category struct {
   CategoryID int `json:"CategoryId"`
   CategoryName string `json:"CategoryName"`
   Description string `json:"Description"`
   Picture string `json:"Picture"` // this is supposed to be like a blob or something
}

func main() {
	db, err := sql.Open("sqlite3", "northwind.db")
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
	router.HandleFunc("/", BaseHandler).Methods("GET")

	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		GetCustomerOrdersHandler(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		GetAllOrdersHandler(w, r, db)
	}).Methods("GET")

	fmt.Println("Starting server on the port 8080...")
	http.ListenAndServe(":8080", router)
}

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func GetAllOrdersHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	orders, err := fetchAllOrders(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Endpoint Hit: All Orders Endpoint")

	response, err := json.Marshal(orders)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func fetchAllOrders(db *sql.DB) ([]Order, error) {
	rows, err := db.Query(`SELECT * FROM Orders`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.OrderID,
			&order.CustomerID,
			&order.EmployeeID,
			&order.OrderDate,
			&order.RequiredDate,
			&order.ShippedDate,
			&order.ShipVia,
			&order.Freight,
			&order.ShipName,
			&order.ShipAddress,
			&order.ShipCity,
			&order.ShipRegion,
			&order.ShipPostalCode,
			&order.ShipCountry,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func GetCustomerOrdersHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	customers, err := fetchCustomerOrders(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Endpoint Hit: All Customers Endpoint")

	response, err := json.Marshal(customers)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func fetchCustomerOrders(db *sql.DB) ([]CustomerOrder, error) {
	rows, err := db.Query(`
      SELECT 
         Customers.CustomerID, 
         Customers.CompanyName,
         Customers.ContactName,
         Orders.OrderID, 
         Orders.OrderDate, 
         Orders.ShipCity, 
         Orders.ShipName, 
         Orders.ShipAddress, 
         Orders.ShipRegion
      FROM Customers
      INNER JOIN Orders 
      ON Customers.CustomerID = Orders.CustomerID
   `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []CustomerOrder

	for rows.Next() {
		var order CustomerOrder
		err := rows.Scan(
			&order.CustomerID,
			&order.CompanyName,
			&order.ContactName,
			&order.OrderID,
			&order.OrderDate,
			&order.ShipCity,
			&order.ShipName,
			&order.ShipAddress,
			&order.ShipRegion,
		)

		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func fetchCustomers(db *sql.DB) ([]Customer, error) {
	rows, err := db.Query(`SELECT * FROM Customers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer

	for rows.Next() {
		var customer Customer
		err := rows.Scan(
			&customer.CustomerID,
			&customer.CompanyName,
			&customer.ContactName,
		)

		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func fetchOrders(db *sql.DB) ([]Order, error) {
	rows, err := db.Query(`SELECT * FROM Orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.OrderID,
			&order.CustomerID,
			&order.EmployeeID,
			&order.OrderDate,
			&order.RequiredDate,
			&order.ShippedDate,
			&order.ShipVia,
			&order.ShipAddress,
		)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func fetchTerritories(db *sql.DB) ([]string, error) {
   rows, err := db.Query(`SELECT * FROM Territories`)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   var territories []string

   for rows.Next() {
      var territory string
      err := rows.Scan(&territory)
      if err != nil {
         return nil, err
      }

      territories = append(territories, territory)
   }

   return territories, nil
}

func fetchOrderDetails(db *sql.DB) ([]OrderDetails, error) {
   rows, err := db.Query(`SELECT * FROM OrderDetails`)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   var orderDetails []OrderDetails

   for rows.Next() {
      var orderDetail OrderDetails
      err := rows.Scan(
         &orderDetail.OrderID,
         &orderDetail.ProductID,
         &orderDetail.UnitPrice,
         &orderDetail.Quantity,
         &orderDetail.Discount,
      )
      if err != nil {
         return nil, err
      }

      orderDetails = append(orderDetails, orderDetail)
   }

   return orderDetails, nil
}
