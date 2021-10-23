package main

import (
	"context"
	"log"
	"os"

	"github.com/bm-krishna/tenant-service/pkg/cmd"
)

func main() {
	log.Println("Tenatn service")
	ctx := context.Background()
	connectionError := cmd.DBConnection(ctx)
	if connectionError != nil {
		log.Fatal("Failed to Connect db", connectionError)
		os.Exit(1)
	}
	err := cmd.ConfigServer(ctx)
	// do graceful shoutdown
	if err != nil {
		log.Fatal("Failed to Config server", os.Stderr, err)
		os.Exit(1)
	}
}
