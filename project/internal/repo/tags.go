package repo

import (
	"context"
	"newsportal/dao/model"
)

type TagRepo interface {
	GetPublishedTags(ctx context.Context) ([]model.Tag, error)
}
