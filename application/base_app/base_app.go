package baseapp

import (
	"context"

	"bitbucket.org/be-proj/osp-base/domain/repository/base"
)

type baseApp struct {
	ba base.BaseRepository
}

var _ BaseAppInterface = &baseApp{}

type BaseAppInterface interface {
	BaseQuery(ctx context.Context, key string) (value string, err error)
}

func (b *baseApp) BaseQuery(ctx context.Context, key string) (value string, err error) {
	return b.ba.BaseQuery(ctx, key)
}
