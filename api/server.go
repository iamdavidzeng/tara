package api

import (
	"fmt"

	"github.com/iamdavidzeng/tara/api/router"
)

// Run start RESTful API service
func Run() {
	// Init config
	if err := Cfg.Init(); err != nil {
		fmt.Printf("Initiate config failed: %v\n", err)
	}

	// Init mysql connection
	_, err := InitMysql()
	if err != nil {
		fmt.Printf("Fail to onnect database: %v\n", err)
		return
	}

	router := router.Init()

	router.Run()
}
