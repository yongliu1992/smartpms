package data

import (
	"context"
	"github.com/yongliu1992/smartpms/app/property/service/internal/biz"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent/community"
	"time"
)

var _ biz.CommunityRepo = (*communityRepo)(nil)

//Community 小区
type Community struct {
	ID int `db:"id"`
	// Comment: 小区名字
	Name *string `db:"name"`
	// Comment: 省份id
	ProvinceID *int `db:"province_id"`
	// Comment: 市/州id
	CityID *int `db:"city_id"`
	// Comment: 区/县id
	AreaID *int `db:"area_id"`
	// Comment: 面积
	AreaNum *int `db:"area_num"`
	// Comment: 管理员id
	AdminID *int `db:"admin_id"`
	// Comment: 小区编码
	Number *int `db:"number"`
	// Comment: 状态，1入驻审核2入驻成功3取消入驻
	State *int `db:"state"`
	// Comment: 开始时间
	StartTime *time.Time `db:"start_time"`
	// Comment: 结束时间
	EndTime   *time.Time `db:"end_time"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type communityRepo struct {
	data *Data
}

func (c communityRepo) List(ctx context.Context, co biz.ListCommunityReq) (*[]biz.Community, error) {
	if co.Page < 1 {
		co.Page = 1
	}
	result, err := c.data.db.Community.Query().Where(community.AdminID(co.AdminID)).Offset((co.Page - 1) * co.PageSize).Limit(co.PageSize).All(ctx)
	if err != nil {
		return nil, err
	}
	var res = make([]biz.Community, len(result))
	for k, v := range result {
		b := biz.Community{ID: v.ID, ProvinceID: v.ProvinceID, CityID: v.CityID, AreaID: v.AreaID, State: v.State, Number: v.CommNumber, AreaNum: v.AreaNum}
		res[k] = b
	}
	return &res, err

}

func (c communityRepo) Add(ctx context.Context, co biz.Community) (biz.Community, error) {
	result, err := c.data.db.Community.
		Create().
		SetName(co.Name).
		SetAdminID(co.AdminID).
		SetCommNumber(co.Number).
		SetAreaID(co.AreaID).
		SetAreaNum(co.AreaNum).
		SetCityID(co.CityID).
		SetEndTime(co.EndTime).
		SetStartTime(co.StartTime).
		SetProvinceID(co.ProvinceID).
		SetState(co.State).
		Save(ctx)
	return biz.Community{ID: result.ID}, err
}

func (c communityRepo) Update(ctx context.Context, co *biz.Community) (bool, error) {
	affectRows, err := c.data.db.Community.
		Update().
		Where(community.ID(co.ID)).
		SetName(co.Name).
		SetAdminID(co.AdminID).
		SetCommNumber(co.Number).
		SetAreaID(co.AreaID).
		SetAreaNum(co.AreaNum).
		SetCityID(co.CityID).
		SetEndTime(co.EndTime).
		SetStartTime(co.StartTime).
		SetProvinceID(co.ProvinceID).
		SetState(co.State).
		Save(ctx)
	return affectRows > 0, err
}

func NewCommunityRepo(data *Data) biz.CommunityRepo {
	return &communityRepo{data: data}
}
