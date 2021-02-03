package main

import (
	"log"
	"os"

	"github.com/bm-krishna/tenant-service/pkg/cmd"
)

func main() {
	log.Println("Tenatn service")
	err := cmd.ConfigServer()
	// do graceful shoutdown
	if err != nil {
		log.Fatal("Failed to Config server", os.Stderr, err)
		os.Exit(1)
	}
}
