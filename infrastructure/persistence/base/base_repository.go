package base

import (
	"context"

	baseEntity "github.com/nansuri/godeep-base/domain/entity/base"

	"github.com/nansuri/godeep-base/domain/repository/base"
	"gorm.io/gorm"
)

type BaseRepo struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepo {
	return &BaseRepo{db}
}

var _ base.BaseRepository = &BaseRepo{}

func (b *BaseRepo) BaseQuery(ctx context.Context, key string) (value string, err error) {

	var tbc baseEntity.TmBaseContext

	r := b.db.Model(&tbc).Where("enum_type = ?", key).Order("id ASC").Pluck("value", &value)
	if r.Error != nil {
		return "", r.Error
	}

	return value, nil
}
