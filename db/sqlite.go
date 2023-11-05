package db
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var db *sql.DB

func Init() {
    var err error
    db, err = sql.Open("sqlite3", "store.db")
    if err != nil {
        log.Fatal(err)
    }

    createTables()
}

func createTables() {
    createCustomersTable()
    createInvoicesTable()
}

func createCustomersTable() {
    query := `
        CREATE TABLE IF NOT EXISTS customers (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            email TEXT,
            address TEXT
        )
    `

    _, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
}

func createInvoicesTable() {
    query := `
        CREATE TABLE IF NOT EXISTS invoices (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            customer_id INTEGER,
            invoice_date TEXT,
            amount REAL
        )
    `

    _, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
}
