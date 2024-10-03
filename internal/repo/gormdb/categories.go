package gormdb

import (
	"context"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"log"
	"newsportal/internal/repo/model"
	"newsportal/internal/repo/query"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

type CategoryFilter struct {
	StatusID int32
}

func (r *CategoryRepo) GetCategoriesByFilter(ctx context.Context, filter CategoryFilter) ([]*model.Category, error) {
	categoryQuery := query.Use(r.db).Category.WithContext(ctx)
	titleField := field.NewString(categoryQuery.TableName(), "orderNumber")
	statusField := field.NewInt32(categoryQuery.TableName(), "statusId")

	q := categoryQuery

	if filter.StatusID > 0 {
		q = q.Where(statusField.Eq(filter.StatusID))
	}

	categories, err := q.Order(titleField.Asc()).Find()
	if err != nil {
		log.Printf("Ошибка при получении категорий: %v", err)
		return nil, err
	}

	return categories, nil
}
