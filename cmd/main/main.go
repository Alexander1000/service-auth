package main

import (
	"log"
	"flag"
	"net/http"
	"fmt"
	"context"

	"github.com/Alexander1000/service-auth/internal/config"
	"github.com/Alexander1000/service-auth/internal/trap"
	"github.com/Alexander1000/service-auth/internal/database"
	"github.com/Alexander1000/service-auth/internal/storage"
	"github.com/Alexander1000/service-auth/internal/api/v1/registration"
	"github.com/Alexander1000/service-auth/internal/api/v1/authenticate"
	"github.com/Alexander1000/service-auth/internal/api/v1/authorize"
)

func main() {
	log.Println("starting application")

	configPath := flag.String("c", "", "config file")
	flag.Parse()

	if len(*configPath) == 0 {
		log.Fatalf("unknown config file")
	}

	var err error
	var cfg *config.Config
	if cfg, err = config.LoadFromFile(*configPath); err != nil {
		log.Fatalf("error in load config from file: %v", err)
	}

	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("error in connect to database: %v", err)
	}
	defer db.Close()

	strg := storage.New(db)

	http.Handle("/v1/registration", registration.New(strg))

	http.Handle("/v1/authenticate", authenticate.New(strg))

	http.Handle("/v1/authorize", authorize.New())

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil); err != nil {
			log.Fatalf("error in start application: %v", err)
		}
	}()

	signalTrap := trap.NewTrap()
	ctx := context.Background()

	if err := signalTrap.WaitShutdown(ctx); err != nil {
		log.Printf("error in caught signal: %v", err)
	}

	log.Println("application terminated")
}
