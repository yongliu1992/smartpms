package data

import "github.com/yongliu1992/smartpms/app/property/service/internal/biz"
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

func (sp *shopRepo) ListShop(ctx context.Context, pageNum, pageSize, villageId int) (*biz.Shop, error) {
	//TODO implement me
	panic("implement me")
}

func NewShopRepo(data *Data) *shopRepo {
	return &shopRepo{
		data: data,
	}
}
