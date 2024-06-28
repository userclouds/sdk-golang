package userstore

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gofrs/uuid"

	"userclouds.com/idp/userstore/selectorconfigparser"
	"userclouds.com/infra/pagination"
	"userclouds.com/infra/ucerr"
	"userclouds.com/infra/uctypes/set"
)

// DataType is an enum for supported data types
type DataType string

// DataType constants
// NOTE: keep in sync with mapDataType defined in TenantUserStoreConfig.tsx
const (
	DataTypeInvalid         DataType = ""
	DataTypeAddress         DataType = "address"
	DataTypeBirthdate       DataType = "birthdate"
	DataTypeBoolean         DataType = "boolean"
	DataTypeComposite       DataType = "composite"
	DataTypeDate            DataType = "date"
	DataTypeEmail           DataType = "email"
	DataTypeInteger         DataType = "integer"
	DataTypeE164PhoneNumber DataType = "e164_phonenumber"
	DataTypePhoneNumber     DataType = "phonenumber"
	DataTypeSSN             DataType = "ssn"
	DataTypeString          DataType = "string"
	DataTypeTimestamp       DataType = "timestamp"
	DataTypeUUID            DataType = "uuid"
)

//go:generate genconstant DataType

// CompositeField represents the settings for a composite data type field
type CompositeField struct {
	DataType            ResourceID `json:"data_type"`
	Name                string     `json:"name" validate:"length:1,128" required:"true" description:"Each part of name must be capitalized or all-caps, separated by underscores. Names may contain alphanumeric characters, and the first part must start with a letter, while other parts may start with a number. (ex. ID_Field_1)"`
	CamelCaseName       string     `json:"camel_case_name" description:"Read-only camel-case version of field name, with underscores stripped out. (ex. IDField1)"`
	StructName          string     `json:"struct_name" description:"Read-only snake-case version of field name, with all letters lowercase. (ex. id_field_1)"`
	Required            bool       `json:"required" description:"Whether a value must be specified for the field."`
	IgnoreForUniqueness bool       `json:"ignore_for_uniqueness" description:"If true, field value will be ignored when comparing two composite values for a uniqueness check."`
}

//go:generate genvalidate CompositeField

// CompositeAttributes represents the attributes for a composite data type
type CompositeAttributes struct {
	IncludeID bool             `json:"include_id" description:"Whether the composite data type must include an id field."`
	Fields    []CompositeField `json:"fields" description:"The set of fields associated with a composite data type."`
}

//go:generate genvalidate CompositeAttributes

// ColumnDataType represents the settings for a data type
type ColumnDataType struct {
	ID                   uuid.UUID           `json:"id"`
	Name                 string              `json:"name" validate:"length:1,128" required:"true"`
	Description          string              `json:"description"`
	IsCompositeFieldType bool                `json:"is_composite_field_type" description:"Whether the data type can be used for a composite field."`
	IsNative             bool                `json:"is_native" description:"Whether this is a native non-editable data type."`
	CompositeAttributes  CompositeAttributes `json:"composite_attributes"`
}

//go:generate genvalidate ColumnDataType

// ColumnField represents the settings for a column field
type ColumnField struct {
	Type                DataType `json:"type" required:"true"`
	Name                string   `json:"name" validate:"length:1,128" required:"true" description:"Each part of name must be capitalized or all-caps, separated by underscores. Names may contain alphanumeric characters, and the first part must start with a letter, while other parts may start with a number. (ex. ID_Field_1)"`
	CamelCaseName       string   `json:"camel_case_name" description:"Read-only camel-case version of field name, with underscores stripped out. (ex. IDField1)"`
	StructName          string   `json:"struct_name" description:"Read-only snake-case version of field name, with all letters lowercase. (ex. id_field_1)"`
	Required            bool     `json:"required" description:"Whether a value must be specified for the field."`
	IgnoreForUniqueness bool     `json:"ignore_for_uniqueness" description:"If true, field value will be ignored when comparing two composite value for a uniqueness check."`
}

//go:generate genvalidate ColumnField

