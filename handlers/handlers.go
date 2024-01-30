package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

   "go-northwind-server/models"
)

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

func fetchAllOrders(db *sql.DB) ([]models.Order, error) {
	rows, err := db.Query(`SELECT * FROM Orders`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
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

func fetchCustomerOrders(db *sql.DB) ([]models.CustomerOrder, error) {
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

	var orders []models.CustomerOrder

	for rows.Next() {
		var order models.CustomerOrder
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

func GetCustomersHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
   customers, err := fetchCustomers(db)
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

func fetchCustomers(db *sql.DB) ([]models.Customer, error) {
	rows, err := db.Query(`SELECT * FROM Customers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer

	for rows.Next() {
		var customer models.Customer
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

func fetchOrders(db *sql.DB) ([]models.Order, error) {
	rows, err := db.Query(`SELECT * FROM Orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order

	for rows.Next() {
		var order models.Order
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

func GetTerritoriesHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
   territories, err := fetchTerritories(db)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println("Endpoint Hit: All Customers Endpoint")

   response, err := json.Marshal(territories)
   if err != nil {
      fmt.Println(err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }

   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(http.StatusOK)
   w.Write(response)
}

func fetchTerritories(db *sql.DB) ([]models.Territory, error) {
   rows, err := db.Query(`SELECT * FROM territories`)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   var territories []models.Territory

   for rows.Next() {
      var territory models.Territory
      err := rows.Scan(
         &territory.TerritoryID, 
         &territory.TerritoryDescription, 
         &territory.RegionID,
      )
      if err != nil {
         return nil, err
      }

      territories = append(territories, territory)
   }

   return territories, nil
}


func GetOrderDetailsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
   orderDetails, err := fetchOrderDetails(db)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println("Endpoint Hit: All Order Details Endpoint")

   response, err := json.Marshal(orderDetails)
   if err != nil {
      fmt.Println(err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }

   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(http.StatusOK)
   w.Write(response)
}

func fetchOrderDetails(db *sql.DB) ([]models.OrderDetails, error) {
   rows, err := db.Query(`SELECT * FROM OrderDetails`)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   var orderDetails []models.OrderDetails

   for rows.Next() {
      var orderDetail models.OrderDetails
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


func GetCustomerOrderCountHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
   customerOrderTotals, err := customerOrderCount(db)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println("Endpoint Hit: All Customer Order Counts Endpoint")

   response, err := json.Marshal(customerOrderTotals)
   if err != nil {
      fmt.Println(err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }

   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(http.StatusOK)
   w.Write(response)
}

func customerOrderCount(db *sql.DB) ([]models.CustomerOrderCount, error) {
   rows, err := db.Query(`
      SELECT 
         Customers.CustomerID AS ID,
         Customers.ContactName AS CustomerName,
         Customers.City AS CustomerCity,
         COUNT(Orders.OrderID) AS OrderCount
      FROM Orders 
      JOIN Customers ON Customers.CustomerID = Orders.CustomerID
      GROUP BY ID
      ORDER BY OrderCount DESC;
   `)

   if err != nil {
      return nil, err
   }
   defer rows.Close()

   var customerOrders []models.CustomerOrderCount

   for rows.Next() {
      var customerOrder models.CustomerOrderCount
      err := rows.Scan(
         &customerOrder.ID,
         &customerOrder.CustomerName,
         &customerOrder.CustomerCity,
         &customerOrder.OrderCount,
      )
      if err != nil {
         return nil, err
      }

      customerOrders = append(customerOrders, customerOrder)
   }

   return customerOrders, nil
}
