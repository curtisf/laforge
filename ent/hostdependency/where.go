// Code generated by entc, DO NOT EDIT.

package hostdependency

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
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
func IDGT(id int) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HostID applies equality check predicate on the "host_id" field. It's identical to HostIDEQ.
func HostID(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostID), v))
	})
}

// NetworkID applies equality check predicate on the "network_id" field. It's identical to NetworkIDEQ.
func NetworkID(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNetworkID), v))
	})
}

// HostIDEQ applies the EQ predicate on the "host_id" field.
func HostIDEQ(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostID), v))
	})
}

// HostIDNEQ applies the NEQ predicate on the "host_id" field.
func HostIDNEQ(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHostID), v))
	})
}

// HostIDIn applies the In predicate on the "host_id" field.
func HostIDIn(vs ...string) predicate.HostDependency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HostDependency(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHostID), v...))
	})
}

// HostIDNotIn applies the NotIn predicate on the "host_id" field.
func HostIDNotIn(vs ...string) predicate.HostDependency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HostDependency(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHostID), v...))
	})
}

// HostIDGT applies the GT predicate on the "host_id" field.
func HostIDGT(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHostID), v))
	})
}

// HostIDGTE applies the GTE predicate on the "host_id" field.
func HostIDGTE(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHostID), v))
	})
}

// HostIDLT applies the LT predicate on the "host_id" field.
func HostIDLT(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHostID), v))
	})
}

// HostIDLTE applies the LTE predicate on the "host_id" field.
func HostIDLTE(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHostID), v))
	})
}

// HostIDContains applies the Contains predicate on the "host_id" field.
func HostIDContains(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHostID), v))
	})
}

// HostIDHasPrefix applies the HasPrefix predicate on the "host_id" field.
func HostIDHasPrefix(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHostID), v))
	})
}

// HostIDHasSuffix applies the HasSuffix predicate on the "host_id" field.
func HostIDHasSuffix(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHostID), v))
	})
}

// HostIDEqualFold applies the EqualFold predicate on the "host_id" field.
func HostIDEqualFold(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHostID), v))
	})
}

// HostIDContainsFold applies the ContainsFold predicate on the "host_id" field.
func HostIDContainsFold(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHostID), v))
	})
}

// NetworkIDEQ applies the EQ predicate on the "network_id" field.
func NetworkIDEQ(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNetworkID), v))
	})
}

// NetworkIDNEQ applies the NEQ predicate on the "network_id" field.
func NetworkIDNEQ(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNetworkID), v))
	})
}

// NetworkIDIn applies the In predicate on the "network_id" field.
func NetworkIDIn(vs ...string) predicate.HostDependency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HostDependency(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNetworkID), v...))
	})
}

// NetworkIDNotIn applies the NotIn predicate on the "network_id" field.
func NetworkIDNotIn(vs ...string) predicate.HostDependency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HostDependency(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNetworkID), v...))
	})
}

// NetworkIDGT applies the GT predicate on the "network_id" field.
func NetworkIDGT(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNetworkID), v))
	})
}

// NetworkIDGTE applies the GTE predicate on the "network_id" field.
func NetworkIDGTE(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNetworkID), v))
	})
}

// NetworkIDLT applies the LT predicate on the "network_id" field.
func NetworkIDLT(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNetworkID), v))
	})
}

// NetworkIDLTE applies the LTE predicate on the "network_id" field.
func NetworkIDLTE(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNetworkID), v))
	})
}

// NetworkIDContains applies the Contains predicate on the "network_id" field.
func NetworkIDContains(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNetworkID), v))
	})
}

// NetworkIDHasPrefix applies the HasPrefix predicate on the "network_id" field.
func NetworkIDHasPrefix(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNetworkID), v))
	})
}

// NetworkIDHasSuffix applies the HasSuffix predicate on the "network_id" field.
func NetworkIDHasSuffix(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNetworkID), v))
	})
}

// NetworkIDEqualFold applies the EqualFold predicate on the "network_id" field.
func NetworkIDEqualFold(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNetworkID), v))
	})
}

// NetworkIDContainsFold applies the ContainsFold predicate on the "network_id" field.
func NetworkIDContainsFold(v string) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNetworkID), v))
	})
}

// HasHostDependencyToDependOnHost applies the HasEdge predicate on the "HostDependencyToDependOnHost" edge.
func HasHostDependencyToDependOnHost() predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostDependencyToDependOnHostTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HostDependencyToDependOnHostTable, HostDependencyToDependOnHostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostDependencyToDependOnHostWith applies the HasEdge predicate on the "HostDependencyToDependOnHost" edge with a given conditions (other predicates).
func HasHostDependencyToDependOnHostWith(preds ...predicate.Host) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostDependencyToDependOnHostInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HostDependencyToDependOnHostTable, HostDependencyToDependOnHostColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHostDependencyToDependByHost applies the HasEdge predicate on the "HostDependencyToDependByHost" edge.
func HasHostDependencyToDependByHost() predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostDependencyToDependByHostTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HostDependencyToDependByHostTable, HostDependencyToDependByHostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostDependencyToDependByHostWith applies the HasEdge predicate on the "HostDependencyToDependByHost" edge with a given conditions (other predicates).
func HasHostDependencyToDependByHostWith(preds ...predicate.Host) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostDependencyToDependByHostInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HostDependencyToDependByHostTable, HostDependencyToDependByHostColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHostDependencyToNetwork applies the HasEdge predicate on the "HostDependencyToNetwork" edge.
func HasHostDependencyToNetwork() predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostDependencyToNetworkTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HostDependencyToNetworkTable, HostDependencyToNetworkColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostDependencyToNetworkWith applies the HasEdge predicate on the "HostDependencyToNetwork" edge with a given conditions (other predicates).
func HasHostDependencyToNetworkWith(preds ...predicate.Network) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostDependencyToNetworkInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HostDependencyToNetworkTable, HostDependencyToNetworkColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHostDependencyToEnvironment applies the HasEdge predicate on the "HostDependencyToEnvironment" edge.
func HasHostDependencyToEnvironment() predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostDependencyToEnvironmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HostDependencyToEnvironmentTable, HostDependencyToEnvironmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostDependencyToEnvironmentWith applies the HasEdge predicate on the "HostDependencyToEnvironment" edge with a given conditions (other predicates).
func HasHostDependencyToEnvironmentWith(preds ...predicate.Environment) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostDependencyToEnvironmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HostDependencyToEnvironmentTable, HostDependencyToEnvironmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HostDependency) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HostDependency) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
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
func Not(p predicate.HostDependency) predicate.HostDependency {
	return predicate.HostDependency(func(s *sql.Selector) {
		p(s.Not())
	})
}
