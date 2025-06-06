// Code generated by entc, DO NOT EDIT.

package device

import (
	"entgo.io/ent/dialect/sql"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GpuVendor applies equality check predicate on the "gpuVendor" field. It's identical to GpuVendorEQ.
func GpuVendor(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGpuVendor), v))
	})
}

// Adevice applies equality check predicate on the "adevice" field. It's identical to AdeviceEQ.
func Adevice(v sonic.Map) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdevice), v))
	})
}

// GpuVendorEQ applies the EQ predicate on the "gpuVendor" field.
func GpuVendorEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGpuVendor), v))
	})
}

// GpuVendorNEQ applies the NEQ predicate on the "gpuVendor" field.
func GpuVendorNEQ(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGpuVendor), v))
	})
}

// GpuVendorIn applies the In predicate on the "gpuVendor" field.
func GpuVendorIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGpuVendor), v...))
	})
}

// GpuVendorNotIn applies the NotIn predicate on the "gpuVendor" field.
func GpuVendorNotIn(vs ...string) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGpuVendor), v...))
	})
}

// GpuVendorGT applies the GT predicate on the "gpuVendor" field.
func GpuVendorGT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGpuVendor), v))
	})
}

// GpuVendorGTE applies the GTE predicate on the "gpuVendor" field.
func GpuVendorGTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGpuVendor), v))
	})
}

// GpuVendorLT applies the LT predicate on the "gpuVendor" field.
func GpuVendorLT(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGpuVendor), v))
	})
}

// GpuVendorLTE applies the LTE predicate on the "gpuVendor" field.
func GpuVendorLTE(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGpuVendor), v))
	})
}

// GpuVendorContains applies the Contains predicate on the "gpuVendor" field.
func GpuVendorContains(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGpuVendor), v))
	})
}

// GpuVendorHasPrefix applies the HasPrefix predicate on the "gpuVendor" field.
func GpuVendorHasPrefix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGpuVendor), v))
	})
}

// GpuVendorHasSuffix applies the HasSuffix predicate on the "gpuVendor" field.
func GpuVendorHasSuffix(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGpuVendor), v))
	})
}

// GpuVendorEqualFold applies the EqualFold predicate on the "gpuVendor" field.
func GpuVendorEqualFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGpuVendor), v))
	})
}

// GpuVendorContainsFold applies the ContainsFold predicate on the "gpuVendor" field.
func GpuVendorContainsFold(v string) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGpuVendor), v))
	})
}

// AdeviceEQ applies the EQ predicate on the "adevice" field.
func AdeviceEQ(v sonic.Map) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdevice), v))
	})
}

// AdeviceNEQ applies the NEQ predicate on the "adevice" field.
func AdeviceNEQ(v sonic.Map) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdevice), v))
	})
}

// AdeviceIn applies the In predicate on the "adevice" field.
func AdeviceIn(vs ...sonic.Map) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAdevice), v...))
	})
}

// AdeviceNotIn applies the NotIn predicate on the "adevice" field.
func AdeviceNotIn(vs ...sonic.Map) predicate.Device {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Device(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAdevice), v...))
	})
}

// AdeviceGT applies the GT predicate on the "adevice" field.
func AdeviceGT(v sonic.Map) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAdevice), v))
	})
}

// AdeviceGTE applies the GTE predicate on the "adevice" field.
func AdeviceGTE(v sonic.Map) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAdevice), v))
	})
}

// AdeviceLT applies the LT predicate on the "adevice" field.
func AdeviceLT(v sonic.Map) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAdevice), v))
	})
}

// AdeviceLTE applies the LTE predicate on the "adevice" field.
func AdeviceLTE(v sonic.Map) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAdevice), v))
	})
}

// AdeviceIsNil applies the IsNil predicate on the "adevice" field.
func AdeviceIsNil() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAdevice)))
	})
}

// AdeviceNotNil applies the NotNil predicate on the "adevice" field.
func AdeviceNotNil() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAdevice)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
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
func Not(p predicate.Device) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		p(s.Not())
	})
}
