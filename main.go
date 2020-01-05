package main

import (
	"fmt"
	"tara/dependencies"
	"tara/routes"
)

func main() {

	if err := dependencies.InitConfig(); err != nil {
		fmt.Printf("Initiate config failed: %v\n", err)
	}
	db, err := dependencies.InitDB()
	if err != nil {
		fmt.Printf("Connect to database failed: %v\n", err)
		return
	}

	defer db.Close()

	router := routes.InitRoute()

	router.Run()
}
