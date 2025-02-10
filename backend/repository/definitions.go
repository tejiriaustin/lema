package repository

import (
	"context"

	"github.com/tejiriaustin/lema/models"
)

type (
	Creator[T models.Models] interface {
		Create(ctx context.Context, data T) (*T, error)
	}
	Finder[T models.Models] interface {
		FindOne(ctx context.Context, queryFilter *Query, preloads ...string) (*T, error)
		FindManyPaginated(ctx context.Context, queryFilter *Query, page, perPage int64, preloads ...string) ([]*T, *Paginator, error)
		Select(ctx context.Context, target interface{}, query string, args ...interface{}) error
	}
	Deleter[T models.Models] interface {
		DeleteMany(ctx context.Context, queryFilter *Query) error
	}
	Updater[T models.Models] interface {
		Update(ctx context.Context, dataObject T) (*T, error)
	}
	Counter[T models.Models] interface {
		Count(ctx context.Context, queryFilter *Query) (int64, error)
	}

	RepoInterface[T models.Models] interface {
		Creator[T]
		Finder[T]
		Deleter[T]
		Updater[T]
		Counter[T]
	}
)
