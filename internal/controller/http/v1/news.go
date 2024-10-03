package v1

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"newsportal/internal/service"
	"strconv"
)

type newsRoutes struct {
	portal *service.NewsPortal
	log    *slog.Logger
}

// Инициализация роутов для новостей
func newNewsRoutes(g *echo.Group, portal *service.NewsPortal, log *slog.Logger) {
	r := &newsRoutes{
		portal: portal,
		log:    log,
	}

	g.GET("", r.getNews)            // /news
	g.GET("/count", r.getNewsCount) // /news/count
	g.GET("/:id", r.getNewsByID)    // /news/:id
}

type GetNewsParams struct {
	Page       int   `query:"page" validate:"omitempty,gte=1"`
	PageSize   int   `query:"pageSize" validate:"omitempty,gte=1"`
	CategoryID int32 `query:"categoryId" validate:"omitempty,gt=0"`
	TagID      int32 `query:"tagId" validate:"omitempty,gt=0"`
}

// Контроллер для получения списка новостей
func (r *newsRoutes) getNews(c echo.Context) error {
	var params GetNewsParams

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверные параметры запроса"})
	}

	if err := c.Validate(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Устанавливаем значения по умолчанию, если они не переданы
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 10
	}

	news, err := r.portal.GetNewsByFilter(c.Request().Context(), service.NewsFilter{
		Page:       params.Page,
		PageSize:   params.PageSize,
		CategoryID: params.CategoryID,
		TagID:      params.TagID,
	})
	if err != nil {
		r.log.Error("Ошибка при получении новостей: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Внутренняя ошибка сервера"})
	}

	return c.JSON(http.StatusOK, news)
}

// Контроллер для получения новости по ID
func (r *newsRoutes) getNewsByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.log.Error("Некорректный параметр ID", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Некорректный ID")
	}

	news, err := r.portal.GetNewsByID(c.Request().Context(), int32(id))
	if err != nil {
		r.log.Error("Ошибка при получении новости", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении новости")
	}

	if news == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Новость не найдена")
	}
	return c.JSON(http.StatusOK, news)
}

type GetNewsCountParams struct {
	CategoryID int32 `query:"categoryId" validate:"omitempty,gt=0"`
	TagID      int32 `query:"tagId" validate:"omitempty,gt=0"`
}

// Контроллер для подсчета новостей
func (r *newsRoutes) getNewsCount(c echo.Context) error {
	var params GetNewsCountParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверные параметры запроса"})
	}

	if err := c.Validate(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	count, err := r.portal.GetNewsCountByFilter(c.Request().Context(), service.NewsFilter{
		CategoryID: params.CategoryID,
		TagID:      params.TagID,
	})
	if err != nil {
		r.log.Error("Ошибка при подсчете новостей", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при подсчете новостей")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"count": count,
	})
}
