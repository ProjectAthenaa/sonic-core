// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/profilegroup"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/task"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/taskgroup"
	"github.com/google/uuid"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// StartTime holds the value of the "StartTime" field.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges              TaskEdges `json:"edges"`
	task_profile_group *uuid.UUID
	task_group_tasks   *uuid.UUID
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Product holds the value of the Product edge.
	Product []*Product `json:"Product,omitempty"`
	// ProxyList holds the value of the ProxyList edge.
	ProxyList []*ProxyList `json:"ProxyList,omitempty"`
	// ProfileGroup holds the value of the ProfileGroup edge.
	ProfileGroup *ProfileGroup `json:"ProfileGroup,omitempty"`
	// TaskGroup holds the value of the TaskGroup edge.
	TaskGroup *TaskGroup `json:"TaskGroup,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "Product"}
}

// ProxyListOrErr returns the ProxyList value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) ProxyListOrErr() ([]*ProxyList, error) {
	if e.loadedTypes[1] {
		return e.ProxyList, nil
	}
	return nil, &NotLoadedError{edge: "ProxyList"}
}

// ProfileGroupOrErr returns the ProfileGroup value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) ProfileGroupOrErr() (*ProfileGroup, error) {
	if e.loadedTypes[2] {
		if e.ProfileGroup == nil {
			// The edge ProfileGroup was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profilegroup.Label}
		}
		return e.ProfileGroup, nil
	}
	return nil, &NotLoadedError{edge: "ProfileGroup"}
}

// TaskGroupOrErr returns the TaskGroup value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) TaskGroupOrErr() (*TaskGroup, error) {
	if e.loadedTypes[3] {
		if e.TaskGroup == nil {
			// The edge TaskGroup was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: taskgroup.Label}
		}
		return e.TaskGroup, nil
	}
	return nil, &NotLoadedError{edge: "TaskGroup"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldCreatedAt, task.FieldUpdatedAt, task.FieldStartTime:
			values[i] = new(sql.NullTime)
		case task.FieldID:
			values[i] = new(uuid.UUID)
		case task.ForeignKeys[0]: // task_profile_group
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case task.ForeignKeys[1]: // task_group_tasks
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Task", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case task.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case task.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case task.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field StartTime", values[i])
			} else if value.Valid {
				t.StartTime = new(time.Time)
				*t.StartTime = value.Time
			}
		case task.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field task_profile_group", values[i])
			} else if value.Valid {
				t.task_profile_group = new(uuid.UUID)
				*t.task_profile_group = *value.S.(*uuid.UUID)
			}
		case task.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field task_group_tasks", values[i])
			} else if value.Valid {
				t.task_group_tasks = new(uuid.UUID)
				*t.task_group_tasks = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryProduct queries the "Product" edge of the Task entity.
func (t *Task) QueryProduct() *ProductQuery {
	return (&TaskClient{config: t.config}).QueryProduct(t)
}

// QueryProxyList queries the "ProxyList" edge of the Task entity.
func (t *Task) QueryProxyList() *ProxyListQuery {
	return (&TaskClient{config: t.config}).QueryProxyList(t)
}

// QueryProfileGroup queries the "ProfileGroup" edge of the Task entity.
func (t *Task) QueryProfileGroup() *ProfileGroupQuery {
	return (&TaskClient{config: t.config}).QueryProfileGroup(t)
}

// QueryTaskGroup queries the "TaskGroup" edge of the Task entity.
func (t *Task) QueryTaskGroup() *TaskGroupQuery {
	return (&TaskClient{config: t.config}).QueryTaskGroup(t)
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return (&TaskClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	if v := t.StartTime; v != nil {
		builder.WriteString(", StartTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

func (t Tasks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
