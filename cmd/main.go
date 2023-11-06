package main

import (
	"log"

	"github.com/SantiagoBedoya/linkedin-go-ms/internal/database"
	"github.com/SantiagoBedoya/linkedin-go-ms/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize Database Client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
