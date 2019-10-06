package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// DBStruct to ensure database configuration
type DBStruct struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// ConnectToDB to connect to database
func ConnectToDB(stageEnv string) (*sql.DB, error) {
	port := os.Getenv(fmt.Sprintf("%s_DB_PORT", strings.ToUpper(stageEnv)))
	host := os.Getenv(fmt.Sprintf("%s_DB_HOST", strings.ToUpper(stageEnv)))
	user := os.Getenv(fmt.Sprintf("%s_DB_USER", strings.ToUpper(stageEnv)))
	password := os.Getenv(fmt.Sprintf("%s_DB_PASSWORD", strings.ToUpper(stageEnv)))
	name := os.Getenv(fmt.Sprintf("%s_DB_NAME", strings.ToUpper(stageEnv)))

	dbConfig := DBStruct{
		Port:     port,
		Host:     host,
		User:     user,
		Password: password,
		Name:     name,
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
