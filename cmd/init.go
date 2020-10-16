package main

import (
	"github.com/joho/godotenv"

	"github.com/phuwn/lightning/src/server"
	"github.com/phuwn/lightning/src/service"
	"github.com/phuwn/lightning/src/store"
	"github.com/phuwn/tools/log"
	"github.com/phuwn/tools/util"
)

// init server stuff
func init() {
	env := util.Getenv("RUN_MODE", "")
	if env == "local" || env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	store := store.New()
	service := service.New()

	server.NewServerCfg(store, service)
}
