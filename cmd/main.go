package main

import (
	"context"
	"fmt"
	"tara/api/config"
	"tara/api/db"
	"tara/api/routes"
)

func main() {

	// Init config
	if err := config.InitConfig(); err != nil {
		fmt.Printf("Initiate config failed: %v\n", err)
	}

	// Init mysql connection
	mySQLClient, err := db.InitDB()
	if err != nil {
		fmt.Printf("Fail to onnect database: %v\n", err)
		return
	}
	defer mySQLClient.Close()

	// Init mongodb connection
	mongoClient, err := db.InitMongo()
	if err != nil {
		fmt.Printf("Fail to connect mongodb: %v\n", err)
	}
	defer mongoClient.Disconnect(context.TODO())

	router := routes.InitRoute()

	router.Run()
}
