// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// AdhocPlan is the client for interacting with the AdhocPlan builders.
	AdhocPlan *AdhocPlanClient
	// AgentStatus is the client for interacting with the AgentStatus builders.
	AgentStatus *AgentStatusClient
	// AgentTask is the client for interacting with the AgentTask builders.
	AgentTask *AgentTaskClient
	// AuthUser is the client for interacting with the AuthUser builders.
	AuthUser *AuthUserClient
	// Build is the client for interacting with the Build builders.
	Build *BuildClient
	// BuildCommit is the client for interacting with the BuildCommit builders.
	BuildCommit *BuildCommitClient
	// Command is the client for interacting with the Command builders.
	Command *CommandClient
	// Competition is the client for interacting with the Competition builders.
	Competition *CompetitionClient
	// DNS is the client for interacting with the DNS builders.
	DNS *DNSClient
	// DNSRecord is the client for interacting with the DNSRecord builders.
	DNSRecord *DNSRecordClient
	// Disk is the client for interacting with the Disk builders.
	Disk *DiskClient
	// Environment is the client for interacting with the Environment builders.
	Environment *EnvironmentClient
	// FileDelete is the client for interacting with the FileDelete builders.
	FileDelete *FileDeleteClient
	// FileDownload is the client for interacting with the FileDownload builders.
	FileDownload *FileDownloadClient
	// FileExtract is the client for interacting with the FileExtract builders.
	FileExtract *FileExtractClient
	// Finding is the client for interacting with the Finding builders.
	Finding *FindingClient
	// GinFileMiddleware is the client for interacting with the GinFileMiddleware builders.
	GinFileMiddleware *GinFileMiddlewareClient
	// Host is the client for interacting with the Host builders.
	Host *HostClient
	// HostDependency is the client for interacting with the HostDependency builders.
	HostDependency *HostDependencyClient
	// Identity is the client for interacting with the Identity builders.
	Identity *IdentityClient
	// IncludedNetwork is the client for interacting with the IncludedNetwork builders.
	IncludedNetwork *IncludedNetworkClient
	// Network is the client for interacting with the Network builders.
	Network *NetworkClient
	// Plan is the client for interacting with the Plan builders.
	Plan *PlanClient
	// PlanDiff is the client for interacting with the PlanDiff builders.
	PlanDiff *PlanDiffClient
	// ProvisionedHost is the client for interacting with the ProvisionedHost builders.
	ProvisionedHost *ProvisionedHostClient
	// ProvisionedNetwork is the client for interacting with the ProvisionedNetwork builders.
	ProvisionedNetwork *ProvisionedNetworkClient
	// ProvisioningStep is the client for interacting with the ProvisioningStep builders.
	ProvisioningStep *ProvisioningStepClient
	// Repository is the client for interacting with the Repository builders.
	Repository *RepositoryClient
	// Script is the client for interacting with the Script builders.
	Script *ScriptClient
	// ServerTask is the client for interacting with the ServerTask builders.
	ServerTask *ServerTaskClient
	// Status is the client for interacting with the Status builders.
	Status *StatusClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// Team is the client for interacting with the Team builders.
	Team *TeamClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
	// User is the client for interacting with the User builders.
	User *UserClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once

	// completion callbacks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook

	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Committer method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	tx.mu.Lock()
	hooks := append([]CommitHook(nil), tx.onCommit...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onCommit = append(tx.onCommit, f)
}

type (
	// Rollbacker is the interface that wraps the Rollbacker method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	tx.mu.Lock()
	hooks := append([]RollbackHook(nil), tx.onRollback...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onRollback = append(tx.onRollback, f)
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.AdhocPlan = NewAdhocPlanClient(tx.config)
	tx.AgentStatus = NewAgentStatusClient(tx.config)
	tx.AgentTask = NewAgentTaskClient(tx.config)
	tx.AuthUser = NewAuthUserClient(tx.config)
	tx.Build = NewBuildClient(tx.config)
	tx.BuildCommit = NewBuildCommitClient(tx.config)
	tx.Command = NewCommandClient(tx.config)
	tx.Competition = NewCompetitionClient(tx.config)
	tx.DNS = NewDNSClient(tx.config)
	tx.DNSRecord = NewDNSRecordClient(tx.config)
	tx.Disk = NewDiskClient(tx.config)
	tx.Environment = NewEnvironmentClient(tx.config)
	tx.FileDelete = NewFileDeleteClient(tx.config)
	tx.FileDownload = NewFileDownloadClient(tx.config)
	tx.FileExtract = NewFileExtractClient(tx.config)
	tx.Finding = NewFindingClient(tx.config)
	tx.GinFileMiddleware = NewGinFileMiddlewareClient(tx.config)
	tx.Host = NewHostClient(tx.config)
	tx.HostDependency = NewHostDependencyClient(tx.config)
	tx.Identity = NewIdentityClient(tx.config)
	tx.IncludedNetwork = NewIncludedNetworkClient(tx.config)
	tx.Network = NewNetworkClient(tx.config)
	tx.Plan = NewPlanClient(tx.config)
	tx.PlanDiff = NewPlanDiffClient(tx.config)
	tx.ProvisionedHost = NewProvisionedHostClient(tx.config)
	tx.ProvisionedNetwork = NewProvisionedNetworkClient(tx.config)
	tx.ProvisioningStep = NewProvisioningStepClient(tx.config)
	tx.Repository = NewRepositoryClient(tx.config)
	tx.Script = NewScriptClient(tx.config)
	tx.ServerTask = NewServerTaskClient(tx.config)
	tx.Status = NewStatusClient(tx.config)
	tx.Tag = NewTagClient(tx.config)
	tx.Team = NewTeamClient(tx.config)
	tx.Token = NewTokenClient(tx.config)
	tx.User = NewUserClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: AdhocPlan.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
