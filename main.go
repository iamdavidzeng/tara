package main

import (
	"context"
	"fmt"
	"tara/dependencies"
	"tara/routes"
)

func main() {

	// Init config
	if err := dependencies.InitConfig(); err != nil {
		fmt.Printf("Initiate config failed: %v\n", err)
	}

	// Init mysql connection
	db, err := dependencies.InitDB()
	if err != nil {
		fmt.Printf("Fail to onnect database: %v\n", err)
		return
	}
	defer db.Close()

	// Init mongodb connection
	client, err := dependencies.InitMongo()
	if err != nil {
		fmt.Printf("Fail to connect mongodb: %v\n", err)
	}
	defer client.Disconnect(context.TODO())

	router := routes.InitRoute()

	router.Run()
}
