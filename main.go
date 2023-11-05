

package main

import (
    "project/db"
    "project/models"
)

func main() {
    db.Init()

    
    customer1 := models.Customer{Name: "Wang Joe", Email: "wang@gmail.com", Address: "45 Kg St"}
    customer2 := models.Customer{Name: "Alice Chan", Email: "alice@yahoo.com", Address: "77 Elm St"}

    db.InsertCustomer(&customer1)
    db.InsertCustomer(&customer2)

   
    invoice1 := models.Invoice{CustomerID: 1, InvoiceDate: "2023-11-01", Amount: 100.00}
    invoice2 := models.Invoice{CustomerID: 1, InvoiceDate: "2023-11-02", Amount: 150.50}
    invoice3 := models.Invoice{CustomerID: 2, InvoiceDate: "2023-11-01", Amount: 75.25}

    db.InsertInvoice(&invoice1)
    db.InsertInvoice(&invoice2)
    db.InsertInvoice(&invoice3)
}
