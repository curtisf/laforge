// Code generated by entc, DO NOT EDIT.

package plan

import (
	"fmt"
	"io"
	"strconv"
)

const (
	// Label holds the string label denoting the plan type in the database.
	Label = "plan"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStepNumber holds the string denoting the step_number field in the database.
	FieldStepNumber = "step_number"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldBuildID holds the string denoting the build_id field in the database.
	FieldBuildID = "build_id"

	// EdgePrevPlan holds the string denoting the prevplan edge name in mutations.
	EdgePrevPlan = "PrevPlan"
	// EdgeNextPlan holds the string denoting the nextplan edge name in mutations.
	EdgeNextPlan = "NextPlan"
	// EdgePlanToBuild holds the string denoting the plantobuild edge name in mutations.
	EdgePlanToBuild = "PlanToBuild"
	// EdgePlanToTeam holds the string denoting the plantoteam edge name in mutations.
	EdgePlanToTeam = "PlanToTeam"
	// EdgePlanToProvisionedNetwork holds the string denoting the plantoprovisionednetwork edge name in mutations.
	EdgePlanToProvisionedNetwork = "PlanToProvisionedNetwork"
	// EdgePlanToProvisionedHost holds the string denoting the plantoprovisionedhost edge name in mutations.
	EdgePlanToProvisionedHost = "PlanToProvisionedHost"
	// EdgePlanToProvisioningStep holds the string denoting the plantoprovisioningstep edge name in mutations.
	EdgePlanToProvisioningStep = "PlanToProvisioningStep"

	// Table holds the table name of the plan in the database.
	Table = "plans"
	// PrevPlanTable is the table the holds the PrevPlan relation/edge. The primary key declared below.
	PrevPlanTable = "plan_NextPlan"
	// NextPlanTable is the table the holds the NextPlan relation/edge. The primary key declared below.
	NextPlanTable = "plan_NextPlan"
	// PlanToBuildTable is the table the holds the PlanToBuild relation/edge.
	PlanToBuildTable = "plans"
	// PlanToBuildInverseTable is the table name for the Build entity.
	// It exists in this package in order to avoid circular dependency with the "build" package.
	PlanToBuildInverseTable = "builds"
	// PlanToBuildColumn is the table column denoting the PlanToBuild relation/edge.
	PlanToBuildColumn = "plan_plan_to_build"
	// PlanToTeamTable is the table the holds the PlanToTeam relation/edge.
	PlanToTeamTable = "plans"
	// PlanToTeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	PlanToTeamInverseTable = "teams"
	// PlanToTeamColumn is the table column denoting the PlanToTeam relation/edge.
	PlanToTeamColumn = "plan_plan_to_team"
	// PlanToProvisionedNetworkTable is the table the holds the PlanToProvisionedNetwork relation/edge.
	PlanToProvisionedNetworkTable = "provisioned_networks"
	// PlanToProvisionedNetworkInverseTable is the table name for the ProvisionedNetwork entity.
	// It exists in this package in order to avoid circular dependency with the "provisionednetwork" package.
	PlanToProvisionedNetworkInverseTable = "provisioned_networks"
	// PlanToProvisionedNetworkColumn is the table column denoting the PlanToProvisionedNetwork relation/edge.
	PlanToProvisionedNetworkColumn = "plan_plan_to_provisioned_network"
	// PlanToProvisionedHostTable is the table the holds the PlanToProvisionedHost relation/edge.
	PlanToProvisionedHostTable = "plans"
	// PlanToProvisionedHostInverseTable is the table name for the ProvisionedHost entity.
	// It exists in this package in order to avoid circular dependency with the "provisionedhost" package.
	PlanToProvisionedHostInverseTable = "provisioned_hosts"
	// PlanToProvisionedHostColumn is the table column denoting the PlanToProvisionedHost relation/edge.
	PlanToProvisionedHostColumn = "plan_plan_to_provisioned_host"
	// PlanToProvisioningStepTable is the table the holds the PlanToProvisioningStep relation/edge.
	PlanToProvisioningStepTable = "provisioning_steps"
	// PlanToProvisioningStepInverseTable is the table name for the ProvisioningStep entity.
	// It exists in this package in order to avoid circular dependency with the "provisioningstep" package.
	PlanToProvisioningStepInverseTable = "provisioning_steps"
	// PlanToProvisioningStepColumn is the table column denoting the PlanToProvisioningStep relation/edge.
	PlanToProvisioningStepColumn = "plan_plan_to_provisioning_step"
)

// Columns holds all SQL columns for plan fields.
var Columns = []string{
	FieldID,
	FieldStepNumber,
	FieldType,
	FieldBuildID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Plan type.
var ForeignKeys = []string{
	"plan_plan_to_build",
	"plan_plan_to_team",
	"plan_plan_to_provisioned_host",
}

var (
	// PrevPlanPrimaryKey and PrevPlanColumn2 are the table columns denoting the
	// primary key for the PrevPlan relation (M2M).
	PrevPlanPrimaryKey = []string{"plan_id", "PrevPlan_id"}
	// NextPlanPrimaryKey and NextPlanColumn2 are the table columns denoting the
	// primary key for the NextPlan relation (M2M).
	NextPlanPrimaryKey = []string{"plan_id", "PrevPlan_id"}
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

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeStartBuild       Type = "start_build"
	TypeStartTeam        Type = "start_team"
	TypeProvisionNetwork Type = "provision_network"
	TypeProvisionHost    Type = "provision_host"
	TypeExecuteStep      Type = "execute_step"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeStartBuild, TypeStartTeam, TypeProvisionNetwork, TypeProvisionHost, TypeExecuteStep:
		return nil
	default:
		return fmt.Errorf("plan: invalid enum value for type field: %q", _type)
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
