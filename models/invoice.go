package models

type Invoice struct {
    ID          int
    CustomerID  int
    InvoiceDate string
    Amount      float64
}