// ColumnConstraints represents the data type constraints for a column
type ColumnConstraints struct {
	ImmutableRequired bool          `json:"immutable_required" description:"Can be enabled when unique_id_required is enabled. If true, values for the associated column cannot be modified, but can be added or removed."`
	PartialUpdates    bool          `json:"partial_updates" description:"Can be enabled for array columns that have UniqueRequired or UniqueIDRequired enabled. When enabled, a mutation request will update the specified subset of values for the associated column."`
	UniqueIDRequired  bool          `json:"unique_id_required" description:"Can be enabled for column type composite or address. If true, each value for the associated column must have a unique string ID, which can either be provided or generated by backend."`
	UniqueRequired    bool          `json:"unique_required" description:"If true, each value for the associated column must be unique for the user. This is primarily useful for array columns."`
	Fields            []ColumnField `json:"fields" description:"The set of fields associated with a column of type composite. Fields cannot be specified if the column type is not composite."`
}

// Equals returns true if the column constraints are equal
func (cc ColumnConstraints) Equals(other ColumnConstraints) bool {
	if cc.ImmutableRequired != other.ImmutableRequired ||
		cc.PartialUpdates != other.PartialUpdates ||
		cc.UniqueIDRequired != other.UniqueIDRequired ||
		cc.UniqueRequired != other.UniqueRequired ||
		len(cc.Fields) != len(other.Fields) {
		return false
	}

	for i, field := range cc.Fields {
		if field != other.Fields[i] {
			return false
		}
	}

	return true
}

//go:generate genvalidate ColumnConstraints

