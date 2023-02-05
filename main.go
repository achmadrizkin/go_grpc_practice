package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/achmadrizkin/go_grpc_practice/database"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
}
