package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/droxey/gogogadget/config"
	"github.com/droxey/gogogadget/databases"
	m "github.com/droxey/gogogadget/middlewares"
	"github.com/droxey/gogogadget/routes"
)

func main() {
	configPath := flag.String("config", "./config/development.json", "path of the config file")
	flag.Parse()

	// Read config
	config, err := config.FromFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(m.DBMiddleware(databases.Init(config)))
	e.Logger.SetLevel(log.INFO)
	routes.Init(e)

	// Start server with GraceShutdown
	go func() {
		if err := e.Start(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)); err != nil {
			e.Logger.Info("shutting down the server.")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
