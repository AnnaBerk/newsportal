package v1

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"newsportal/internal/service"
	"strconv"
)

type categoryRoutes struct {
	portal *service.NewsPortal
	log    *slog.Logger
}

// Инициализация роутов для категорий
func newCategoryRoutes(g *echo.Group, portal *service.NewsPortal, log *slog.Logger) {
	r := &categoryRoutes{
		portal: portal,
		log:    log,
	}

	g.GET("", r.getCategories) // /categories
}

func (r *categoryRoutes) getCategories(c echo.Context) error {
	statusID, _ := strconv.Atoi(c.QueryParam("statusId"))

	categories, err := r.portal.GetCategoriesByFilter(c.Request().Context(), service.CategoryFilter{
		StatusID: int32(statusID),
	})
	if err != nil {
		r.log.Error("Ошибка при получении категорий", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении категорий")
	}

	return c.JSON(http.StatusOK, categories)
}
