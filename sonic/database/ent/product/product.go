// Code generated by entc, DO NOT EDIT.

package product

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldLookupType holds the string denoting the lookuptype field in the database.
	FieldLookupType = "lookup_type"
	// FieldPositiveKeywords holds the string denoting the positivekeywords field in the database.
	FieldPositiveKeywords = "positive_keywords"
	// FieldNegativeKeywords holds the string denoting the negativekeywords field in the database.
	FieldNegativeKeywords = "negative_keywords"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldSizes holds the string denoting the sizes field in the database.
	FieldSizes = "sizes"
	// FieldColors holds the string denoting the colors field in the database.
	FieldColors = "colors"
	// FieldSite holds the string denoting the site field in the database.
	FieldSite = "site"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "Task"
	// EdgeCalendar holds the string denoting the calendar edge name in mutations.
	EdgeCalendar = "Calendar"
	// Table holds the table name of the product in the database.
	Table = "products"
	// TaskTable is the table that holds the Task relation/edge. The primary key declared below.
	TaskTable = "task_Product"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// CalendarTable is the table that holds the Calendar relation/edge.
	CalendarTable = "products"
	// CalendarInverseTable is the table name for the Calendar entity.
	// It exists in this package in order to avoid circular dependency with the "calendar" package.
	CalendarInverseTable = "calendars"
	// CalendarColumn is the table column denoting the Calendar relation/edge.
	CalendarColumn = "calendar_quick_task"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldImage,
	FieldLookupType,
	FieldPositiveKeywords,
	FieldNegativeKeywords,
	FieldLink,
	FieldQuantity,
	FieldSizes,
	FieldColors,
	FieldSite,
	FieldMetadata,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "products"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"calendar_quick_task",
}

var (
	// TaskPrimaryKey and TaskColumn2 are the table columns denoting the
	// primary key for the Task relation (M2M).
	TaskPrimaryKey = []string{"task_id", "product_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultQuantity holds the default value on creation for the "Quantity" field.
	DefaultQuantity int32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// LookupType defines the type for the "LookupType" enum field.
type LookupType string

// LookupType values.
const (
	LookupTypeKeywords LookupType = "Keywords"
	LookupTypeLink     LookupType = "Link"
	LookupTypeOther    LookupType = "Other"
)

func (_lookuptype LookupType) String() string {
	return string(_lookuptype)
}

// LookupTypeValidator is a validator for the "LookupType" field enum values. It is called by the builders before save.
func LookupTypeValidator(_lookuptype LookupType) error {
	switch _lookuptype {
	case LookupTypeKeywords, LookupTypeLink, LookupTypeOther:
		return nil
	default:
		return fmt.Errorf("product: invalid enum value for LookupType field: %q", _lookuptype)
	}
}

// Site defines the type for the "Site" enum field.
type Site string

// Site values.
const (
	SiteFinishLine     Site = "FinishLine"
	SiteJD_Sports      Site = "JD_Sports"
	SiteYeezySupply    Site = "YeezySupply"
	SiteSupreme        Site = "Supreme"
	SiteEastbay_US     Site = "Eastbay_US"
	SiteChamps_US      Site = "Champs_US"
	SiteFootaction_US  Site = "Footaction_US"
	SiteFootlocker_US  Site = "Footlocker_US"
	SiteBestbuy        Site = "Bestbuy"
	SitePokemon_Center Site = "Pokemon_Center"
	SitePanini_US      Site = "Panini_US"
	SiteTopss          Site = "Topss"
	SiteNordstorm      Site = "Nordstorm"
	SiteEnd            Site = "End"
	SiteTarget         Site = "Target"
	SiteAmazon         Site = "Amazon"
	SiteSolebox        Site = "Solebox"
	SiteOnygo          Site = "Onygo"
	SiteSnipes         Site = "Snipes"
	SiteSsense         Site = "Ssense"
	SiteWalmart        Site = "Walmart"
	SiteHibbet         Site = "Hibbet"
	SiteNewBalance     Site = "NewBalance"
)

func (_site Site) String() string {
	return string(_site)
}

// SiteValidator is a validator for the "Site" field enum values. It is called by the builders before save.
func SiteValidator(_site Site) error {
	switch _site {
	case SiteFinishLine, SiteJD_Sports, SiteYeezySupply, SiteSupreme, SiteEastbay_US, SiteChamps_US, SiteFootaction_US, SiteFootlocker_US, SiteBestbuy, SitePokemon_Center, SitePanini_US, SiteTopss, SiteNordstorm, SiteEnd, SiteTarget, SiteAmazon, SiteSolebox, SiteOnygo, SiteSnipes, SiteSsense, SiteWalmart, SiteHibbet, SiteNewBalance:
		return nil
	default:
		return fmt.Errorf("product: invalid enum value for Site field: %q", _site)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (_lookuptype LookupType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(_lookuptype.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (_lookuptype *LookupType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*_lookuptype = LookupType(str)
	if err := LookupTypeValidator(*_lookuptype); err != nil {
		return fmt.Errorf("%s is not a valid LookupType", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (_site Site) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(_site.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (_site *Site) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*_site = Site(str)
	if err := SiteValidator(*_site); err != nil {
		return fmt.Errorf("%s is not a valid Site", str)
	}
	return nil
}
