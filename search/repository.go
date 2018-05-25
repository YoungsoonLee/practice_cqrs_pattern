package search

import (
	"context"

	"github.com/YoungsoonLee/exam_cqrs_pattern/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	SearchMeow(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return impl.InsertMeow(ctx, meow)
}

func SearchMeow(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error) {
	return impl.SearchMeow(ctx, query, skip, take)
}
