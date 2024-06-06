package main

import (
	"database/sql"
	"fmt"
)

func initDB() (*sql.DB, error) {
	dsn := "root:admin123@/go_tp"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("DB connection successful")
	return db, nil
}
