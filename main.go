package main

import (
	"log"

	"github.com/emil-1003/InvestmentServiceBackendGolang/pkg/database"
	"github.com/emil-1003/InvestmentServiceBackendGolang/pkg/server"
)

const (
	apiPath    = "api"
	apiVersion = "v1"
	apiPort    = ":8585"
	apiName    = "Investment Service API"
)

func init() {
	// Start db connection
	if err := database.ConnectToDb(); err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
}

func main() {
	// Start server
	srv, err := server.New(apiName, apiVersion, apiPort, apiPath)
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}

	log.Printf("Starting %s version %s, listening on %s", srv.Name, srv.Version, srv.Port)
	log.Fatal(srv.ListenAndServe())
}
