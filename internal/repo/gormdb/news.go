package gormdb

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"newsportal/dao/model"
	"newsportal/internal/repo"
)

type NewsRepo struct {
	db *gorm.DB
}

func NewNewsRepo(db *gorm.DB) *NewsRepo {
	return &NewsRepo{db: db}
}

func (r *NewsRepo) GetNewsByFilters(ctx context.Context, filter repo.NewsFilter) ([]repo.NewsWithCategories, error) {
	newsQuery := r.db.WithContext(ctx).
		Model(&model.News{}).
		Select(`news.*, categories.title AS category_title`).
		Joins(`LEFT JOIN categories ON news."categoryId" = categories."categoryId"`)

	if filter.StatusID > 0 {
		newsQuery = newsQuery.Where(`news."statusId" = ?`, filter.StatusID)
	}
	if filter.CategoryID > 0 {
		newsQuery = newsQuery.Where(`news."categoryId" = ?`, filter.CategoryID)
	}
	if filter.TagID > 0 {
		newsQuery = newsQuery.Where(`? = ANY(news."tagIds")`, filter.TagID)
	}

	newsQuery = newsQuery.Offset((filter.Offset - 1) * filter.Limit).Limit(filter.Limit)
	newsQuery = newsQuery.Order(`news."publicationDate" DESC, news."newsId" DESC`)

	var news []repo.NewsWithCategories
	if err := newsQuery.Find(&news).Error; err != nil {
		log.Printf("Ошибка при получении новостей: %v", err)
		return nil, err
	}

	return news, nil
}

func (r *NewsRepo) CountNewsByFilters(ctx context.Context, filter repo.NewsFilter) (int64, error) {
	var count int64
	newsQuery := r.db.WithContext(ctx).Model(&model.News{})

	if filter.StatusID > 0 {
		newsQuery = newsQuery.Where(`"statusId" = ?`, filter.StatusID)
	}
	if filter.CategoryID > 0 {
		newsQuery = newsQuery.Where(`"categoryId" = ?`, filter.CategoryID)
	}
	if filter.TagID > 0 {
		newsQuery = newsQuery.Where(`? = ANY("tagIds")`, filter.TagID)
	}
	if err := newsQuery.Count(&count).Error; err != nil {
		log.Printf("Ошибка при подсчете новостей: %v", err)
		return 0, err
	}

	return count, nil
}

func (r *NewsRepo) GetNewsByID(ctx context.Context, id int32) (*repo.NewsWithCategories, error) {
	var newsWithCategories repo.NewsWithCategories

	err := r.db.WithContext(ctx).
		Model(&model.News{}).
		Select(`news.*, categories.title AS category_title`).
		Joins(`LEFT JOIN categories ON news."categoryId" = categories."categoryId"`).
		Where(`"newsId" = ?`, id).
		First(&newsWithCategories).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &newsWithCategories, nil
}
