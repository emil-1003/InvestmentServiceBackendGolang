package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type Config struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     string
}

func ConnectToDb() error {
	c := Config{
		Name:     "name",     // Type database name
		Username: "username", // Type username
		Password: "password", // Type password
		Host:     "host",     // Type host
		Port:     "port",     // Type port
	}

	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.Username, c.Password, c.Host, c.Port, c.Name)

	var err error
	DB, err = sql.Open("mysql", dbInfo)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}

	return err
}
