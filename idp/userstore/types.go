package userstore

import (
	"regexp"
	"strings"

	"github.com/gofrs/uuid"

	"userclouds.com/infra/ucerr"
)

// DataType is an enum for supported data types
type DataType int

// DataType constants (leaving gaps intentionally to allow future related types to be grouped)
// NOTE: keep in sync with mapDataType defined in TenantUserStoreConfig.tsx
const (
	DataTypeInvalid DataType = 0
	DataTypeBoolean DataType = 1
	DataTypeInteger DataType = 2

	DataTypeString DataType = 100

	DataTypeTimestamp DataType = 200
	//DataTypeBirthdate DataType = 201

	DataTypeUUID DataType = 300
	//DataTypeSSN DataType = 301
	//DataTypeCreditCardNumber DataType = 302

	//DataTypeJSONB   DataType = 400
	DataTypeAddress DataType = 401
	//DataTypePhoneNumber DataType = 402
)

//go:generate genconstant DataType

// Address is a native userstore type that represents a physical address
type Address struct {
	Country            string `json:"country,omitempty"`
	Name               string `json:"name,omitempty"`
	Organization       string `json:"organization,omitempty"`
	StreetAddressLine1 string `json:"street_address_line_1,omitempty"`
	StreetAddressLine2 string `json:"street_address_line_2,omitempty"`
	DependentLocality  string `json:"dependent_locality,omitempty"`
	Locality           string `json:"locality,omitempty"`
	AdministrativeArea string `json:"administrative_area,omitempty"`
	PostCode           string `json:"post_code,omitempty"`
	SortingCode        string `json:"sorting_code,omitempty"`
}

//go:generate gendbjson Address

// ColumnIndexType is an enum for supported column index types
type ColumnIndexType int

const (
	// ColumnIndexTypeNone is the default value
	ColumnIndexTypeNone ColumnIndexType = iota

	// ColumnIndexTypeIndexed indicates that the column should be indexed
	ColumnIndexTypeIndexed

	// ColumnIndexTypeUnique indicates that the column should be indexed and unique
	ColumnIndexTypeUnique
)

//go:generate genconstant ColumnIndexType

// Column represents a single field/column/value to be collected/stored/managed
// in the user data store of a tenant.
type Column struct {
	// Columns may be renamed, but their ID cannot be changed.
	ID           uuid.UUID       `json:"id"`
	Name         string          `json:"name" validate:"length:1,128"`
	Type         DataType        `json:"type"`
	IsArray      bool            `json:"is_array"`
	DefaultValue string          `json:"default_value"`
	IndexType    ColumnIndexType `json:"index_type"`
}

