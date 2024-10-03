package repo

import (
	"context"
	"github.com/lib/pq"
	"time"
)

type NewsFilter struct {
	CategoryID int32
	TagID      int32
	StatusID   int32
	Limit      int
	Offset     int
}

type NewsWithCategories struct {
	NewsID          int32         `gorm:"column:newsId"`
	Title           string        `gorm:"column:title"`
	Foreword        string        `gorm:"column:foreword"`
	Content         string        `gorm:"column:content"`
	PublicationDate time.Time     `gorm:"column:publicationDate"`
	CategoryID      int32         `gorm:"column:categoryId"`
	TagIds          pq.Int32Array `gorm:"column:tagIds"`
	CategoryTitle   string        `gorm:"column:category_title"`
}

type NewsRepo interface {
	GetNewsByFilters(ctx context.Context, filter NewsFilter) ([]NewsWithCategories, error)
	CountNewsByFilters(ctx context.Context, filter NewsFilter) (int64, error)
	GetNewsByID(ctx context.Context, id int32) (*NewsWithCategories, error)
}
