package database

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/ChatService/internal/configs"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
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

func InitializeGraphDB(graphConfig configs.GraphDataBase /*, logger *zap.Logger*/) (neo4j.DriverWithContext, func(), error) {
	fmt.Println("Initializing graph database...")
	ctx := context.Background()
	dbUri := fmt.Sprintf("%v://%v", graphConfig.Database, graphConfig.Host)
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(graphConfig.Username, graphConfig.Password, ""))

	if err != nil {
		fmt.Printf("Failed to connect to Neo4j %v\n", err)
		return nil, nil, err
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		fmt.Printf("Failed to verify connect to Neo4j %v\n", err)
		return nil, nil, err
	}

	cleanup := func() {
		driver.Close(ctx)
	}
	return driver, cleanup, nil
}

func InitializeDB(databaseConfig configs.Database /*, logger *zap.Logger*/) (*sql.DB, func(), error) {
	fmt.Println("Initializing database...")
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

	return tx, nil
}
