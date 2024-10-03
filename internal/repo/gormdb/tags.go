package gormdb

import (
	"context"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"log"
	"newsportal/dao/model"
	"newsportal/dao/query"
	"newsportal/internal/repo"
)

type TagRepo struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) *TagRepo {
	return &TagRepo{db: db}
}

func (r *TagRepo) GetTagsByFilter(ctx context.Context, filter repo.TagFilter) ([]*model.Tag, error) {
	tagQuery := query.Use(r.db).Tag.WithContext(ctx)
	titleField := field.NewString(tagQuery.TableName(), "title")
	statusField := field.NewInt32(tagQuery.TableName(), "statusId")

	q := tagQuery

	if filter.StatusID > 0 {
		q = q.Where(statusField.Eq(filter.StatusID))
	}

	tags, err := q.Order(titleField.Asc()).Find()
	if err != nil {
		log.Printf("Ошибка при получении тегов: %v", err)
		return nil, err
	}

	return tags, nil
}

func (r *TagRepo) GetTagTitles(ctx context.Context, tagIDs []int32) (map[int32]string, error) {
	var tags []model.Tag
	err := r.db.WithContext(ctx).
		Model(&model.Tag{}).
		Where(`"tagId" IN ?`, tagIDs).
		Find(&tags).Error
	if err != nil {
		log.Printf("Ошибка при получении названий тегов: %v", err)
		return nil, err
	}

	tagTitles := make(map[int32]string)
	for _, tag := range tags {
		tagTitles[tag.TagID] = tag.Title
	}

	return tagTitles, nil
}
