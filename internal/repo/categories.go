package repo

import (
	"context"
	"newsportal/dao/model"
)

type CategoryFilter struct {
	StatusID int32
}

type CategoryRepo interface {
	GetCategoriesByFilter(ctx context.Context, filter CategoryFilter) ([]*model.Category, error)
}
