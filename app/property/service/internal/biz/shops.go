package biz

import (
	"context"
	"time"
)

type ListShopsReq struct {
	VillageId int
	State     int
	Id        int
}

type ListShopsResp struct {
	Data []Shop
}

// Shop is the model entity for the Shop schema.
type Shop struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FloorID holds the value of the "floor_id" field.
	FloorID string `json:"floor_id,omitempty"`
	// RoomNum holds the value of the "room_num" field.
	RoomNum int `json:"room_num,omitempty"`
	// Layer holds the value of the "layer" field.
	Layer int `json:"layer,omitempty"`
	// BuiltUpArea holds the value of the "built_up_area" field.
	BuiltUpArea float32 `json:"built_up_area,omitempty"`
	// CommunityID holds the value of the "community_id" field.
	CommunityID int `json:"community_id,omitempty"`
	// FeeRate holds the value of the "fee_rate" field.
	FeeRate float32 `json:"fee_rate,omitempty"`
	// RoomArea holds the value of the "room_area" field.
	RoomArea float32 `json:"room_area,omitempty"`
	// Rent holds the value of the "rent" field.
	Rent float32 `json:"rent,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type ShopsRepo interface {
	CreateShop(ctx context.Context, s *Shop) (*Shop, error)
	UpdateShop(ctx context.Context, s *Shop) (*Shop, error)
	ListShop(ctx context.Context, pageNum, pageSize, villageId int) (*Shop, error)
}

type ShopUseCase struct {
	repo ShopsRepo
}

func (uc *ShopUseCase) ListShop(ctx context.Context, pageNum, pageSize, villageId int) (*Shop, error) {
	return uc.repo.ListShop(ctx, pageNum, pageSize, villageId)
}

func (uc *ShopUseCase) UpdateShop(ctx context.Context, s *Shop) (*Shop, error) {
	return uc.repo.UpdateShop(ctx, s)
}

func (uc *ShopUseCase) CreateShop(ctx context.Context, s *Shop) (*Shop, error) {
	return uc.repo.CreateShop(ctx, s)
}

func NewShopUseCase(repo ShopsRepo) *ShopUseCase {
	return &ShopUseCase{repo: repo}
}
