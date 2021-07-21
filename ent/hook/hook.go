// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/gen0cide/laforge/ent"
)

// The AgentStatusFunc type is an adapter to allow the use of ordinary
// function as AgentStatus mutator.
type AgentStatusFunc func(context.Context, *ent.AgentStatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AgentStatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AgentStatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AgentStatusMutation", m)
	}
	return f(ctx, mv)
}

// The AgentTaskFunc type is an adapter to allow the use of ordinary
// function as AgentTask mutator.
type AgentTaskFunc func(context.Context, *ent.AgentTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AgentTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AgentTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AgentTaskMutation", m)
	}
	return f(ctx, mv)
}

// The AuthUserFunc type is an adapter to allow the use of ordinary
// function as AuthUser mutator.
type AuthUserFunc func(context.Context, *ent.AuthUserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AuthUserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AuthUserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AuthUserMutation", m)
	}
	return f(ctx, mv)
}

// The BuildFunc type is an adapter to allow the use of ordinary
// function as Build mutator.
type BuildFunc func(context.Context, *ent.BuildMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BuildFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BuildMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BuildMutation", m)
	}
	return f(ctx, mv)
}

// The CommandFunc type is an adapter to allow the use of ordinary
// function as Command mutator.
type CommandFunc func(context.Context, *ent.CommandMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CommandFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CommandMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CommandMutation", m)
	}
	return f(ctx, mv)
}

// The CompetitionFunc type is an adapter to allow the use of ordinary
// function as Competition mutator.
type CompetitionFunc func(context.Context, *ent.CompetitionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompetitionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CompetitionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompetitionMutation", m)
	}
	return f(ctx, mv)
}

// The DNSFunc type is an adapter to allow the use of ordinary
// function as DNS mutator.
type DNSFunc func(context.Context, *ent.DNSMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DNSFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DNSMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DNSMutation", m)
	}
	return f(ctx, mv)
}

// The DNSRecordFunc type is an adapter to allow the use of ordinary
// function as DNSRecord mutator.
type DNSRecordFunc func(context.Context, *ent.DNSRecordMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DNSRecordFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DNSRecordMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DNSRecordMutation", m)
	}
	return f(ctx, mv)
}

// The DiskFunc type is an adapter to allow the use of ordinary
// function as Disk mutator.
type DiskFunc func(context.Context, *ent.DiskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DiskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DiskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DiskMutation", m)
	}
	return f(ctx, mv)
}

// The EnvironmentFunc type is an adapter to allow the use of ordinary
// function as Environment mutator.
type EnvironmentFunc func(context.Context, *ent.EnvironmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EnvironmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EnvironmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EnvironmentMutation", m)
	}
	return f(ctx, mv)
}

// The FileDeleteFunc type is an adapter to allow the use of ordinary
// function as FileDelete mutator.
type FileDeleteFunc func(context.Context, *ent.FileDeleteMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FileDeleteFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FileDeleteMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FileDeleteMutation", m)
	}
	return f(ctx, mv)
}

// The FileDownloadFunc type is an adapter to allow the use of ordinary
// function as FileDownload mutator.
type FileDownloadFunc func(context.Context, *ent.FileDownloadMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FileDownloadFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FileDownloadMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FileDownloadMutation", m)
	}
	return f(ctx, mv)
}

// The FileExtractFunc type is an adapter to allow the use of ordinary
// function as FileExtract mutator.
type FileExtractFunc func(context.Context, *ent.FileExtractMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FileExtractFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FileExtractMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FileExtractMutation", m)
	}
	return f(ctx, mv)
}

// The FindingFunc type is an adapter to allow the use of ordinary
// function as Finding mutator.
type FindingFunc func(context.Context, *ent.FindingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FindingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FindingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FindingMutation", m)
	}
	return f(ctx, mv)
}

// The GinFileMiddlewareFunc type is an adapter to allow the use of ordinary
// function as GinFileMiddleware mutator.
type GinFileMiddlewareFunc func(context.Context, *ent.GinFileMiddlewareMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GinFileMiddlewareFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GinFileMiddlewareMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GinFileMiddlewareMutation", m)
	}
	return f(ctx, mv)
}

