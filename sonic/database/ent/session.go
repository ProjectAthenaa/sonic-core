// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/session"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/user"
	"github.com/google/uuid"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DeviceName holds the value of the "DeviceName" field.
	DeviceName string `json:"DeviceName,omitempty"`
	// OS holds the value of the "OS" field.
	OS string `json:"OS,omitempty"`
	// DeviceType holds the value of the "DeviceType" field.
	DeviceType session.DeviceType `json:"DeviceType,omitempty"`
	// IP holds the value of the "IP" field.
	IP string `json:"IP,omitempty"`
	// Expired holds the value of the "Expired" field.
	Expired bool `json:"Expired,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges         SessionEdges `json:"edges"`
	user_sessions *uuid.UUID
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldExpired:
			values[i] = new(sql.NullBool)
		case session.FieldDeviceName, session.FieldOS, session.FieldDeviceType, session.FieldIP:
			values[i] = new(sql.NullString)
		case session.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case session.FieldID:
			values[i] = new(uuid.UUID)
		case session.ForeignKeys[0]: // user_sessions
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Session", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case session.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case session.FieldDeviceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeviceName", values[i])
			} else if value.Valid {
				s.DeviceName = value.String
			}
		case session.FieldOS:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field OS", values[i])
			} else if value.Valid {
				s.OS = value.String
			}
		case session.FieldDeviceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeviceType", values[i])
			} else if value.Valid {
				s.DeviceType = session.DeviceType(value.String)
			}
		case session.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field IP", values[i])
			} else if value.Valid {
				s.IP = value.String
			}
		case session.FieldExpired:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field Expired", values[i])
			} else if value.Valid {
				s.Expired = value.Bool
			}
		case session.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_sessions", values[i])
			} else if value.Valid {
				s.user_sessions = new(uuid.UUID)
				*s.user_sessions = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Session entity.
func (s *Session) QueryUser() *UserQuery {
	return (&SessionClient{config: s.config}).QueryUser(s)
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", DeviceName=")
	builder.WriteString(s.DeviceName)
	builder.WriteString(", OS=")
	builder.WriteString(s.OS)
	builder.WriteString(", DeviceType=")
	builder.WriteString(fmt.Sprintf("%v", s.DeviceType))
	builder.WriteString(", IP=")
	builder.WriteString(s.IP)
	builder.WriteString(", Expired=")
	builder.WriteString(fmt.Sprintf("%v", s.Expired))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
