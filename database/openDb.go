package database

import (
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"database/sql"
	"fmt"
)

type Config struct {
	User     string
	Host     string
	Password string
	Port     string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	//Open the database and get the DB descriptor
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s host=%s password=%s port=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Host, cfg.Password, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, errors.Wrap(err, "Open DB, func NewPostgresDB")
	}
	//Checking the connection to the DB
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "Checking DB connection, func NewPostgresDB")
	}

	return db, nil
}
