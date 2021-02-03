package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bm-krishna-source/tenant-service/pkg/protocal/tenant"
	"github.com/bm-krishna-source/tenant-service/pkg/server"
)

type Config struct {
	Port string
}

func ConfigServer() error {
	fmt.Println("ConfigServer")
	ctx := context.Background()
	tenantServer := &server.TenantServer{}
	port := "7575"
	portEnv := os.Getenv("SERVER_PORT")
	if portEnv != "" {
		port = portEnv
	}
	config := &Config{
		Port: port,
	}
	err := tenant.BootStrapServer(ctx, tenantServer, config.Port)
	if err != nil {
		log.Fatal("Failed to BootStrap Server")
		return err
	}
	return nil
}
