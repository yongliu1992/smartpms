package data

import "time"

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
}
