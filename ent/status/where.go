// Code generated by entc, DO NOT EDIT.

package status

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// EndedAt applies equality check predicate on the "ended_at" field. It's identical to EndedAtEQ.
func EndedAt(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndedAt), v))
	})
}

// Failed applies equality check predicate on the "failed" field. It's identical to FailedEQ.
func Failed(v bool) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFailed), v))
	})
}

// Completed applies equality check predicate on the "completed" field. It's identical to CompletedEQ.
func Completed(v bool) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompleted), v))
	})
}

// Error applies equality check predicate on the "error" field. It's identical to ErrorEQ.
func Error(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldError), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StatusForEQ applies the EQ predicate on the "status_for" field.
func StatusForEQ(v StatusFor) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusFor), v))
	})
}

// StatusForNEQ applies the NEQ predicate on the "status_for" field.
func StatusForNEQ(v StatusFor) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatusFor), v))
	})
}

// StatusForIn applies the In predicate on the "status_for" field.
func StatusForIn(vs ...StatusFor) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatusFor), v...))
	})
}

// StatusForNotIn applies the NotIn predicate on the "status_for" field.
func StatusForNotIn(vs ...StatusFor) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatusFor), v...))
	})
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartedAt), v...))
	})
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartedAt), v...))
	})
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartedAt), v))
	})
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartedAt), v))
	})
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartedAt)))
	})
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartedAt)))
	})
}

// EndedAtEQ applies the EQ predicate on the "ended_at" field.
func EndedAtEQ(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndedAt), v))
	})
}

// EndedAtNEQ applies the NEQ predicate on the "ended_at" field.
func EndedAtNEQ(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndedAt), v))
	})
}

// EndedAtIn applies the In predicate on the "ended_at" field.
func EndedAtIn(vs ...time.Time) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndedAt), v...))
	})
}

// EndedAtNotIn applies the NotIn predicate on the "ended_at" field.
func EndedAtNotIn(vs ...time.Time) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndedAt), v...))
	})
}

// EndedAtGT applies the GT predicate on the "ended_at" field.
func EndedAtGT(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndedAt), v))
	})
}

// EndedAtGTE applies the GTE predicate on the "ended_at" field.
func EndedAtGTE(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndedAt), v))
	})
}

// EndedAtLT applies the LT predicate on the "ended_at" field.
func EndedAtLT(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndedAt), v))
	})
}

// EndedAtLTE applies the LTE predicate on the "ended_at" field.
func EndedAtLTE(v time.Time) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndedAt), v))
	})
}

// EndedAtIsNil applies the IsNil predicate on the "ended_at" field.
func EndedAtIsNil() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndedAt)))
	})
}

// EndedAtNotNil applies the NotNil predicate on the "ended_at" field.
func EndedAtNotNil() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndedAt)))
	})
}

// FailedEQ applies the EQ predicate on the "failed" field.
func FailedEQ(v bool) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFailed), v))
	})
}

// FailedNEQ applies the NEQ predicate on the "failed" field.
func FailedNEQ(v bool) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFailed), v))
	})
}

// CompletedEQ applies the EQ predicate on the "completed" field.
func CompletedEQ(v bool) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompleted), v))
	})
}

// CompletedNEQ applies the NEQ predicate on the "completed" field.
func CompletedNEQ(v bool) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompleted), v))
	})
}

// ErrorEQ applies the EQ predicate on the "error" field.
func ErrorEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldError), v))
	})
}

// ErrorNEQ applies the NEQ predicate on the "error" field.
func ErrorNEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldError), v))
	})
}

// ErrorIn applies the In predicate on the "error" field.
func ErrorIn(vs ...string) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldError), v...))
	})
}

// ErrorNotIn applies the NotIn predicate on the "error" field.
func ErrorNotIn(vs ...string) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldError), v...))
	})
}

// ErrorGT applies the GT predicate on the "error" field.
func ErrorGT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldError), v))
	})
}

// ErrorGTE applies the GTE predicate on the "error" field.
func ErrorGTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldError), v))
	})
}

// ErrorLT applies the LT predicate on the "error" field.
func ErrorLT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldError), v))
	})
}

// ErrorLTE applies the LTE predicate on the "error" field.
func ErrorLTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldError), v))
	})
}

