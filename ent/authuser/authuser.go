// Code generated by entc, DO NOT EDIT.

package authuser

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the authuser type in the database.
	Label = "auth_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldCompany holds the string denoting the company field in the database.
	FieldCompany = "company"
	// FieldOccupation holds the string denoting the occupation field in the database.
	FieldOccupation = "occupation"
	// FieldPrivateKeyPath holds the string denoting the private_key_path field in the database.
	FieldPrivateKeyPath = "private_key_path"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldProvider holds the string denoting the provider field in the database.
	FieldProvider = "provider"
	// EdgeAuthUserToToken holds the string denoting the authusertotoken edge name in mutations.
	EdgeAuthUserToToken = "AuthUserToToken"
	// EdgeAuthUserToServerTasks holds the string denoting the authusertoservertasks edge name in mutations.
	EdgeAuthUserToServerTasks = "AuthUserToServerTasks"
	// Table holds the table name of the authuser in the database.
	Table = "auth_users"
	// AuthUserToTokenTable is the table the holds the AuthUserToToken relation/edge.
	AuthUserToTokenTable = "tokens"
	// AuthUserToTokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	AuthUserToTokenInverseTable = "tokens"
	// AuthUserToTokenColumn is the table column denoting the AuthUserToToken relation/edge.
	AuthUserToTokenColumn = "auth_user_auth_user_to_token"
	// AuthUserToServerTasksTable is the table the holds the AuthUserToServerTasks relation/edge.
	AuthUserToServerTasksTable = "server_tasks"
	// AuthUserToServerTasksInverseTable is the table name for the ServerTask entity.
	// It exists in this package in order to avoid circular dependency with the "servertask" package.
	AuthUserToServerTasksInverseTable = "server_tasks"
	// AuthUserToServerTasksColumn is the table column denoting the AuthUserToServerTasks relation/edge.
	AuthUserToServerTasksColumn = "server_task_server_task_to_auth_user"
)

// Columns holds all SQL columns for authuser fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldFirstName,
	FieldLastName,
	FieldEmail,
	FieldPhone,
	FieldCompany,
	FieldOccupation,
	FieldPrivateKeyPath,
	FieldRole,
	FieldProvider,
}

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
	// DefaultFirstName holds the default value on creation for the "first_name" field.
	DefaultFirstName string
	// DefaultLastName holds the default value on creation for the "last_name" field.
	DefaultLastName string
	// DefaultEmail holds the default value on creation for the "email" field.
	DefaultEmail string
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// DefaultCompany holds the default value on creation for the "company" field.
	DefaultCompany string
	// DefaultOccupation holds the default value on creation for the "occupation" field.
	DefaultOccupation string
	// DefaultPrivateKeyPath holds the default value on creation for the "private_key_path" field.
	DefaultPrivateKeyPath string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Role defines the type for the "role" enum field.
type Role string

// Role values.
const (
	RoleUSER  Role = "USER"
	RoleADMIN Role = "ADMIN"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleUSER, RoleADMIN:
		return nil
	default:
		return fmt.Errorf("authuser: invalid enum value for role field: %q", r)
	}
}

// Provider defines the type for the "provider" enum field.
type Provider string

// Provider values.
const (
	ProviderLOCAL  Provider = "LOCAL"
	ProviderGITHUB Provider = "GITHUB"
	ProviderOPENID Provider = "OPENID"
)

func (pr Provider) String() string {
	return string(pr)
}

// ProviderValidator is a validator for the "provider" field enum values. It is called by the builders before save.
func ProviderValidator(pr Provider) error {
	switch pr {
	case ProviderLOCAL, ProviderGITHUB, ProviderOPENID:
		return nil
	default:
		return fmt.Errorf("authuser: invalid enum value for provider field: %q", pr)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (r Role) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(r.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (r *Role) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*r = Role(str)
	if err := RoleValidator(*r); err != nil {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (pr Provider) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(pr.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (pr *Provider) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*pr = Provider(str)
	if err := ProviderValidator(*pr); err != nil {
		return fmt.Errorf("%s is not a valid Provider", str)
	}
	return nil
}
