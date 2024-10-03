package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log/slog"
	v1 "newsportal/internal/controller/http/v1"
	"newsportal/internal/repo/gormdb"
	"newsportal/internal/service"
	"newsportal/pkg/httpserver"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Run() {
	// Configuration
	cfg := MustLoad()

	// Logger
	log := setupLogger(cfg.Env)

	log.Info(
		"starting newsportal",
		slog.String("env", cfg.Env),
		slog.String("version", "1.0.0"),
	)
	log.Debug("debug messages are enabled")

	// Initialize the database
	db, err := initDB(*cfg)
	if err != nil {
		log.Error("Ошибка инициализации базы данных", err)
		os.Exit(1)
	}
	db = db.Debug()
	sqlDB, err := db.DB()
	if err != nil {
		log.Error("Ошибка получения подключения к базе данных: %v", err)
	}
	defer sqlDB.Close()

	// Initialize repositories
	log.Info("Initializing repositories...")
	repositories := service.Repositories{
		CategoryRepo: gormdb.NewCategoryRepo(db),
		TagRepo:      gormdb.NewTagRepo(db),
		NewsRepo:     gormdb.NewNewsRepo(db),
	}

	// Initialize services
	log.Info("Initializing services...")
	services := service.NewNewsPortal(repositories, *log)

	// Echo handler
	log.Info("Initializing handlers and routes...")
	handler := echo.New()

	handler.Validator = &CustomValidator{validator: validator.New()}
	v1.NewRouter(handler, services, log)

	// HTTP server
	log.Info("Starting http server...")
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Port))

	// Waiting signal for graceful shutdown
	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error("app - Run - httpServer.Notify", err)
	}

	// Graceful shutdown
	log.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error("app - Run - httpServer.Shutdown", err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func initDB(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.ConnectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}
	return db, nil
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
