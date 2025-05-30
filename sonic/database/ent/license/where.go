// Code generated by entc, DO NOT EDIT.

package license

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.License {
	return predicate.License(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.License {
	return predicate.License(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Key applies equality check predicate on the "Key" field. It's identical to KeyEQ.
func Key(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// HardwareID applies equality check predicate on the "HardwareID" field. It's identical to HardwareIDEQ.
func HardwareID(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHardwareID), v))
	})
}

// MobileHardwareID applies equality check predicate on the "MobileHardwareID" field. It's identical to MobileHardwareIDEQ.
func MobileHardwareID(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMobileHardwareID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// KeyEQ applies the EQ predicate on the "Key" field.
func KeyEQ(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// KeyNEQ applies the NEQ predicate on the "Key" field.
func KeyNEQ(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKey), v))
	})
}

// KeyIn applies the In predicate on the "Key" field.
func KeyIn(vs ...string) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKey), v...))
	})
}

// KeyNotIn applies the NotIn predicate on the "Key" field.
func KeyNotIn(vs ...string) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKey), v...))
	})
}

// KeyGT applies the GT predicate on the "Key" field.
func KeyGT(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKey), v))
	})
}

// KeyGTE applies the GTE predicate on the "Key" field.
func KeyGTE(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKey), v))
	})
}

// KeyLT applies the LT predicate on the "Key" field.
func KeyLT(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKey), v))
	})
}

// KeyLTE applies the LTE predicate on the "Key" field.
func KeyLTE(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKey), v))
	})
}

// KeyContains applies the Contains predicate on the "Key" field.
func KeyContains(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKey), v))
	})
}

// KeyHasPrefix applies the HasPrefix predicate on the "Key" field.
func KeyHasPrefix(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKey), v))
	})
}

// KeyHasSuffix applies the HasSuffix predicate on the "Key" field.
func KeyHasSuffix(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKey), v))
	})
}

// KeyEqualFold applies the EqualFold predicate on the "Key" field.
func KeyEqualFold(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKey), v))
	})
}

// KeyContainsFold applies the ContainsFold predicate on the "Key" field.
func KeyContainsFold(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKey), v))
	})
}

// HardwareIDEQ applies the EQ predicate on the "HardwareID" field.
func HardwareIDEQ(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHardwareID), v))
	})
}

// HardwareIDNEQ applies the NEQ predicate on the "HardwareID" field.
func HardwareIDNEQ(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHardwareID), v))
	})
}

// HardwareIDIn applies the In predicate on the "HardwareID" field.
func HardwareIDIn(vs ...string) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHardwareID), v...))
	})
}

// HardwareIDNotIn applies the NotIn predicate on the "HardwareID" field.
func HardwareIDNotIn(vs ...string) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHardwareID), v...))
	})
}

// HardwareIDGT applies the GT predicate on the "HardwareID" field.
func HardwareIDGT(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHardwareID), v))
	})
}

// HardwareIDGTE applies the GTE predicate on the "HardwareID" field.
func HardwareIDGTE(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHardwareID), v))
	})
}

// HardwareIDLT applies the LT predicate on the "HardwareID" field.
func HardwareIDLT(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHardwareID), v))
	})
}

// HardwareIDLTE applies the LTE predicate on the "HardwareID" field.
func HardwareIDLTE(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHardwareID), v))
	})
}

// HardwareIDContains applies the Contains predicate on the "HardwareID" field.
func HardwareIDContains(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHardwareID), v))
	})
}

// HardwareIDHasPrefix applies the HasPrefix predicate on the "HardwareID" field.
func HardwareIDHasPrefix(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHardwareID), v))
	})
}

// HardwareIDHasSuffix applies the HasSuffix predicate on the "HardwareID" field.
func HardwareIDHasSuffix(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHardwareID), v))
	})
}

// HardwareIDIsNil applies the IsNil predicate on the "HardwareID" field.
func HardwareIDIsNil() predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHardwareID)))
	})
}

// HardwareIDNotNil applies the NotNil predicate on the "HardwareID" field.
func HardwareIDNotNil() predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHardwareID)))
	})
}

