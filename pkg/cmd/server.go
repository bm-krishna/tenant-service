package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bm-krishna/tenant-service/pkg/protocal/tenant"
	"github.com/bm-krishna/tenant-service/pkg/repository/store"
	"github.com/bm-krishna/tenant-service/pkg/server"
)

// Config ...
type config struct {
	Port string
}

// ConfigServer will take the config from env and start the serve
func ConfigServer(ctx context.Context) error {
	fmt.Println("ConfigServer")
	// ctx := context.Background()
	tenantServer := &server.TenantServer{}
	port := "7575"
	portEnv := os.Getenv("SERVER_PORT")
	if portEnv != "" {
		port = portEnv
	}
	config := &config{
		Port: port,
	}
	err := tenant.BootStrapServer(ctx, tenantServer, config.Port)
	log.Println("server running on port" + port)
	if err != nil {
		log.Fatal("Failed to BootStrap Server")
		return err
	}
	return nil
}

// BuildConnectionFromEnv returns connection object either evn props or default connection
func buildConnectionFromEnv() map[string]string {
	defaultConnection := true
	// connection data from .env file
	dbhost := os.Getenv("DB_HOST")
	if dbhost != "" {
		defaultConnection = false
	}
	dbport := os.Getenv("DB_PORT")
	if dbport != "" {
		defaultConnection = false
	}
	dbUserName := os.Getenv("DB_USER")
	if dbUserName != "" {
		defaultConnection = false
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword != "" {
		defaultConnection = false
	}
	dbSchemaName := os.Getenv("DB_NAME")
	if dbSchemaName != "" {
		defaultConnection = false
	}
	if defaultConnection {
		log.Println("Falling back to default connection")
		return map[string]string{
			"Host":       "localhost",
			"Port":       "3306",
			"UserName":   "root",
			"PassWord":   "12345678",
			"SchemaName": "tenant",
		}
	}
	return map[string]string{
		"Host":       dbhost,
		"Port":       dbport,
		"UserName":   dbUserName,
		"PassWord":   dbPassword,
		"SchemaName": dbSchemaName,
	}
}

//DBConnection ...
func DBConnection(ctx context.Context) error {
	var store store.SQLStore
	storeName := "mysql"
	dirverName := os.Getenv("DRIVER_NAME")
	if dirverName != "" {
		storeName = dirverName
	}
	connection := buildConnectionFromEnv()
	err := store.Provider(ctx, storeName, connection)
	if err != nil {
		return err
	}
	return nil
}
