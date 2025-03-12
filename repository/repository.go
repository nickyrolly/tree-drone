// This file contains the repository implementation layer.
package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	// Db   *sql.DB
	Gorm *gorm.DB
}

type NewRepositoryOptions struct {
	Driver string
	Url    string
	Host   string
	// Port            int
	// SSLMode         bool
	DBName string
	// Username        string
	// Password        string
	// MaxIdleConns    int
	// MaxOpenConns    int
	// ConnMaxLifetime time.Duration
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	// var cfg = config.Get()

	fmt.Println("")

	var gormDB *gorm.DB
	switch dbDriver := opts.Driver; dbDriver {
	case "postgresql":
		gormDB = NewPostgreSQL(opts)
	case "pgx":
		gormDB = NewPostgreSQL(opts)
	default:
		gormDB = NewPostgreSQL(opts)
	}

	return &Repository{
		Gorm: gormDB,
	}
}

func NewPostgreSQL(opts NewRepositoryOptions) *gorm.DB {
	fmt.Println("-- Postgre SQL --")
	fmt.Println("Driver : ", opts.Driver)
	fmt.Println("Url : ", opts.Url)
	connection, err := sql.Open(opts.Driver, opts.Url)
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	connection.SetMaxIdleConns(5)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	connection.SetMaxOpenConns(50)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	connection.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: connection,
	}), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		// TODO: set logger by environment
		// Logger:                 logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("--SUCCESS KONEK : Postgre SQL --")

	return db
}

// func NewLite(opts NewRepositoryOptions) *gorm.DB {
// 	// var cfg = config.Get()

// 	db, err := gorm.Open(sqlite.Open(opts.DBName), &gorm.Config{
// 		PrepareStmt:            true,
// 		SkipDefaultTransaction: true,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return db
// }
