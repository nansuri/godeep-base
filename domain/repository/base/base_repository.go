package base

import "context"

type BaseRepository interface {
	BaseQuery(ctx context.Context, key string) (value string, err error)
}
