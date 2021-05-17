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
		panic(fmt.Errorf("failed to load config: %s", err))
	}

	// Init mysql connection
	if err := db.D.Init(); err != nil {
		panic(fmt.Errorf("failed to connect mysql: %s", err))
	}

	router := router.Init()
	router.Run(fmt.Sprintf("%v:%v", config.Cfg.Web.Address, config.Cfg.Web.Port))
}
