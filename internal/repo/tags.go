package repo

import (
	"context"
	"newsportal/dao/model"
)

type TagFilter struct {
	StatusID int32
}

type TagRepo interface {
	GetTagsByFilter(ctx context.Context, filter TagFilter) ([]*model.Tag, error)
	GetTagTitles(ctx context.Context, tagIDs []int32) (map[int32]string, error)
}
