package gormdb

import (
	"context"
	"gorm.io/gorm"
	"log"
	"newsportal/dao/model"
	"newsportal/dao/query"
)

type TagRepo struct {
	db *gorm.DB
}

// NewTagRepo создаёт новый экземпляр репозитория тегов
func NewTagRepo(db *gorm.DB) *TagRepo {
	return &TagRepo{db: db}
}

func (r *TagRepo) GetTags(ctx context.Context) ([]*model.Tag, error) {
	tagQuery := query.Use(r.db).Tag.WithContext(ctx)

	t, err := tagQuery.Find()
	if err != nil {
		log.Printf("Ошибка при получении тегов: %v", err)
		return nil, err
	}

	return t, nil
}
