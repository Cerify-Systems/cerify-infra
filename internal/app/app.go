package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cerify/cerify-infra/internal/config"
	"github.com/cerify/cerify-infra/internal/handlers"
	"github.com/cerify/cerify-infra/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type App struct {
	config     *config.Config
	logger     *zap.Logger
	server     *http.Server
	mongodb    *mongo.Client
	services   *services.Services
}

func New(cfg *config.Config) (*App, error) {
	logger, _ := zap.NewProduction()

	// Initialize MongoDB connection
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoDB.URI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Initialize services
	services, err := services.New(cfg, mongoClient, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize services: %v", err)
	}

	// Initialize router
	router := gin.Default()
	handlers.RegisterRoutes(router, services, cfg)

	// Configure HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	return &App{
		config:   cfg,
		logger:   logger,
		server:   server,
		mongodb:  mongoClient,
		services: services,
	}, nil
}

func (a *App) Start() error {
	a.logger.Info("Starting server", zap.Int("port", a.config.Server.Port))
	return a.server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %v", err)
	}

	if err := a.mongodb.Disconnect(ctx); err != nil {
		return fmt.Errorf("mongodb disconnect failed: %v", err)
	}

	return nil
} 