package db

import (
	"context"

	"github.com/YoungsoonLee/exam_cqrs_pattern/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ListMeows(ctx context.Context, skip uint16, take uint16) ([]schema.Meow, error)
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

func ListMeows(ctx context.Context, skip uint16, take uint16) ([]schema.Meow, error) {
	return impl.ListMeows(ctx, skip, take)
}
