// This file contains the repository implementation layer.
package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	Options NewRepositoryOptions
	Gorm    *gorm.DB
}

type NewRepositoryOptions struct {
	Driver string
	Url    string
}

type RepositoryInterface interface {
	SetEstate(estate *Estate) error
	GetOptions() NewRepositoryOptions // Add this method to get the options.
	CreateEstate(e *Estate) error
}

func (r *Repository) GetOptions() NewRepositoryOptions {
	return r.Options
}

func (r *Repository) SetEstate(estate *Estate) error {
	return nil
}
func (r *Repository) CreateEstate(e *Estate) error {
	result := r.Gorm.Create(e)
	if result.Error != nil {
		log.Println("Error : ", result.Error)
		return result.Error
	}
	log.Println("Success")
	return nil
}

func NewRepository(opts NewRepositoryOptions) RepositoryInterface {
	fmt.Println("Driver : ", opts.Driver)
	fmt.Println("Url : ", opts.Url)
	gormDB, err := gorm.Open(getDialector(opts.Driver, opts.Url), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

	return &Repository{
		Options: opts,
		Gorm:    gormDB,
	}
}

func getDialector(driver, url string) gorm.Dialector {
	switch driver {
	case "postgres":
		return postgres.Open(url)
	default:
		// return postgres.Open(url)
		panic(fmt.Sprintf("driver '%s' not supported", driver)) // Now will panic with the driver
	}
}
