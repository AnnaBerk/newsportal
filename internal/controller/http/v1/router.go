package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"newsportal/internal/service"
)

func NewRouter(handler *echo.Echo, services *service.NewsPortal, log *slog.Logger) {
	handler.Use(middleware.Recover())

	v1 := handler.Group("/api/v1")
	{
		newNewsRoutes(v1.Group("/news"), services, log)
		newCategoryRoutes(v1.Group("/categories"), services, log)
		newTagRoutes(v1.Group("/tags"), services, log)
	}
}
