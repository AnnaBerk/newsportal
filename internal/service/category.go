package service

import (
	"context"
	"newsportal/internal/repo"
)

type CategoryDTO struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

type CategoryFilter struct {
	StatusID int32
}

func (s *NewsPortal) GetCategoriesByFilter(ctx context.Context, filter CategoryFilter) ([]*CategoryDTO, error) {
	s.log.Info("Получение категорий по фильтру", "filter", filter)

	// Установка дефолтных значений для фильтров, если не заданы
	//if filter.StatusID <= 0 {
	//	filter.StatusID = 1
	//}

	categories, err := s.CategoryRepo.GetCategoriesByFilter(ctx, repo.CategoryFilter{
		StatusID: filter.StatusID,
	})
	if err != nil {
		s.log.Error("Ошибка при получении категорий", "error", err)
		return nil, err
	}

	var categoryDTOs []*CategoryDTO
	for _, cat := range categories {
		categoryDTOs = append(categoryDTOs, &CategoryDTO{
			ID:    cat.CategoryID,
			Title: cat.Title,
		})
	}

	return categoryDTOs, nil
}
