package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nickyrolly/tree-drone/internal/config"

	repository "github.com/nickyrolly/tree-drone/repository"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	e.Use(middleware.Logger())

	url := cfg.GetString("database.url")
	if url == "" {
		url = os.Getenv("DATABASE_URL")
	}

	env := os.Getenv("ENV")

	repo := repository.NewRepository(repository.NewRepositoryOptions{
		Driver: cfg.GetString("database.driver"),
		Url:    url,
	})

	db := repo.Gorm

	models := []interface{}{
		&repository.Estate{},
		&repository.EstateTree{},
	}

	for _, model := range models {
		if err := db.Migrator().AutoMigrate(model); err != nil {
			e.Logger.Error("Migration error for %T: %+v", model, err)
			return
		}
		e.Logger.Infof("Successfully migrated %T", model)
	}

	if env == "development" {
		var IRepo repository.RepositoryInterface

		estate := repository.Estate{}
		estate.SetLength(5)
		estate.SetWidth(10)

		IRepo = repo
		IRepo.SetEstate(estate)

		e.Logger.Infof("Successfully created seed estate")

	}

}
