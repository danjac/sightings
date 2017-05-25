package main

import (
	"fmt"
	"github.com/danjac/sightings/api"
	"github.com/danjac/sightings/config"
	"net/http"
)

func main() {

	cfg, err := config.Configure()
	if err != nil {
		panic(err)
	}

	defer cfg.Close()

	router := api.NewRouter(cfg, "/api/v1")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router); err != nil {
		panic(err)
	}
}
