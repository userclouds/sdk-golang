package policy

import (
	"encoding/json"
	"regexp"

	"github.com/gofrs/uuid"

	"userclouds.com/infra/ucerr"
)

var validIdentifier = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_-]*$`)

const maxIdentifierLength = 128

// TransformationPolicy describes a token transformation policy
type TransformationPolicy struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name" validate:"notempty"`
	Function   string    `json:"function" validate:"notempty"`
	Parameters string    `json:"parameters"`
}

//go:generate genvalidate TransformationPolicy

func (g TransformationPolicy) extraValidate() error {

	if len(g.Name) > maxIdentifierLength || !validIdentifier.MatchString(string(g.Name)) {
		return ucerr.Friendlyf(nil, `Transformation policy name "%s" is too long or has invalid characters`, g.Name)
	}

	params := map[string]interface{}{}
	if err := json.Unmarshal([]byte(g.Parameters), &params); g.Parameters != "" && err != nil {
		paramsArr := []interface{}{}
		if err := json.Unmarshal([]byte(g.Parameters), &paramsArr); err != nil {
			return ucerr.New("TransformationPolicy.Parameters must be either empty, or a JSON dictionary or JSON array")
		}
	}

	return nil
}

// Equals returns true if the two policies are equal, ignoring the ID field
func (g *TransformationPolicy) Equals(other *TransformationPolicy) bool {
	return (g.ID == other.ID || g.ID == uuid.Nil || other.ID == uuid.Nil) &&
		g.Name == other.Name && g.Function == other.Function && g.Parameters == other.Parameters
}

// AccessPolicy describes a token transformation policy
type AccessPolicy struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name" validate:"notempty"`
	Function   string    `json:"function" validate:"notempty"`
	Parameters string    `json:"parameters"`
	Version    int       `json:"version"` // NB: this is currently emitted by the server, but not read by the server (for UI only)
}

//go:generate genvalidate AccessPolicy

func (a AccessPolicy) extraValidate() error {

	if len(a.Name) > maxIdentifierLength || !validIdentifier.MatchString(string(a.Name)) {
		return ucerr.Friendlyf(nil, `Access policy name "%s" is too long or has invalid characters`, a.Name)
	}

	params := map[string]interface{}{}
	if err := json.Unmarshal([]byte(a.Parameters), &params); a.Parameters != "" && err != nil {
		return ucerr.New("AccessPolicy.Parameters must be either empty, or a JSON dictionary")
	}

	return nil
}

// Equals returns true if the two policies are equal, ignoring the ID field
func (a *AccessPolicy) Equals(other *AccessPolicy) bool {
	return (a.ID == other.ID || a.ID == uuid.Nil || other.ID == uuid.Nil) &&
		a.Name == other.Name && a.Function == other.Function && a.Parameters == other.Parameters
}

// ClientContext is passed by the client at resolution time
type ClientContext map[string]interface{}

// AccessPolicyContext gets passed to the access policy's function(context, params) at resolution time
type AccessPolicyContext struct {
	Server ServerContext `json:"server"`
	Client ClientContext `json:"client"`
}

// ServerContext is automatically injected by the server at resolution time
type ServerContext struct {
	// TODO: add token creation time
	IPAddress string          `json:"ip_address"`
	Resolver  ResolverContext `json:"resolver"`
	Action    Action          `json:"action"`
}

// ResolverContext contains automatic data about the authenticated user/system at resolution time
type ResolverContext struct {
	Username string `json:"username"`
}

// Action identifies the reason access policy is being invoked
type Action string

// Different reasons for running access policy
const (
	ActionResolve Action = "Resolve"
	ActionInspect Action = "Inspect"
	ActionLookup  Action = "Lookup"
	ActionDelete  Action = "Delete"
	ActionExecute Action = "Execute" // TODO: should this be a unique action?
)
