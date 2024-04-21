package config

import (
	"errors"
	"fmt"
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDatabase() (*gorm.DB, error) {
	logger := GetLogger("PG")

	var db *gorm.DB
	var err error

	env := os.Getenv("ENV")

	switch env {
	case "prod":
		db, err = initPostgres(logger)
	case "dev":
		db, err = initSQLite(logger)
	default:
		return nil, errors.New("environment not specified or invalid")
	}

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&schemas.User{}, &schemas.Category{}, &schemas.Task{})

	if err != nil {
		logger.ErrF("Auto migration failed: %s", err)
		return nil, err
	}

	return db, nil
}

func initPostgres(logger *Logger) (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrF("Postgres connection failed: %s", err)
		return nil, err
	}
	return db, nil
}

func initSQLite(logger *Logger) (*gorm.DB, error) {
	dbPath := "C:\\PersonalProjects\\routine\\routine_back\\db\\test.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Creating database ...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}
		err = file.Close()
		if err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrF("SQLite connection failed: %s", err)
		return nil, err
	}
	return db, nil
}
