package data

import (
	"github.com/yongliu1992/smartpms/app/property/service/internal/biz"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent/shop"
)
import "context"

var _ biz.ShopsRepo = (*shopRepo)(nil)

type shopRepo struct {
	data *Data
}

func (sp *shopRepo) CreateShop(ctx context.Context, s *biz.Shop) (*biz.Shop, error) {
	data, err := sp.data.db.Shop.
		Create().
		SetCommunityID(s.CommunityID).
		SetLayer(s.Layer).
		SetRemark(s.Remark).
		SetRent(s.Rent).
		SetRoomNum(s.RoomNum).
		SetBuiltUpArea(s.BuiltUpArea).
		SetRoomArea(s.RoomArea).
		SetFeeRate(s.FeeRate).
		SetFloorID(s.FloorID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Shop{ID: data.ID, Remark: data.Remark}, err
}

func (sp *shopRepo) UpdateShop(ctx context.Context, s *biz.Shop) (*biz.Shop, error) {
	//TODO implement me
	panic("implement me")
}

func (sp *shopRepo) ListShop(ctx context.Context, pageNum, pageSize, villageId int) (resp []*biz.Shop, err error) {
	if pageNum <= 0 {
		pageNum = 1
	}
	data, err := sp.data.db.Debug().Shop.Query().Where(shop.And(shop.CommunityID(villageId))).Offset((pageNum - 1) * pageSize).Limit(pageNum).All(ctx)
	for _, v := range data {
		bp := biz.Shop{ID: v.ID, FloorID: v.FloorID,CreatedAt: v.CreatedAt,BuiltUpArea: v.BuiltUpArea}
		resp = append(resp, &bp)
	}
	return resp, err

}

func NewShopRepo(data *Data) *shopRepo {
	return &shopRepo{
		data: data,
	}
}
