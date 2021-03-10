// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/provisionedhost"
)

// ProvisionedHost is the model entity for the ProvisionedHost schema.
type ProvisionedHost struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SubnetIP holds the value of the "subnet_ip" field.
	SubnetIP string `json:"subnet_ip,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProvisionedHostQuery when eager-loading is set.
	Edges ProvisionedHostEdges `json:"edges"`

	// Edges put into the main struct to be loaded via hcl
	// ProvisionedHostToTag holds the value of the ProvisionedHostToTag edge.
	HCLProvisionedHostToTag []*Tag `json:"ProvisionedHostToTag,omitempty"`
	// ProvisionedHostToStatus holds the value of the ProvisionedHostToStatus edge.
	HCLProvisionedHostToStatus []*Status `json:"ProvisionedHostToStatus,omitempty"`
	// ProvisionedHostToProvisionedNetwork holds the value of the ProvisionedHostToProvisionedNetwork edge.
	HCLProvisionedHostToProvisionedNetwork []*ProvisionedNetwork `json:"ProvisionedHostToProvisionedNetwork,omitempty"`
	// ProvisionedHostToHost holds the value of the ProvisionedHostToHost edge.
	HCLProvisionedHostToHost []*Host `json:"ProvisionedHostToHost,omitempty"`
	// ProvisionedHostToProvisioningStep holds the value of the ProvisionedHostToProvisioningStep edge.
	HCLProvisionedHostToProvisioningStep []*ProvisioningStep `json:"ProvisionedHostToProvisioningStep,omitempty"`
	// ProvisionedHostToAgentStatus holds the value of the ProvisionedHostToAgentStatus edge.
	HCLProvisionedHostToAgentStatus []*AgentStatus `json:"ProvisionedHostToAgentStatus,omitempty"`
	//

}

// ProvisionedHostEdges holds the relations/edges for other nodes in the graph.
type ProvisionedHostEdges struct {
	// ProvisionedHostToTag holds the value of the ProvisionedHostToTag edge.
	ProvisionedHostToTag []*Tag `json:"ProvisionedHostToTag,omitempty"`
	// ProvisionedHostToStatus holds the value of the ProvisionedHostToStatus edge.
	ProvisionedHostToStatus []*Status `json:"ProvisionedHostToStatus,omitempty"`
	// ProvisionedHostToProvisionedNetwork holds the value of the ProvisionedHostToProvisionedNetwork edge.
	ProvisionedHostToProvisionedNetwork []*ProvisionedNetwork `json:"ProvisionedHostToProvisionedNetwork,omitempty"`
	// ProvisionedHostToHost holds the value of the ProvisionedHostToHost edge.
	ProvisionedHostToHost []*Host `json:"ProvisionedHostToHost,omitempty"`
	// ProvisionedHostToProvisioningStep holds the value of the ProvisionedHostToProvisioningStep edge.
	ProvisionedHostToProvisioningStep []*ProvisioningStep `json:"ProvisionedHostToProvisioningStep,omitempty"`
	// ProvisionedHostToAgentStatus holds the value of the ProvisionedHostToAgentStatus edge.
	ProvisionedHostToAgentStatus []*AgentStatus `json:"ProvisionedHostToAgentStatus,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// ProvisionedHostToTagOrErr returns the ProvisionedHostToTag value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedHostEdges) ProvisionedHostToTagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.ProvisionedHostToTag, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToTag"}
}

// ProvisionedHostToStatusOrErr returns the ProvisionedHostToStatus value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedHostEdges) ProvisionedHostToStatusOrErr() ([]*Status, error) {
	if e.loadedTypes[1] {
		return e.ProvisionedHostToStatus, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToStatus"}
}

// ProvisionedHostToProvisionedNetworkOrErr returns the ProvisionedHostToProvisionedNetwork value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedHostEdges) ProvisionedHostToProvisionedNetworkOrErr() ([]*ProvisionedNetwork, error) {
	if e.loadedTypes[2] {
		return e.ProvisionedHostToProvisionedNetwork, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToProvisionedNetwork"}
}

