package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/vaibhav-ch123/asset-management/database"
	"github.com/vaibhav-ch123/asset-management/server"
)

const (
	ShutDownTimeOut = 5 * time.Minute
)

func main() {

	if envErr := godotenv.Load(); envErr != nil {
		log.Fatalf("failed to load .env file: %+v", envErr)
	}

	sigChn := make(chan os.Signal, 1)

	signal.Notify(sigChn, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := server.SetupRoute()

	if err := database.ConnectAndMigrate(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	); err != nil {
        log.Fatalf("failed to connect database: %+v", err)
	}
	fmt.Println("Database connected!")

	fmt.Println("Server Running..")
	go func() {
		srvErr := srv.Run(":" + os.Getenv("SERVER_PORT"))
		if srvErr == http.ErrServerClosed {
			fmt.Println("Server stopped!")
		} else if srvErr != nil {
			log.Fatalf("failed to run server: %+v", srvErr)
		}

	}()

	<-sigChn

	if srvErr := srv.ShutDownServer(ShutDownTimeOut); srvErr != nil {
		log.Fatalf("failed to shutdown server: %+v", srvErr)
	}

}
