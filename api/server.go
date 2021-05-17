package api

import (
	"fmt"

	"github.com/iamdavidzeng/tara/api/router"
	"github.com/iamdavidzeng/tara/internal/db"
	"github.com/iamdavidzeng/tara/pkg/config"
)

// Run start RESTful API service
func Run() {
	// Init config
	if err := config.Cfg.Init(); err != nil {
		fmt.Printf("Initiate config failed: %v\n", err)
	}

	// Init mysql connection
	if err := db.D.Init(); err != nil {
		fmt.Printf("Fail to onnect database: %v\n", err)
	}

	router := router.Init()

	router.Run()
}
