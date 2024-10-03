package service

import (
	"context"
	"newsportal/internal/repo/gormdb"
)

type TagDTO struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

type TagFilter struct {
	StatusID int32
}

func (s *NewsPortal) GetTagsByFilter(ctx context.Context, filter TagFilter) ([]*TagDTO, error) {
	s.log.Info("Получение тегов по фильтру", "filter", filter)

	// Установка дефолтных значений для фильтров, если не заданы
	//if filter.StatusID <= 0 {
	//	filter.StatusID = 1
	//}

	tags, err := s.TagRepo.GetTagsByFilter(ctx, gormdb.TagFilter{
		StatusID: filter.StatusID,
	})
	if err != nil {
		s.log.Error("Ошибка при получении тегов", "error", err)
		return nil, err
	}

	var tagDTOs []*TagDTO
	for _, tag := range tags {
		tagDTOs = append(tagDTOs, &TagDTO{
			ID:    tag.TagID,
			Title: tag.Title,
		})
	}

	return tagDTOs, nil
}
