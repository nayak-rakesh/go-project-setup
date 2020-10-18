package config

import (
	"fmt"
	"os"
)

type dbConfig struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
}

// GetDBConnString ...
func GetDBConnString() string {
	c := dbConfig{
		dbUser:     os.Getenv("PGSQL_USER"),
		dbPassword: os.Getenv("PGSQL_PASS"),
		dbHost:     os.Getenv("PGSQL_HOST"),
		dbPort:     os.Getenv("PGSQL_PORT"),
		dbName:     os.Getenv("PGSQL_DB"),
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPassword,
		c.dbHost,
		c.dbPort,
		c.dbName,
	)
}
