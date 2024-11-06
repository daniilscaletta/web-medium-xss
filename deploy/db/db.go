package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDBConnection() (*sql.DB, error) {
	dsn := "root:qwerty@tcp(127.0.0.1:3306)/vkakids"
	driver := "mysql"

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("error to connect db: %v", err)
	}
	fmt.Println("Success connection to DB")
	return db, nil
}
