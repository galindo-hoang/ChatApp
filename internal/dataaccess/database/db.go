package database

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/ChatService/internal/configs"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	//go:embed migrations/mysql/0001.initialize.sql
	byteMigrate []byte
)

func InitializeAndMigrateUpDB(databaseConfig configs.Database /*, logger *zap.Logger*/) (*sql.DB, func(), error) {
	var host = os.Getenv("DB_HOST")
	if len(host) != 0 {
		databaseConfig.Host = host
	}
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	db, err := sql.Open(string(databaseConfig.Type), connectionString)
	if err != nil {
		log.Printf("error connecting to database: %+v", err)
		return nil, nil, err
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, err
}

func migrateUp(context context.Context, db *gorm.DB) error {
	if err := db.AutoMigrate(Accounts{}, Messages{}); err != nil {
		return err
	}
	return nil
}

func InitializeGorm(db *sql.DB, databaseConfig configs.Database) (*gorm.DB, error) {
	var tx *gorm.DB
	var err error

	switch databaseConfig.Type {
	case "mysql":
		tx, err = gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	case "postgresql":
		tx, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported database type %s\n", databaseConfig.Type)
	}

	if err != nil {
		return nil, err
	}

	err = migrateUp(context.Background(), tx)
	if err != nil {
		log.Printf("error migrating database: %+v", err)
	}
	return tx, nil
}
