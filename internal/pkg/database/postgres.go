package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func NewPostgres() (*sqlx.DB, error) {
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBSslMode := os.Getenv("SSL_MODE")

	DBDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		DBHost, DBUser, DBPassword, DBName, DBPort, DBSslMode,
	)

	db, err := sqlx.Connect("postgres", DBDSN)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(5)
	db.SetConnMaxLifetime(30)

	return db, nil
}
