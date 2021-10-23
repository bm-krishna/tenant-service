package store

import (
	"context"
	"database/sql"

	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DriverConnection type wich have connection props
type DriverConnectionConfig struct {
	Host       string
	Port       string
	UserName   string
	PassWord   string
	SchemaName string
}

var connConfig *DriverConnectionConfig

var pool *sql.DB // Database connection pool.
var once sync.Once

//SQLStore to Setup store by using Connection Contract in driver package
type SQLStore struct {
}

// Connection establish by using singleton pattern connection by using sync.Once
func (provider *SQLStore) New(config map[string]string) *DriverConnectionConfig {
	once.Do(func() {
		connConfig = &DriverConnectionConfig{Host: config["Host"], Port: config["Port"], UserName: config["UserName"], PassWord: config["PassWord"], SchemaName: config["SchemaName"]}
	})
	return connConfig
}

//Provider is implementation for Connection interface in driver
func (provider *SQLStore) Provider(ctx context.Context, driverName string, config map[string]string) error {
	if driverName == "" {
		return errors.New("Provide Store Name")
	}
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		connConfig.UserName, connConfig.PassWord, connConfig.Host, connConfig.Port, connConfig.SchemaName,
	)
	log.Println(dataSourceName)
	// Opening a driver typically will not attempt to connect to the database.
	pool, err = sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		select {
		case <-appSignal:
			stop()
		}
	}()

	// Ping(ctx)
	return nil
}

// Ping the database to verify DSN provided by the user is valid and the
// server accessible. If the ping fails exit the program with an error.
func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}
