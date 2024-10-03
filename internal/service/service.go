package service

import (
	"log/slog"
	"newsportal/internal/repo/gormdb"
)

type NewsPortal struct {
	NewsRepo     *gormdb.NewsRepo
	CategoryRepo *gormdb.CategoryRepo
	TagRepo      *gormdb.TagRepo
	log          slog.Logger
}

type Repositories struct {
	NewsRepo     *gormdb.NewsRepo
	CategoryRepo *gormdb.CategoryRepo
	TagRepo      *gormdb.TagRepo
}

// Конструктор для создания сервиса
func NewNewsPortal(repos Repositories, logger slog.Logger) *NewsPortal {
	return &NewsPortal{
		NewsRepo:     repos.NewsRepo,
		CategoryRepo: repos.CategoryRepo,
		TagRepo:      repos.TagRepo,
		log:          logger,
	}
}