// ProvisionedHostToHostOrErr returns the ProvisionedHostToHost value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedHostEdges) ProvisionedHostToHostOrErr() ([]*Host, error) {
	if e.loadedTypes[3] {
		return e.ProvisionedHostToHost, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToHost"}
}

// ProvisionedHostToProvisioningStepOrErr returns the ProvisionedHostToProvisioningStep value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedHostEdges) ProvisionedHostToProvisioningStepOrErr() ([]*ProvisioningStep, error) {
	if e.loadedTypes[4] {
		return e.ProvisionedHostToProvisioningStep, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToProvisioningStep"}
}

// ProvisionedHostToAgentStatusOrErr returns the ProvisionedHostToAgentStatus value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedHostEdges) ProvisionedHostToAgentStatusOrErr() ([]*AgentStatus, error) {
	if e.loadedTypes[5] {
		return e.ProvisionedHostToAgentStatus, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToAgentStatus"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProvisionedHost) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case provisionedhost.FieldID:
			values[i] = &sql.NullInt64{}
		case provisionedhost.FieldSubnetIP:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProvisionedHost", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProvisionedHost fields.
func (ph *ProvisionedHost) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provisionedhost.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ph.ID = int(value.Int64)
		case provisionedhost.FieldSubnetIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subnet_ip", values[i])
			} else if value.Valid {
				ph.SubnetIP = value.String
			}
		}
	}
	return nil
}

// QueryProvisionedHostToTag queries the "ProvisionedHostToTag" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToTag() *TagQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToTag(ph)
}

// QueryProvisionedHostToStatus queries the "ProvisionedHostToStatus" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToStatus() *StatusQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToStatus(ph)
}

// QueryProvisionedHostToProvisionedNetwork queries the "ProvisionedHostToProvisionedNetwork" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToProvisionedNetwork() *ProvisionedNetworkQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToProvisionedNetwork(ph)
}

// QueryProvisionedHostToHost queries the "ProvisionedHostToHost" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToHost() *HostQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToHost(ph)
}

// QueryProvisionedHostToProvisioningStep queries the "ProvisionedHostToProvisioningStep" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToProvisioningStep() *ProvisioningStepQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToProvisioningStep(ph)
}

// QueryProvisionedHostToAgentStatus queries the "ProvisionedHostToAgentStatus" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToAgentStatus() *AgentStatusQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToAgentStatus(ph)
}

// Update returns a builder for updating this ProvisionedHost.
// Note that you need to call ProvisionedHost.Unwrap() before calling this method if this ProvisionedHost
// was returned from a transaction, and the transaction was committed or rolled back.
func (ph *ProvisionedHost) Update() *ProvisionedHostUpdateOne {
	return (&ProvisionedHostClient{config: ph.config}).UpdateOne(ph)
}

// Unwrap unwraps the ProvisionedHost entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ph *ProvisionedHost) Unwrap() *ProvisionedHost {
	tx, ok := ph.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProvisionedHost is not a transactional entity")
	}
	ph.config.driver = tx.drv
	return ph
}

// String implements the fmt.Stringer.
func (ph *ProvisionedHost) String() string {
	var builder strings.Builder
	builder.WriteString("ProvisionedHost(")
	builder.WriteString(fmt.Sprintf("id=%v", ph.ID))
	builder.WriteString(", subnet_ip=")
	builder.WriteString(ph.SubnetIP)
	builder.WriteByte(')')
	return builder.String()
}

// ProvisionedHosts is a parsable slice of ProvisionedHost.
type ProvisionedHosts []*ProvisionedHost

func (ph ProvisionedHosts) config(cfg config) {
	for _i := range ph {
		ph[_i].config = cfg
	}
}