var validIdentifier = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_-]*$`)

func (c *Column) extraValidate() error {

	if !validIdentifier.MatchString(string(c.Name)) {
		return ucerr.Friendlyf(nil, `"%s" is not a valid column name`, c.Name)
	}

	return nil
}

//go:generate genvalidate Column

// Equals returns true if the two columns are equal
func (c *Column) Equals(other *Column) bool {
	return (c.ID == other.ID || c.ID == uuid.Nil || other.ID == uuid.Nil) &&
		c.Name == other.Name &&
		c.Type == other.Type &&
		c.IsArray == other.IsArray &&
		c.DefaultValue == other.DefaultValue &&
		c.IndexType == other.IndexType
}

// Record is a single "row" of data containing 0 or more Columns from userstore's schema
// The key is the name of the column
type Record map[string]interface{}

func typedValue[T any](r Record, key string, defaultValue T) T {
	if r[key] != nil {
		if value, ok := r[key].(T); ok {
			return value
		}
	}

	return defaultValue
}

// BoolValue returns a boolean value for the specified key
func (r Record) BoolValue(key string) bool {
	return typedValue(r, key, false)
}

// StringValue returns a string value for the specified key
func (r Record) StringValue(key string) string {
	return typedValue(r, key, "")
}

// UUIDValue returns a UUID value for the specified key
func (r Record) UUIDValue(key string) uuid.UUID {
	value, err := uuid.FromString(r.StringValue(key))
	if err != nil {
		return uuid.Nil
	}
	return value
}

//go:generate gendbjson Record

// ResourceID is a struct that contains a name and ID, only one of which is required to be set
type ResourceID struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Validate implements Validateable
func (r ResourceID) Validate() error {
	if r.ID == uuid.Nil && r.Name == "" {
		return ucerr.Friendlyf(nil, "either ID or Name must be set")
	}
	return nil
}

// ColumnOutputConfig is a struct that contains a column and the transformer to apply to that column
type ColumnOutputConfig struct {
	Column      ResourceID `json:"column"`
	Transformer ResourceID `json:"transformer"`
}

// Accessor represents a customer-defined view and permissions policy on userstore data
type Accessor struct {
	ID uuid.UUID `json:"id"`

	// Name of accessor, must be unique
	Name string `json:"name" validate:"length:1,128"`

	// Description of the accessor
	Description string `json:"description"`

	// Version of the accessor
	Version int `json:"version"`

	// Configuration for which user records to return
	SelectorConfig UserSelectorConfig `json:"selector_config"`

	// Purposes for which this accessor is used
	Purposes []ResourceID `json:"purposes" validate:"skip"`

	// List of userstore columns being accessed and the transformers to apply to each column
	Columns []ColumnOutputConfig `json:"columns" validate:"skip"`

	// Policy for what data is returned by this accessor, based on properties of the caller and the user records
	AccessPolicy ResourceID `json:"access_policy" validate:"skip"`

	// Policy for token resolution in the case of transformers that tokenize data
	TokenAccessPolicy ResourceID `json:"token_access_policy,omitempty" validate:"skip"`
}

func (o *Accessor) extraValidate() error {

	if !validIdentifier.MatchString(string(o.Name)) {
		return ucerr.Friendlyf(nil, `"%s" is not a valid accessor name`, o.Name)
	}

	if len(o.Columns) == 0 {
		return ucerr.Errorf("Accessor.Columns (%v) can't be empty", o.ID)
	}

	for _, ct := range o.Columns {
		if ct.Column.ID == uuid.Nil && ct.Column.Name == "" {
			return ucerr.Errorf("Each element of Accessor.Columns (%v) must have a column ID or name", o.ID)
		}

		if ct.Transformer.ID == uuid.Nil && ct.Transformer.Name == "" {
			return ucerr.Errorf("Each element of Accessor.Columns (%v) must have a transformer ID or name", o.ID)
		}
	}

	if o.AccessPolicy.ID == uuid.Nil && o.AccessPolicy.Name == "" {
		return ucerr.Errorf("Accessor.AccessPolicy (%v) must have an ID or name", o.ID)
	}

	if len(o.Purposes) == 0 {
		return ucerr.Errorf("Accessor.Purposes (%v) can't be empty", o.ID)
	}

	return nil
}

//go:generate genvalidate Accessor

// ColumnInputConfig is a struct that contains a column and the validator to use for that column
type ColumnInputConfig struct {
	Column    ResourceID `json:"column"`
	Validator ResourceID `json:"validator"`
}

// Mutator represents a customer-defined scope and permissions policy for updating userstore data
type Mutator struct {
	ID uuid.UUID `json:"id"`

	// Name of mutator, must be unique
	Name string `json:"name" validate:"length:1,128"`

	// Description of the mutator
	Description string `json:"description"`

	// Version of the mutator
	Version int `json:"version"`

	// Configuration for which user records to modify
	SelectorConfig UserSelectorConfig `json:"selector_config"`

	// The set of userstore columns to modify for each user record
	Columns []ColumnInputConfig `json:"columns" validate:"skip"`

	// Policy for whether the data for each user record can be updated
	AccessPolicy ResourceID `json:"access_policy" validate:"skip"`
}

func (o *Mutator) extraValidate() error {

	if !validIdentifier.MatchString(string(o.Name)) {
		return ucerr.Friendlyf(nil, `"%s" is not a valid mutator name`, o.Name)
	}

	if len(o.Columns) == 0 {
		return ucerr.Errorf("Mutator.Columns (%v) can't be empty", o.ID)
	}

	for _, cv := range o.Columns {
		if cv.Column.ID == uuid.Nil && cv.Column.Name == "" {
			return ucerr.Errorf("Each element of Mutator.Columns (%v) must have a column ID or name", o.ID)
		}

		if cv.Validator.ID == uuid.Nil && cv.Validator.Name == "" {
			return ucerr.Errorf("Each element of Mutator.Columns (%v) must have a validator ID or name", o.ID)
		}
	}

	if o.AccessPolicy.ID == uuid.Nil && o.AccessPolicy.Name == "" {
		return ucerr.Errorf("Mutator.AccessPolicy (%v) must have an ID or name", o.ID)
	}

	return nil
}

//go:generate genvalidate Mutator

// UserSelectorValues are the values passed for the UserSelector of an accessor or mutator
type UserSelectorValues []interface{}

// UserSelectorConfig is the configuration for a UserSelector
type UserSelectorConfig struct {
	WhereClause string `json:"where_clause" validate:"notempty"`
}

func (u UserSelectorConfig) extraValidate() error {
	// make sure the where clause only contains tokens for clauses of the form "{column_id} operator ? [conjunction {column_id} operator ?]*"
	// e.g. "{id} = ANY (?) OR {phone_number} LIKE ?"
	columnsRE := regexp.MustCompile(`{[a-zA-Z0-9_-]+}(->>'[a-zA-Z0-9_-]+')?`)
	operatorRE := regexp.MustCompile(`(?i) (=|<|>|<=|>=|!=|LIKE|ANY)`)
	valuesRE := regexp.MustCompile(`\?|\(\?\)`)
	conjunctionRE := regexp.MustCompile(`(?i) (OR|AND) `)
	if s := strings.TrimSpace(conjunctionRE.ReplaceAllString(operatorRE.ReplaceAllString(valuesRE.ReplaceAllString(columnsRE.ReplaceAllString(u.WhereClause, ""), ""), ""), "")); s != "" {
		return ucerr.Friendlyf(nil, `invalid or unsupported SQL in UserSelectorConfig.WhereClause: "%s", near "%s"`, u.WhereClause, strings.Split(s, " ")[0])
	}
	return nil
}

//go:generate gendbjson UserSelectorConfig

//go:generate genvalidate UserSelectorConfig

// Purpose represents a customer-defined purpose for userstore columns
type Purpose struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"length:1,128"`
	Description string    `json:"description"`
}

func (p *Purpose) extraValidate() error {

	if !validIdentifier.MatchString(string(p.Name)) {
		return ucerr.Friendlyf(nil, `"%s" is not a valid purpose name`, p.Name)
	}

	return nil
}

//go:generate genvalidate Purpose
