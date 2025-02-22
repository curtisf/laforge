// Code generated by entc, DO NOT EDIT.

package build

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the build type in the database.
	Label = "build"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRevision holds the string denoting the revision field in the database.
	FieldRevision = "revision"
	// FieldEnvironmentRevision holds the string denoting the environment_revision field in the database.
	FieldEnvironmentRevision = "environment_revision"
	// FieldCompletedPlan holds the string denoting the completed_plan field in the database.
	FieldCompletedPlan = "completed_plan"
	// EdgeBuildToStatus holds the string denoting the buildtostatus edge name in mutations.
	EdgeBuildToStatus = "BuildToStatus"
	// EdgeBuildToEnvironment holds the string denoting the buildtoenvironment edge name in mutations.
	EdgeBuildToEnvironment = "BuildToEnvironment"
	// EdgeBuildToCompetition holds the string denoting the buildtocompetition edge name in mutations.
	EdgeBuildToCompetition = "BuildToCompetition"
	// EdgeBuildToLatestBuildCommit holds the string denoting the buildtolatestbuildcommit edge name in mutations.
	EdgeBuildToLatestBuildCommit = "BuildToLatestBuildCommit"
	// EdgeBuildToProvisionedNetwork holds the string denoting the buildtoprovisionednetwork edge name in mutations.
	EdgeBuildToProvisionedNetwork = "BuildToProvisionedNetwork"
	// EdgeBuildToTeam holds the string denoting the buildtoteam edge name in mutations.
	EdgeBuildToTeam = "BuildToTeam"
	// EdgeBuildToPlan holds the string denoting the buildtoplan edge name in mutations.
	EdgeBuildToPlan = "BuildToPlan"
	// EdgeBuildToBuildCommits holds the string denoting the buildtobuildcommits edge name in mutations.
	EdgeBuildToBuildCommits = "BuildToBuildCommits"
	// EdgeBuildToAdhocPlans holds the string denoting the buildtoadhocplans edge name in mutations.
	EdgeBuildToAdhocPlans = "BuildToAdhocPlans"
	// Table holds the table name of the build in the database.
	Table = "builds"
	// BuildToStatusTable is the table that holds the BuildToStatus relation/edge.
	BuildToStatusTable = "status"
	// BuildToStatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	BuildToStatusInverseTable = "status"
	// BuildToStatusColumn is the table column denoting the BuildToStatus relation/edge.
	BuildToStatusColumn = "build_build_to_status"
	// BuildToEnvironmentTable is the table that holds the BuildToEnvironment relation/edge.
	BuildToEnvironmentTable = "builds"
	// BuildToEnvironmentInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	BuildToEnvironmentInverseTable = "environments"
	// BuildToEnvironmentColumn is the table column denoting the BuildToEnvironment relation/edge.
	BuildToEnvironmentColumn = "build_build_to_environment"
	// BuildToCompetitionTable is the table that holds the BuildToCompetition relation/edge.
	BuildToCompetitionTable = "builds"
	// BuildToCompetitionInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	BuildToCompetitionInverseTable = "competitions"
	// BuildToCompetitionColumn is the table column denoting the BuildToCompetition relation/edge.
	BuildToCompetitionColumn = "build_build_to_competition"
	// BuildToLatestBuildCommitTable is the table that holds the BuildToLatestBuildCommit relation/edge.
	BuildToLatestBuildCommitTable = "builds"
	// BuildToLatestBuildCommitInverseTable is the table name for the BuildCommit entity.
	// It exists in this package in order to avoid circular dependency with the "buildcommit" package.
	BuildToLatestBuildCommitInverseTable = "build_commits"
	// BuildToLatestBuildCommitColumn is the table column denoting the BuildToLatestBuildCommit relation/edge.
	BuildToLatestBuildCommitColumn = "build_build_to_latest_build_commit"
	// BuildToProvisionedNetworkTable is the table that holds the BuildToProvisionedNetwork relation/edge.
	BuildToProvisionedNetworkTable = "provisioned_networks"
	// BuildToProvisionedNetworkInverseTable is the table name for the ProvisionedNetwork entity.
	// It exists in this package in order to avoid circular dependency with the "provisionednetwork" package.
	BuildToProvisionedNetworkInverseTable = "provisioned_networks"
	// BuildToProvisionedNetworkColumn is the table column denoting the BuildToProvisionedNetwork relation/edge.
	BuildToProvisionedNetworkColumn = "provisioned_network_provisioned_network_to_build"
	// BuildToTeamTable is the table that holds the BuildToTeam relation/edge.
	BuildToTeamTable = "teams"
	// BuildToTeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	BuildToTeamInverseTable = "teams"
	// BuildToTeamColumn is the table column denoting the BuildToTeam relation/edge.
	BuildToTeamColumn = "team_team_to_build"
	// BuildToPlanTable is the table that holds the BuildToPlan relation/edge.
	BuildToPlanTable = "plans"
	// BuildToPlanInverseTable is the table name for the Plan entity.
	// It exists in this package in order to avoid circular dependency with the "plan" package.
	BuildToPlanInverseTable = "plans"
	// BuildToPlanColumn is the table column denoting the BuildToPlan relation/edge.
	BuildToPlanColumn = "plan_plan_to_build"
	// BuildToBuildCommitsTable is the table that holds the BuildToBuildCommits relation/edge.
	BuildToBuildCommitsTable = "build_commits"
	// BuildToBuildCommitsInverseTable is the table name for the BuildCommit entity.
	// It exists in this package in order to avoid circular dependency with the "buildcommit" package.
	BuildToBuildCommitsInverseTable = "build_commits"
	// BuildToBuildCommitsColumn is the table column denoting the BuildToBuildCommits relation/edge.
	BuildToBuildCommitsColumn = "build_commit_build_commit_to_build"
	// BuildToAdhocPlansTable is the table that holds the BuildToAdhocPlans relation/edge.
	BuildToAdhocPlansTable = "adhoc_plans"
	// BuildToAdhocPlansInverseTable is the table name for the AdhocPlan entity.
	// It exists in this package in order to avoid circular dependency with the "adhocplan" package.
	BuildToAdhocPlansInverseTable = "adhoc_plans"
	// BuildToAdhocPlansColumn is the table column denoting the BuildToAdhocPlans relation/edge.
	BuildToAdhocPlansColumn = "adhoc_plan_adhoc_plan_to_build"
)

// Columns holds all SQL columns for build fields.
var Columns = []string{
	FieldID,
	FieldRevision,
	FieldEnvironmentRevision,
	FieldCompletedPlan,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "builds"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"build_build_to_environment",
	"build_build_to_competition",
	"build_build_to_latest_build_commit",
}

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
	// DefaultCompletedPlan holds the default value on creation for the "completed_plan" field.
	DefaultCompletedPlan bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
