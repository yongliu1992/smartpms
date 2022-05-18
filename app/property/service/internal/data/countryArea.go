package data

// CountryArea 地名
type CountryArea struct {
	CodeID   int     `db:"code_id"`
	ParentID *int    `db:"parent_id"`
	CityName *string `db:"city_name"`
}