// Address is a native userstore type that represents a physical address
type Address struct {
	ID                 string `json:"id,omitempty"`
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

// NewAddressSet returns a set of addresses
func NewAddressSet(items ...Address) set.Set[Address] {
	return set.New(
		func(items []Address) {
			sort.Slice(items, func(i, j int) bool {
				return fmt.Sprintf("%+v", items[i]) < fmt.Sprintf("%+v", items[j])
			})
		},
		items...,
	)
}

//go:generate gendbjson Address

// CompositeValue is a map of strings to value
type CompositeValue map[string]interface{}

//go:generate gendbjson CompositeValue

// ColumnIndexType is an enum for supported column index types
type ColumnIndexType string

const (
	// ColumnIndexTypeNone is the default value
	ColumnIndexTypeNone ColumnIndexType = "none"

	// ColumnIndexTypeIndexed indicates that the column should be indexed
	ColumnIndexTypeIndexed ColumnIndexType = "indexed"

	// ColumnIndexTypeUnique indicates that the column should be indexed and unique
	ColumnIndexTypeUnique ColumnIndexType = "unique"
)

//go:generate genconstant ColumnIndexType

// Column represents a single field/column/value to be collected/stored/managed
// in the user data store of a tenant.
type Column struct {
	// Columns may be renamed, but their ID cannot be changed.
	ID           uuid.UUID         `json:"id"`
	Table        string            `json:"table"` // TODO (sgarrity 6/24): validate & mark as required once people update
	Name         string            `json:"name" validate:"length:1,128" required:"true"`
	DataType     ResourceID        `json:"data_type" validate:"skip"`
	Type         DataType          `json:"type" required:"true"`
	IsArray      bool              `json:"is_array" required:"true"`
	DefaultValue string            `json:"default_value"`
	IndexType    ColumnIndexType   `json:"index_type" required:"true"`
	IsSystem     bool              `json:"is_system" description:"Whether this column is a system column. System columns cannot be deleted or modified. This property cannot be changed."`
	Constraints  ColumnConstraints `json:"constraints" description:"Optional constraints for configuring the behavior of the associated column Type."`
}

var validIdentifier = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_-]*$`)

func (c *Column) extraValidate() error {
	// TODO (sgarrity 6/24): see above validation comment, remove the != "" check when ready
	if c.Table != "" && !validIdentifier.MatchString(c.Table) {
		return ucerr.Friendlyf(nil, `"%s" is not a valid table name`, c.Table)
	}

	if !validIdentifier.MatchString(c.Name) {
		return ucerr.Friendlyf(nil, `"%s" is not a valid column name`, c.Name)
	}

	return nil
}

//go:generate genvalidate Column

// GetPaginationKeys is part of the pagination.PageableType interface
func (Column) GetPaginationKeys() pagination.KeyTypes {
	return pagination.KeyTypes{
		"id":      pagination.UUIDKeyType,
		"name":    pagination.StringKeyType,
		"created": pagination.TimestampKeyType,
		"updated": pagination.TimestampKeyType,
	}
}

// EqualsIgnoringNilID returns true if the two columns are equal, ignoring ID if one is nil
func (c *Column) EqualsIgnoringNilID(other *Column) bool {
	return (c.ID == other.ID || c.ID.IsNil() || other.ID.IsNil()) &&
		(strings.EqualFold(c.Table, other.Table) ||
			c.Table == "" && other.Table == "users" ||
			c.Table == "users" && other.Table == "") &&
		strings.EqualFold(c.Name, other.Name) &&
		c.Type == other.Type &&
		c.DataType.EquivalentTo(other.DataType) &&
		c.IsArray == other.IsArray &&
		c.DefaultValue == other.DefaultValue &&
		c.IndexType == other.IndexType &&
		c.IsSystem == other.IsSystem &&
		c.Constraints.Equals(other.Constraints)
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
	return typedValue(r, key, false) || r.StringValue(key) == "true"
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

// isCompatibleWith returns true if the resource id is "compatible"
// with the other resource id - that is, that both resource ids
// could refer to the same fully specified resource. Since the
// resource id ID or Name can be unspecified, and thus treated as
// a wildcard, we need to ensure that whatever is specified in one
// resource id does not conflict with what is specified in the other
// resource id.
func (r ResourceID) isCompatibleWith(other ResourceID) bool {
	expectedMatches := 0
	numMatches := 0

	if !r.ID.IsNil() {
		expectedMatches++
		if r.ID == other.ID {
			numMatches++
		} else if !other.ID.IsNil() {
			return false
		}
	}

	if r.Name != "" {
		expectedMatches++
		if strings.EqualFold(r.Name, other.Name) {
			numMatches++
		} else if other.Name != "" {
			return false
		}
	}

	return expectedMatches == 0 || numMatches > 0
}

// EquivalentTo returns true if the resources are compatible with each other
func (r ResourceID) EquivalentTo(other ResourceID) bool {
	return r.isCompatibleWith(other) && other.isCompatibleWith(r)
}

// Validate implements Validateable
func (r ResourceID) Validate() error {
	if r.ID.IsNil() && r.Name == "" {
		return ucerr.Friendlyf(nil, "either ID or Name must be set")
	}
	return nil
}

// ColumnOutputConfig is a struct that contains a column and the transformer to apply to that column
type ColumnOutputConfig struct {
	Column      ResourceID `json:"column"`
	Transformer ResourceID `json:"transformer"`
}

// GetRetentionTimeoutImmediateDeletion returns the immediate deletion retention timeout
func GetRetentionTimeoutImmediateDeletion() time.Time {
	return time.Time{}
}

// GetRetentionTimeoutIndefinite returns the indefinite retention timeout
func GetRetentionTimeoutIndefinite() time.Time {
	return time.Time{}
}

// DataLifeCycleState identifies the life-cycle state for a piece of data - either
// live or soft-deleted.
type DataLifeCycleState string

// Supported data life cycle states
const (
	DataLifeCycleStateDefault     DataLifeCycleState = ""
	DataLifeCycleStateLive        DataLifeCycleState = "live"
	DataLifeCycleStateSoftDeleted DataLifeCycleState = "softdeleted"

	// maps to softdeleted
	DataLifeCycleStatePostDelete DataLifeCycleState = "postdelete"

	// maps to live
	DataLifeCycleStatePreDelete DataLifeCycleState = "predelete"
)

//go:generate genconstant DataLifeCycleState

// GetConcrete returns the concrete data life cycle state for the given data life cycle state
func (dlcs DataLifeCycleState) GetConcrete() DataLifeCycleState {
	switch dlcs {
	case DataLifeCycleStateDefault, DataLifeCycleStatePreDelete:
		return DataLifeCycleStateLive
	case DataLifeCycleStatePostDelete:
		return DataLifeCycleStateSoftDeleted
	default:
		return dlcs
	}
}

// GetDefaultRetentionTimeout returns the default retention timeout for the data life cycle state
func (dlcs DataLifeCycleState) GetDefaultRetentionTimeout() time.Time {
	if dlcs.GetConcrete() == DataLifeCycleStateLive {
		return GetRetentionTimeoutIndefinite()
	}

	return GetRetentionTimeoutImmediateDeletion()
}

// IsLive return true if the concrete data life cycle state is live
func (dlcs DataLifeCycleState) IsLive() bool {
	return dlcs.GetConcrete() == DataLifeCycleStateLive
}

// Accessor represents a customer-defined view and permissions policy on userstore data
type Accessor struct {
	ID uuid.UUID `json:"id"`

	// Name of accessor, must be unique
	Name string `json:"name" validate:"length:1,128" required:"true"`

	// Description of the accessor
	Description string `json:"description"`

	// Version of the accessor
	Version int `json:"version"`

	// Specify whether to access live or soft-deleted data
	DataLifeCycleState DataLifeCycleState `json:"data_life_cycle_state"`

	// Configuration for which user records to return
	SelectorConfig UserSelectorConfig `json:"selector_config" required:"true"`

	// Purposes for which this accessor is used
	Purposes []ResourceID `json:"purposes" validate:"skip" required:"true"`

	// List of userstore columns being accessed and the transformers to apply to each column
	Columns []ColumnOutputConfig `json:"columns" validate:"skip" required:"true"`

	// Policy for what data is returned by this accessor, based on properties of the caller and the user records
	AccessPolicy ResourceID `json:"access_policy" validate:"skip" required:"true"`

	// Policy for token resolution in the case of transformers that tokenize data
	TokenAccessPolicy ResourceID `json:"token_access_policy,omitempty" validate:"skip"`

	// Whether this accessor is a system accessor
	IsSystem bool `json:"is_system" description:"Whether this accessor is a system accessor. System accessors cannot be deleted or modified. This property cannot be changed."`

	// Whether this accessor is audit logged each time it is executed
	IsAuditLogged bool `json:"is_audit_logged" description:"Whether this accessor is audit logged each time it is executed."`
}

func (o *Accessor) extraValidate() error {

	if !validIdentifier.MatchString(string(o.Name)) {
		return ucerr.Friendlyf(nil, `"%s" is not a valid accessor name`, o.Name)
	}

	if len(o.Columns) == 0 {
		return ucerr.Friendlyf(nil, "Accessor.Columns (%v) can't be empty", o.ID)
	}

	for _, ct := range o.Columns {
		if err := ct.Column.Validate(); err != nil {
			return ucerr.Friendlyf(err, "Each element of Accessor.Columns (%v) must have a column ID or name", o.ID)
		}

		if err := ct.Transformer.Validate(); err != nil {
			return ucerr.Friendlyf(err, "Each element of Accessor.Columns (%v) must have a transformer ID or name", o.ID)
		}
	}

	if err := o.AccessPolicy.Validate(); err != nil {
		return ucerr.Friendlyf(err, "Accessor.AccessPolicy (%v) must have an ID or name", o.ID)
	}

	if len(o.Purposes) == 0 {
		return ucerr.Friendlyf(nil, "Accessor.Purposes (%v) can't be empty", o.ID)
	}

	return nil
}

//go:generate genvalidate Accessor

// GetPaginationKeys is part of the pagination.PageableType interface
func (Accessor) GetPaginationKeys() pagination.KeyTypes {
	return pagination.KeyTypes{
		"id":      pagination.UUIDKeyType,
		"name":    pagination.StringKeyType,
		"created": pagination.TimestampKeyType,
		"updated": pagination.TimestampKeyType,
	}
}

// ColumnInputConfig is a struct that contains a column and the normalizer to use for that column
type ColumnInputConfig struct {
	Column     ResourceID `json:"column"`
	Normalizer ResourceID `json:"normalizer"`

	// Validator is deprecated in favor of Normalizer
	Validator ResourceID `json:"validator"`
}

// Mutator represents a customer-defined scope and permissions policy for updating userstore data
type Mutator struct {
	ID uuid.UUID `json:"id"`

	// Name of mutator, must be unique
	Name string `json:"name" validate:"length:1,128" required:"true"`

	// Description of the mutator
	Description string `json:"description"`

	// Version of the mutator
	Version int `json:"version"`

	// Configuration for which user records to modify
	SelectorConfig UserSelectorConfig `json:"selector_config" required:"true"`

	// The set of userstore columns to modify for each user record
	Columns []ColumnInputConfig `json:"columns" validate:"skip" required:"true"`

	// Policy for whether the data for each user record can be updated
	AccessPolicy ResourceID `json:"access_policy" validate:"skip" required:"true"`

	IsSystem bool `json:"is_system" description:"Whether this mutator is a system mutator. System mutators cannot be deleted or modified. This property cannot be changed."`
}

func (o *Mutator) extraValidate() error {

	if !validIdentifier.MatchString(string(o.Name)) {
		return ucerr.Friendlyf(nil, `"%s" is not a valid mutator name`, o.Name)
	}

	totalColumns := len(o.Columns)
	if totalColumns == 0 && !o.IsSystem {
		return ucerr.Friendlyf(nil, "Mutator with ID (%v) can't have empty Columns", o.ID)
	}

	totalValidNormalizers := 0
	totalValidValidators := 0
	for _, cv := range o.Columns {
		if err := cv.Column.Validate(); err != nil {
			return ucerr.Friendlyf(err, "Mutator with ID (%v): each element of Columns must have a column ID or name", o.ID)
		}

		if err := cv.Normalizer.Validate(); err == nil {
			totalValidNormalizers++
		}

		if err := cv.Validator.Validate(); err == nil {
			totalValidValidators++
		}
	}

	if totalValidNormalizers != totalColumns && totalValidValidators != totalColumns {
		return ucerr.Friendlyf(nil, "Mutator with ID (%v): each element of Columns must have either a normalizer or validator ID or name", o.ID)
	}

	if err := o.AccessPolicy.Validate(); err != nil {
		return ucerr.Friendlyf(err, "Mutator with ID (%v): AccessPolicy must have an ID or name", o.ID)
	}

	return nil
}

//go:generate genvalidate Mutator

// UserSelectorValues are the values passed for the UserSelector of an accessor or mutator
type UserSelectorValues []interface{}

// UserSelectorConfig is the configuration for a UserSelector
type UserSelectorConfig struct {
	WhereClause string `json:"where_clause" validate:"notempty" example:"{id} = ANY (?)"`
}

// MatchesAll returns true if the UserSelectorConfig is configured to match all users
func (u UserSelectorConfig) MatchesAll() bool {
	return u.WhereClause == "ALL"
}

func (u UserSelectorConfig) extraValidate() error {
	if u.MatchesAll() {
		return nil
	}
	return ucerr.Wrap(selectorconfigparser.ParseWhereClause(u.WhereClause))
}

//go:generate gendbjson UserSelectorConfig

//go:generate genvalidate UserSelectorConfig

// Purpose represents a customer-defined purpose for userstore columns
type Purpose struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"length:1,128" required:"true"`
	Description string    `json:"description"`
	IsSystem    bool      `json:"is_system" description:"Whether this purpose is a system purpose. System purposes cannot be deleted or modified. This property cannot be changed."`
}

