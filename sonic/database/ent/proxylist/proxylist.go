// Code generated by entc, DO NOT EDIT.

package proxylist

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the proxylist type in the database.
	Label = "proxy_list"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "App"
	// EdgeProxies holds the string denoting the proxies edge name in mutations.
	EdgeProxies = "Proxies"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "Task"
	// Table holds the table name of the proxylist in the database.
	Table = "proxy_lists"
	// AppTable is the table that holds the App relation/edge. The primary key declared below.
	AppTable = "app_ProxyLists"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "apps"
	// ProxiesTable is the table that holds the Proxies relation/edge.
	ProxiesTable = "proxies"
	// ProxiesInverseTable is the table name for the Proxy entity.
	// It exists in this package in order to avoid circular dependency with the "proxy" package.
	ProxiesInverseTable = "proxies"
	// ProxiesColumn is the table column denoting the Proxies relation/edge.
	ProxiesColumn = "proxy_list_proxies"
	// TaskTable is the table that holds the Task relation/edge. The primary key declared below.
	TaskTable = "task_ProxyList"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
)

// Columns holds all SQL columns for proxylist fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldType,
}

var (
	// AppPrimaryKey and AppColumn2 are the table columns denoting the
	// primary key for the App relation (M2M).
	AppPrimaryKey = []string{"app_id", "proxy_list_id"}
	// TaskPrimaryKey and TaskColumn2 are the table columns denoting the
	// primary key for the Task relation (M2M).
	TaskPrimaryKey = []string{"task_id", "proxy_list_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultName holds the default value on creation for the "Name" field.
	DefaultName string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "Type" enum field.
type Type string

// Type values.
const (
	TypeResidential Type = "Residential"
	TypeDatacenter  Type = "Datacenter"
	TypeISP         Type = "ISP"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "Type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeResidential, TypeDatacenter, TypeISP:
		return nil
	default:
		return fmt.Errorf("proxylist: invalid enum value for Type field: %q", _type)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (_type Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(_type.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (_type *Type) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*_type = Type(str)
	if err := TypeValidator(*_type); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}
