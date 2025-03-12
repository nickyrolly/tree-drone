package main

import (

	// "github.com/SawitProRecruitment/UserService/handler"
	// "github.com/SawitProRecruitment/UserService/repository"

	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	generated "github.com/nickyrolly/tree-drone/generated"
	"github.com/nickyrolly/tree-drone/handler"
	"github.com/nickyrolly/tree-drone/internal/config"
	"github.com/nickyrolly/tree-drone/repository"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()

	var server generated.ServerInterface = newServer(cfg)
	generated.RegisterHandlers(e, server)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":" + cfg.GetString("application.port")))
}

func newServer(cfg *viper.Viper) *handler.Server {
	url := cfg.GetString("database.url")
	if url == "" {
		url = os.Getenv("DATABASE_URL")
	}

	fmt.Println("DB Dsn : ", url)
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Driver: cfg.GetString("database.driver"),
		Url:    url,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}
