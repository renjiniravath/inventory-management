package main

import (
	"app/core/logger"
	"app/handler"
	"app/services"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if err := godotenv.Load(); err != nil {
		logger.Error.Fatalf("No .env file found")
	}
	services.GetMySqlConnection()

	e.Static("/", "public")

	e.GET("/item", handler.GetInventoryList)
	e.POST("/item", handler.CreateItem)
	e.PUT("/item/:id", handler.EditItem)
	e.DELETE("/item/:id", handler.DeleteItem)

	e.GET("/warehouse", handler.GetWarehouseList)
	e.POST("/warehouse", handler.CreateWarehouse)
	e.PUT("/warehouse/:id", handler.EditWarehouse)
	e.DELETE("/warehouse/:id", handler.DeleteWarehouse)

	//Starting the api server
	go func() {
		if err := e.Start(":8080"); err != nil {
			logger.Error.Fatalf("Shutting down the server : %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Error.Fatal(err)
	}
}
