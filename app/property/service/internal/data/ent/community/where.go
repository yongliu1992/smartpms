// Code generated by entc, DO NOT EDIT.

package community

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ProvinceID applies equality check predicate on the "province_id" field. It's identical to ProvinceIDEQ.
func ProvinceID(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvinceID), v))
	})
}

// CityID applies equality check predicate on the "city_id" field. It's identical to CityIDEQ.
func CityID(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCityID), v))
	})
}

// AreaID applies equality check predicate on the "area_id" field. It's identical to AreaIDEQ.
func AreaID(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAreaID), v))
	})
}

// AreaNum applies equality check predicate on the "area_num" field. It's identical to AreaNumEQ.
func AreaNum(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAreaNum), v))
	})
}

// AdminID applies equality check predicate on the "admin_id" field. It's identical to AdminIDEQ.
func AdminID(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdminID), v))
	})
}

// CommNumber applies equality check predicate on the "comm_number" field. It's identical to CommNumberEQ.
func CommNumber(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommNumber), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ProvinceIDEQ applies the EQ predicate on the "province_id" field.
func ProvinceIDEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvinceID), v))
	})
}

// ProvinceIDNEQ applies the NEQ predicate on the "province_id" field.
func ProvinceIDNEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProvinceID), v))
	})
}

// ProvinceIDIn applies the In predicate on the "province_id" field.
func ProvinceIDIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProvinceID), v...))
	})
}

// ProvinceIDNotIn applies the NotIn predicate on the "province_id" field.
func ProvinceIDNotIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProvinceID), v...))
	})
}

// ProvinceIDGT applies the GT predicate on the "province_id" field.
func ProvinceIDGT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProvinceID), v))
	})
}

// ProvinceIDGTE applies the GTE predicate on the "province_id" field.
func ProvinceIDGTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProvinceID), v))
	})
}

// ProvinceIDLT applies the LT predicate on the "province_id" field.
func ProvinceIDLT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProvinceID), v))
	})
}

// ProvinceIDLTE applies the LTE predicate on the "province_id" field.
func ProvinceIDLTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProvinceID), v))
	})
}

// CityIDEQ applies the EQ predicate on the "city_id" field.
func CityIDEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCityID), v))
	})
}

// CityIDNEQ applies the NEQ predicate on the "city_id" field.
func CityIDNEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCityID), v))
	})
}

// CityIDIn applies the In predicate on the "city_id" field.
func CityIDIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCityID), v...))
	})
}

// CityIDNotIn applies the NotIn predicate on the "city_id" field.
func CityIDNotIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCityID), v...))
	})
}

// CityIDGT applies the GT predicate on the "city_id" field.
func CityIDGT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCityID), v))
	})
}

// CityIDGTE applies the GTE predicate on the "city_id" field.
func CityIDGTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCityID), v))
	})
}

// CityIDLT applies the LT predicate on the "city_id" field.
func CityIDLT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCityID), v))
	})
}

// CityIDLTE applies the LTE predicate on the "city_id" field.
func CityIDLTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCityID), v))
	})
}

// AreaIDEQ applies the EQ predicate on the "area_id" field.
func AreaIDEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAreaID), v))
	})
}

// AreaIDNEQ applies the NEQ predicate on the "area_id" field.
func AreaIDNEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAreaID), v))
	})
}

// AreaIDIn applies the In predicate on the "area_id" field.
func AreaIDIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAreaID), v...))
	})
}

// AreaIDNotIn applies the NotIn predicate on the "area_id" field.
func AreaIDNotIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAreaID), v...))
	})
}

// AreaIDGT applies the GT predicate on the "area_id" field.
func AreaIDGT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAreaID), v))
	})
}

// AreaIDGTE applies the GTE predicate on the "area_id" field.
func AreaIDGTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAreaID), v))
	})
}

// AreaIDLT applies the LT predicate on the "area_id" field.
func AreaIDLT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAreaID), v))
	})
}

// AreaIDLTE applies the LTE predicate on the "area_id" field.
func AreaIDLTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAreaID), v))
	})
}

// AreaNumEQ applies the EQ predicate on the "area_num" field.
func AreaNumEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAreaNum), v))
	})
}

// AreaNumNEQ applies the NEQ predicate on the "area_num" field.
func AreaNumNEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAreaNum), v))
	})
}

// AreaNumIn applies the In predicate on the "area_num" field.
func AreaNumIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAreaNum), v...))
	})
}

// AreaNumNotIn applies the NotIn predicate on the "area_num" field.
func AreaNumNotIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAreaNum), v...))
	})
}

// AreaNumGT applies the GT predicate on the "area_num" field.
func AreaNumGT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAreaNum), v))
	})
}

// AreaNumGTE applies the GTE predicate on the "area_num" field.
func AreaNumGTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAreaNum), v))
	})
}

// AreaNumLT applies the LT predicate on the "area_num" field.
func AreaNumLT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAreaNum), v))
	})
}

// AreaNumLTE applies the LTE predicate on the "area_num" field.
func AreaNumLTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAreaNum), v))
	})
}

// AdminIDEQ applies the EQ predicate on the "admin_id" field.
func AdminIDEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdminID), v))
	})
}

// AdminIDNEQ applies the NEQ predicate on the "admin_id" field.
func AdminIDNEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdminID), v))
	})
}

// AdminIDIn applies the In predicate on the "admin_id" field.
func AdminIDIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAdminID), v...))
	})
}

// AdminIDNotIn applies the NotIn predicate on the "admin_id" field.
func AdminIDNotIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAdminID), v...))
	})
}

// AdminIDGT applies the GT predicate on the "admin_id" field.
func AdminIDGT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAdminID), v))
	})
}

// AdminIDGTE applies the GTE predicate on the "admin_id" field.
func AdminIDGTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAdminID), v))
	})
}

// AdminIDLT applies the LT predicate on the "admin_id" field.
func AdminIDLT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAdminID), v))
	})
}

// AdminIDLTE applies the LTE predicate on the "admin_id" field.
func AdminIDLTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAdminID), v))
	})
}

// CommNumberEQ applies the EQ predicate on the "comm_number" field.
func CommNumberEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommNumber), v))
	})
}

// CommNumberNEQ applies the NEQ predicate on the "comm_number" field.
func CommNumberNEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCommNumber), v))
	})
}

// CommNumberIn applies the In predicate on the "comm_number" field.
func CommNumberIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCommNumber), v...))
	})
}

// CommNumberNotIn applies the NotIn predicate on the "comm_number" field.
func CommNumberNotIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCommNumber), v...))
	})
}

// CommNumberGT applies the GT predicate on the "comm_number" field.
func CommNumberGT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCommNumber), v))
	})
}

// CommNumberGTE applies the GTE predicate on the "comm_number" field.
func CommNumberGTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCommNumber), v))
	})
}

// CommNumberLT applies the LT predicate on the "comm_number" field.
func CommNumberLT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCommNumber), v))
	})
}

// CommNumberLTE applies the LTE predicate on the "comm_number" field.
func CommNumberLTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCommNumber), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...int) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v int) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartTime), v))
	})
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartTime), v...))
	})
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartTime), v...))
	})
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartTime), v))
	})
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartTime), v))
	})
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartTime), v))
	})
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartTime), v))
	})
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndTime), v))
	})
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndTime), v...))
	})
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndTime), v...))
	})
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndTime), v))
	})
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndTime), v))
	})
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndTime), v))
	})
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndTime), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Community {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Community(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Community) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Community) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Community) predicate.Community {
	return predicate.Community(func(s *sql.Selector) {
		p(s.Not())
	})
}
