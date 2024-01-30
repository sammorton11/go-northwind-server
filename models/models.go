package models

import (
   "database/sql"
)


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

type CustomerOrderCount struct {
   ID int `json:"ID"`
   CustomerName string `json:"CustomerName"`
   CustomerCity string `json:"CustomerCity"`
   OrderCount int `json:"OrderCount"`
}

type Territory struct {
   TerritoryID string `json:"TerritoryId"`
   TerritoryDescription string `json:"TerritoryDescription"`
   RegionID int `json:"RegionId"`
}
