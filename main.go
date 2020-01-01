package main

import (
	"fmt"
	"tara/routes"
	"tara/dependencies"
)

func main() {

	if err := dependencies.InitConfig(); err != nil {
		fmt.Println("Initiate config failed.")
	}
	db, err := dependencies.InitDB()
	if err != nil {
		fmt.Println("Connect to database failed.")
		return
	}

	defer db.Close()

	router := routes.InitRoute()

	router.Run()
}
