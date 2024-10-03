package service

import (
	"log/slog"
	"newsportal/internal/repo"
)

type NewsPortal struct {
	NewsRepo     repo.NewsRepo
	CategoryRepo repo.CategoryRepo
	TagRepo      repo.TagRepo
	log          slog.Logger
}

// Конструктор для создания сервиса
func NewNewsPortal(repos repo.Repositories, logger slog.Logger) *NewsPortal {
	return &NewsPortal{
		NewsRepo:     repos.NewsRepo,
		CategoryRepo: repos.CategoryRepo,
		TagRepo:      repos.TagRepo,
		log:          logger,
	}
}
