package main

import (
	"example/v3/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	handlers.HandlerRequest()
}
