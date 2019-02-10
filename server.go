package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"io"
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

// TemplateRenderer allows us to use Echo's templating system.
type TemplateRenderer struct {
	templates *template.Template
}

// Render pre-compiles templates
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	configPath := flag.String("config", "./config/development.json", "path of the config file")
	flag.Parse()

	// Read config
	config, err := config.FromFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Echo, enable pre-compilation of HTML templates.
	e := echo.New()

	t := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	// Set up middlewares.
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(m.DBMiddleware(databases.Init(config)))
	e.Logger.SetLevel(log.INFO)

	// Load routes.
	routes.Init(e)

	// Start server with GraceShutdown.
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