// HardwareIDEqualFold applies the EqualFold predicate on the "HardwareID" field.
func HardwareIDEqualFold(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHardwareID), v))
	})
}

// HardwareIDContainsFold applies the ContainsFold predicate on the "HardwareID" field.
func HardwareIDContainsFold(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHardwareID), v))
	})
}

// MobileHardwareIDEQ applies the EQ predicate on the "MobileHardwareID" field.
func MobileHardwareIDEQ(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMobileHardwareID), v))
	})
}

// MobileHardwareIDNEQ applies the NEQ predicate on the "MobileHardwareID" field.
func MobileHardwareIDNEQ(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMobileHardwareID), v))
	})
}

// MobileHardwareIDIn applies the In predicate on the "MobileHardwareID" field.
func MobileHardwareIDIn(vs ...string) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMobileHardwareID), v...))
	})
}

// MobileHardwareIDNotIn applies the NotIn predicate on the "MobileHardwareID" field.
func MobileHardwareIDNotIn(vs ...string) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMobileHardwareID), v...))
	})
}

// MobileHardwareIDGT applies the GT predicate on the "MobileHardwareID" field.
func MobileHardwareIDGT(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMobileHardwareID), v))
	})
}

// MobileHardwareIDGTE applies the GTE predicate on the "MobileHardwareID" field.
func MobileHardwareIDGTE(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMobileHardwareID), v))
	})
}

// MobileHardwareIDLT applies the LT predicate on the "MobileHardwareID" field.
func MobileHardwareIDLT(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMobileHardwareID), v))
	})
}

// MobileHardwareIDLTE applies the LTE predicate on the "MobileHardwareID" field.
func MobileHardwareIDLTE(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMobileHardwareID), v))
	})
}

// MobileHardwareIDContains applies the Contains predicate on the "MobileHardwareID" field.
func MobileHardwareIDContains(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMobileHardwareID), v))
	})
}

// MobileHardwareIDHasPrefix applies the HasPrefix predicate on the "MobileHardwareID" field.
func MobileHardwareIDHasPrefix(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMobileHardwareID), v))
	})
}

// MobileHardwareIDHasSuffix applies the HasSuffix predicate on the "MobileHardwareID" field.
func MobileHardwareIDHasSuffix(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMobileHardwareID), v))
	})
}

// MobileHardwareIDIsNil applies the IsNil predicate on the "MobileHardwareID" field.
func MobileHardwareIDIsNil() predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMobileHardwareID)))
	})
}

// MobileHardwareIDNotNil applies the NotNil predicate on the "MobileHardwareID" field.
func MobileHardwareIDNotNil() predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMobileHardwareID)))
	})
}

// MobileHardwareIDEqualFold applies the EqualFold predicate on the "MobileHardwareID" field.
func MobileHardwareIDEqualFold(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMobileHardwareID), v))
	})
}

// MobileHardwareIDContainsFold applies the ContainsFold predicate on the "MobileHardwareID" field.
func MobileHardwareIDContainsFold(v string) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMobileHardwareID), v))
	})
}

// TypeEQ applies the EQ predicate on the "Type" field.
func TypeEQ(v Type) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "Type" field.
func TypeNEQ(v Type) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "Type" field.
func TypeIn(vs ...Type) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "Type" field.
func TypeNotIn(vs ...Type) predicate.License {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.License(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// HasUser applies the HasEdge predicate on the "User" edge.
func HasUser() predicate.License {
	return predicate.License(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "User" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStripe applies the HasEdge predicate on the "Stripe" edge.
func HasStripe() predicate.License {
	return predicate.License(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StripeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StripeTable, StripeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStripeWith applies the HasEdge predicate on the "Stripe" edge with a given conditions (other predicates).
func HasStripeWith(preds ...predicate.Stripe) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StripeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StripeTable, StripeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.License) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.License) predicate.License {
	return predicate.License(func(s *sql.Selector) {
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
func Not(p predicate.License) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		p(s.Not())
	})
}
