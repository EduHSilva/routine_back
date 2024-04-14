package config

import (
	"errors"
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

	err = db.AutoMigrate(&schemas.User{})

	if err != nil {
		logger.ErrF("Auto migration failed: %s", err)
		return nil, err
	}

	return db, nil
}

func initPostgres(logger *Logger) (*gorm.DB, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
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
