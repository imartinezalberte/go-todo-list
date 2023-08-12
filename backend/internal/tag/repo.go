package tag

import (
	"context"

	"gorm.io/gorm"
)

type (
	Repo interface {
		GetTags(context.Context, TagsQuery) (Tags, error)
		GetTag(context.Context, TagQuery) (Tag, error)
		AddTag(context.Context, Tag) (Tag, error)
		UpdateTag(context.Context, Tag) (Tag, error)
		DelTag(context.Context, Tag) error
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) Repo {
	return &repo{db}
}

func (r *repo) GetTags(ctx context.Context, query TagsQuery) (tags Tags, err error) {
	result := r.db.WithContext(ctx).Scopes(query.Scopes()).Find(&tags)
	return tags, result.Error
}

func (r *repo) GetTag(ctx context.Context, tag TagQuery) (c Tag, err error) {
	result := r.db.WithContext(ctx).Find(&c, "id = ?", tag.ID)
	return c, result.Error
}

func (r *repo) AddTag(ctx context.Context, tag Tag) (Tag, error) {
	result := r.db.WithContext(ctx).Create(&tag)
	return tag, result.Error
}

func (r *repo) UpdateTag(ctx context.Context, tag Tag) (Tag, error) {
	result := r.db.WithContext(ctx).Updates(&tag)
	return tag, result.Error
}

func (r *repo) DelTag(ctx context.Context, tag Tag) error {
	result := r.db.WithContext(ctx).Delete(tag)
	return result.Error
}
