package main

import (
	"github.com/RealWorldGoSolution/sec5/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	if err := database.Exec(db); err != nil {
		panic(err)
	}
}
