package service

import (
	"context"
	"newsportal/internal/repo"
	"newsportal/internal/repo/gormdb"
	"slices"
)

type NewsFilter struct {
	Page       int
	PageSize   int
	CategoryID int32
	TagID      int32
}

type ShortNews struct {
	ID              int32    `json:"id"`
	Title           string   `json:"title"`
	Foreword        string   `json:"foreword"`
	CategoryTitle   string   `gorm:"-" json:"categoryTitle"`
	TagTitles       []string `gorm:"-" json:"tagTitles"`
	PublicationDate string   `json:"publicationDate"`
}

type FullNews struct {
	ID              int32    `json:"id"`
	Title           string   `json:"title"`
	Foreword        string   `json:"foreword"`
	Content         string   `json:"content"`
	CategoryTitle   string   `gorm:"-" json:"categoryTitle"`
	TagTitles       []string `gorm:"-" json:"tagTitles"`
	PublicationDate string   `json:"publicationDate"`
}

func uniqueInt32Slice(input []int32) []int32 {
	slices.Sort(input)
	return slices.Compact(input)
}

func (s *NewsPortal) GetNewsByFilter(ctx context.Context, filter NewsFilter) ([]*ShortNews, error) {
	s.log.Info("Получение новостей по фильтру", "filter", filter)

	news, err := s.NewsRepo.GetNewsByFilters(ctx, gormdb.NewsFilter{
		Offset:     filter.Page,
		Limit:      filter.PageSize,
		CategoryID: filter.CategoryID,
		TagID:      filter.TagID,
		StatusID:   int32(repo.TagStatusPublished),
	})
	if err != nil {
		s.log.Error("Ошибка при получении новостей", "error", err)
		return nil, err
	}

	var tagIDs []int32
	for _, n := range news {
		tagIDs = append(tagIDs, n.TagIds...)
	}

	tagIDs = uniqueInt32Slice(tagIDs)

	tagTitles, err := s.TagRepo.GetTagTitles(ctx, tagIDs)
	if err != nil {
		s.log.Error("Ошибка при получении названий тегов", "error", err)
		return nil, err
	}

	var newsDTOs []*ShortNews
	for _, n := range news {
		var tagNames []string
		for _, tagID := range n.TagIds {
			if tagTitle, ok := tagTitles[tagID]; ok {
				tagNames = append(tagNames, tagTitle)
			}
		}

		newsDTOs = append(newsDTOs, &ShortNews{
			ID:              n.NewsID,
			Title:           n.Title,
			Foreword:        n.Foreword,
			CategoryTitle:   n.CategoryTitle,
			TagTitles:       tagNames,
			PublicationDate: n.PublicationDate.Format("2006-01-02"),
		})
	}

	return newsDTOs, nil
}

func (s *NewsPortal) GetNewsCountByFilter(ctx context.Context, filter NewsFilter) (int64, error) {
	s.log.Info("Подсчет новостей по фильтру", "filter", filter)

	count, err := s.NewsRepo.CountNewsByFilters(ctx, gormdb.NewsFilter{
		CategoryID: filter.CategoryID,
		TagID:      filter.TagID,
	})
	if err != nil {
		s.log.Error("Ошибка при подсчете новостей", "error", err)
		return 0, err
	}
	return count, nil
}

func (s *NewsPortal) GetNewsByID(ctx context.Context, id int32) (*FullNews, error) {
	s.log.Info("Получение новости по ID", "id", id)

	news, err := s.NewsRepo.GetNewsByID(ctx, id)
	if err != nil {
		s.log.Error("Ошибка при получении новости", "error", err)
		return nil, err
	}

	if news == nil {
		s.log.Warn("Новость не найдена", "id", id)
		return nil, ErrNewsNotFound
	}

	tagIDs := news.TagIds

	var tagNames []string
	if len(tagIDs) > 0 {
		tagTitles, err := s.TagRepo.GetTagTitles(ctx, tagIDs)
		if err != nil {
			s.log.Error("Ошибка при получении названий тегов", "error", err)
			return nil, err
		}

		for _, tagID := range tagIDs {
			if tagTitle, ok := tagTitles[tagID]; ok {
				tagNames = append(tagNames, tagTitle)
			}
		}
	}

	newsDTO := &FullNews{
		ID:              news.NewsID,
		Title:           news.Title,
		Foreword:        news.Foreword,
		Content:         news.Content,
		CategoryTitle:   news.CategoryTitle,
		TagTitles:       tagNames,
		PublicationDate: news.PublicationDate.Format("2006-01-02"),
	}

	return newsDTO, nil
}