func (p *Purpose) extraValidate() error {

	if !validIdentifier.MatchString(string(p.Name)) {
		return ucerr.Friendlyf(nil, `"%s" is not a valid purpose name`, p.Name)
	}

	return nil
}

//go:generate genvalidate Purpose

// SQLShimDatabase represents an external database that tenant customers can connect to via a SQLShim proxy
type SQLShimDatabase struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name" validate:"notempty"`
	Type     string    `json:"type" validate:"notempty"`
	Host     string    `json:"host" validate:"notempty"`
	Port     int       `json:"port" validate:"notzero"`
	Username string    `json:"username" validate:"notempty"`
	Password string    `json:"password" validate:"skip"`
}

//go:generate genvalidate SQLShimDatabase

// EqualsIgnoringNilIDAndPassword returns true if the two columns are equal, ignoring ID if one is nil, and ignoring password field
func (s SQLShimDatabase) EqualsIgnoringNilIDAndPassword(other SQLShimDatabase) bool {
	return (s.ID == other.ID || s.ID.IsNil() || other.ID.IsNil()) &&
		strings.EqualFold(s.Name, other.Name) &&
		strings.EqualFold(s.Type, other.Type) &&
		strings.EqualFold(s.Host, other.Host) &&
		s.Port == other.Port &&
		strings.EqualFold(s.Username, other.Username)
}
