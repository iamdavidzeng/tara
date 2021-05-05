package server

import (
	"context"
	"fmt"
	"tara/api/config"
	"tara/api/db"
	"tara/api/router"
)

// Run start RESTful API service
func Run() {
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

	router := router.Init()

	router.Run()
}