// ErrorContains applies the Contains predicate on the "error" field.
func ErrorContains(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldError), v))
	})
}

// ErrorHasPrefix applies the HasPrefix predicate on the "error" field.
func ErrorHasPrefix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldError), v))
	})
}

// ErrorHasSuffix applies the HasSuffix predicate on the "error" field.
func ErrorHasSuffix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldError), v))
	})
}

// ErrorIsNil applies the IsNil predicate on the "error" field.
func ErrorIsNil() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldError)))
	})
}

// ErrorNotNil applies the NotNil predicate on the "error" field.
func ErrorNotNil() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldError)))
	})
}

// ErrorEqualFold applies the EqualFold predicate on the "error" field.
func ErrorEqualFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldError), v))
	})
}

// ErrorContainsFold applies the ContainsFold predicate on the "error" field.
func ErrorContainsFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldError), v))
	})
}

// HasStatusToBuild applies the HasEdge predicate on the "StatusToBuild" edge.
func HasStatusToBuild() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToBuildTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToBuildTable, StatusToBuildColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusToBuildWith applies the HasEdge predicate on the "StatusToBuild" edge with a given conditions (other predicates).
func HasStatusToBuildWith(preds ...predicate.Build) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToBuildInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToBuildTable, StatusToBuildColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusToProvisionedNetwork applies the HasEdge predicate on the "StatusToProvisionedNetwork" edge.
func HasStatusToProvisionedNetwork() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToProvisionedNetworkTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToProvisionedNetworkTable, StatusToProvisionedNetworkColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusToProvisionedNetworkWith applies the HasEdge predicate on the "StatusToProvisionedNetwork" edge with a given conditions (other predicates).
func HasStatusToProvisionedNetworkWith(preds ...predicate.ProvisionedNetwork) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToProvisionedNetworkInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToProvisionedNetworkTable, StatusToProvisionedNetworkColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusToProvisionedHost applies the HasEdge predicate on the "StatusToProvisionedHost" edge.
func HasStatusToProvisionedHost() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToProvisionedHostTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToProvisionedHostTable, StatusToProvisionedHostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusToProvisionedHostWith applies the HasEdge predicate on the "StatusToProvisionedHost" edge with a given conditions (other predicates).
func HasStatusToProvisionedHostWith(preds ...predicate.ProvisionedHost) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToProvisionedHostInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToProvisionedHostTable, StatusToProvisionedHostColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusToProvisioningStep applies the HasEdge predicate on the "StatusToProvisioningStep" edge.
func HasStatusToProvisioningStep() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToProvisioningStepTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToProvisioningStepTable, StatusToProvisioningStepColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusToProvisioningStepWith applies the HasEdge predicate on the "StatusToProvisioningStep" edge with a given conditions (other predicates).
func HasStatusToProvisioningStepWith(preds ...predicate.ProvisioningStep) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToProvisioningStepInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToProvisioningStepTable, StatusToProvisioningStepColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusToTeam applies the HasEdge predicate on the "StatusToTeam" edge.
func HasStatusToTeam() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToTeamTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToTeamTable, StatusToTeamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusToTeamWith applies the HasEdge predicate on the "StatusToTeam" edge with a given conditions (other predicates).
func HasStatusToTeamWith(preds ...predicate.Team) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToTeamInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToTeamTable, StatusToTeamColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusToPlan applies the HasEdge predicate on the "StatusToPlan" edge.
func HasStatusToPlan() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToPlanTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToPlanTable, StatusToPlanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusToPlanWith applies the HasEdge predicate on the "StatusToPlan" edge with a given conditions (other predicates).
func HasStatusToPlanWith(preds ...predicate.Plan) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToPlanInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToPlanTable, StatusToPlanColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusToServerTask applies the HasEdge predicate on the "StatusToServerTask" edge.
func HasStatusToServerTask() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToServerTaskTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToServerTaskTable, StatusToServerTaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusToServerTaskWith applies the HasEdge predicate on the "StatusToServerTask" edge with a given conditions (other predicates).
func HasStatusToServerTaskWith(preds ...predicate.ServerTask) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusToServerTaskInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, StatusToServerTaskTable, StatusToServerTaskColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		p(s.Not())
	})
}
