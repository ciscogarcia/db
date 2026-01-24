package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var conn *pgxpool.Pool

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func (d *DBConfig) GetURL() string {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", d.Username, d.Password, d.Host, d.Port, d.Database)
	return url
}

func GetDB() *pgxpool.Pool {
	if conn != nil {
		return conn
	}
	return InitDB()
}

func InitDB() *pgxpool.Pool {
	// Get Credentials from config
	file_path, err := os.Getwd()
	if err != nil {
		// log error
	}
	file_path += "/config.json"

	// TODO: stat file
	// use alternate paths if stat fails
	file, err := os.Open(file_path)
	if err != nil {
		// log error
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	dbConfig := DBConfig{}
	err = decoder.Decode(&dbConfig)
	if err != nil {
		// log error
	}
	// Get connection
	conn, err = pgxpool.New(context.Background(), dbConfig.GetURL())
	if err != nil {
		// log err
	}

	// Return connection
	return conn
}
