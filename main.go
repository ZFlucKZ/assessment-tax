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
	"github.com/ZFlucKZ/assessment-tax/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()

	e := echo.New()
	routes.RegisterRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})


	go func() {
		e.Start(":" + os.Getenv("PORT"))
	}()

	fmt.Println("Server is running at port", os.Getenv("PORT"))
	
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
