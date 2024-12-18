package app

import (
	"cinema-seating/config"
	"cinema-seating/docs"
	v1 "cinema-seating/internal/controller/http/v1"
	"cinema-seating/internal/usecase/cinema"
	cinemaRepo "cinema-seating/internal/usecase/cinema/repo"
	"cinema-seating/pkg/httpserver"
	"cinema-seating/pkg/logger"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	ctx = context.Background()
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// HTTP Server
	handler := gin.Default()

	// Configure CORS
	cfgCORS := cors.Config{
		AllowOrigins:     []string{"*"}, // Replace "*" with specific origins for better security
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}

	handler.Use(cors.New(cfgCORS))
	v1Group := handler.Group("/v1")

	// Make Use case
	seatingUseCase := cinema.New(cinemaRepo.NewJSON())

	// Swagger
	docs.SwaggerInfo.BasePath = "/v1"

	// Map routes
	v1.NewSeatingRoutes(v1Group, l, seatingUseCase)

	// Swagger route
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// HTTP Server
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
