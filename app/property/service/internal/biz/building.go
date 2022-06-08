package biz

import (
	"context"
	"time"
)

type UnitRepo interface {
	Add(ctx context.Context, co Community) (Community, error)
	List(ctx context.Context, co ListCommunityReq) (*[]Community, error)
	Update(ctx context.Context, co *Community) (bool, error)
}

type UnitS struct {
	ID int `db:"id"`
	// Comment: 小区名字
	Name string `db:"name"`
	// CommId 单元id
	CommId    int       `db:"comm_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