// The HostFunc type is an adapter to allow the use of ordinary
// function as Host mutator.
type HostFunc func(context.Context, *ent.HostMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HostFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.HostMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HostMutation", m)
	}
	return f(ctx, mv)
}

// The HostDependencyFunc type is an adapter to allow the use of ordinary
// function as HostDependency mutator.
type HostDependencyFunc func(context.Context, *ent.HostDependencyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HostDependencyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.HostDependencyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HostDependencyMutation", m)
	}
	return f(ctx, mv)
}

// The IdentityFunc type is an adapter to allow the use of ordinary
// function as Identity mutator.
type IdentityFunc func(context.Context, *ent.IdentityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f IdentityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.IdentityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.IdentityMutation", m)
	}
	return f(ctx, mv)
}

// The IncludedNetworkFunc type is an adapter to allow the use of ordinary
// function as IncludedNetwork mutator.
type IncludedNetworkFunc func(context.Context, *ent.IncludedNetworkMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f IncludedNetworkFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.IncludedNetworkMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.IncludedNetworkMutation", m)
	}
	return f(ctx, mv)
}

// The NetworkFunc type is an adapter to allow the use of ordinary
// function as Network mutator.
type NetworkFunc func(context.Context, *ent.NetworkMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NetworkFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.NetworkMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NetworkMutation", m)
	}
	return f(ctx, mv)
}

// The PlanFunc type is an adapter to allow the use of ordinary
// function as Plan mutator.
type PlanFunc func(context.Context, *ent.PlanMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlanFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PlanMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlanMutation", m)
	}
	return f(ctx, mv)
}

// The ProvisionedHostFunc type is an adapter to allow the use of ordinary
// function as ProvisionedHost mutator.
type ProvisionedHostFunc func(context.Context, *ent.ProvisionedHostMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProvisionedHostFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProvisionedHostMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProvisionedHostMutation", m)
	}
	return f(ctx, mv)
}

// The ProvisionedNetworkFunc type is an adapter to allow the use of ordinary
// function as ProvisionedNetwork mutator.
type ProvisionedNetworkFunc func(context.Context, *ent.ProvisionedNetworkMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProvisionedNetworkFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProvisionedNetworkMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProvisionedNetworkMutation", m)
	}
	return f(ctx, mv)
}

// The ProvisioningStepFunc type is an adapter to allow the use of ordinary
// function as ProvisioningStep mutator.
type ProvisioningStepFunc func(context.Context, *ent.ProvisioningStepMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProvisioningStepFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProvisioningStepMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProvisioningStepMutation", m)
	}
	return f(ctx, mv)
}

// The ScriptFunc type is an adapter to allow the use of ordinary
// function as Script mutator.
type ScriptFunc func(context.Context, *ent.ScriptMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScriptFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ScriptMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScriptMutation", m)
	}
	return f(ctx, mv)
}

// The ServerTaskFunc type is an adapter to allow the use of ordinary
// function as ServerTask mutator.
type ServerTaskFunc func(context.Context, *ent.ServerTaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ServerTaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ServerTaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ServerTaskMutation", m)
	}
	return f(ctx, mv)
}

// The StatusFunc type is an adapter to allow the use of ordinary
// function as Status mutator.
type StatusFunc func(context.Context, *ent.StatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatusMutation", m)
	}
	return f(ctx, mv)
}

// The TagFunc type is an adapter to allow the use of ordinary
// function as Tag mutator.
type TagFunc func(context.Context, *ent.TagMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TagFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TagMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TagMutation", m)
	}
	return f(ctx, mv)
}

// The TeamFunc type is an adapter to allow the use of ordinary
// function as Team mutator.
type TeamFunc func(context.Context, *ent.TeamMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TeamFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TeamMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TeamMutation", m)
	}
	return f(ctx, mv)
}

// The TokenFunc type is an adapter to allow the use of ordinary
// function as Token mutator.
type TokenFunc func(context.Context, *ent.TokenMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TokenFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TokenMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TokenMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
