package v1

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"newsportal/internal/service"
	"strconv"
)

type tagRoutes struct {
	portal *service.NewsPortal
	log    *slog.Logger
}

// Инициализация роутов для тегов
func newTagRoutes(g *echo.Group, portal *service.NewsPortal, log *slog.Logger) {
	r := &tagRoutes{
		portal: portal,
		log:    log,
	}

	g.GET("", r.getTags) // /tags
}

func (r *tagRoutes) getTags(c echo.Context) error {
	statusID, _ := strconv.Atoi(c.QueryParam("statusId"))

	tags, err := r.portal.GetTagsByFilter(c.Request().Context(), service.TagFilter{
		StatusID: int32(statusID),
	})
	if err != nil {
		r.log.Error("Ошибка при получении тегов", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении тегов")
	}

	return c.JSON(http.StatusOK, tags)
}
