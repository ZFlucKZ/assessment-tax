package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ZFlucKZ/assessment-tax/config"
	"github.com/ZFlucKZ/assessment-tax/db"
	"github.com/ZFlucKZ/assessment-tax/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.SetupEnv()

	p := config.ConnectDB()

	db.SetDatabase(p)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	go func() {
		e.Start(":" + fmt.Sprintf("%d", config.Env.Port))
	}()

	fmt.Println("Server is running at port", fmt.Sprintf("%d", config.Env.Port))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	fmt.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	fmt.Println("Server gracefully stopped")

	defer cancel()
}
