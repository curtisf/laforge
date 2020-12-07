package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// ProvisioningStep holds the schema definition for the ProvisioningStep entity.
type ProvisioningStep struct {
	ent.Schema
}

// Fields of the ProvisioningStep.
func (ProvisioningStep) Fields() []ent.Field {
	return []ent.Field{
		field.String("provisioner_type"),
		field.Int("step_number"),
	}
}

// Edges of the ProvisioningStep.
func (ProvisioningStep) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provisioned_host", ProvisionedHost.Type),
		edge.To("status", Status.Type),
		edge.To("script", Script.Type),
		edge.To("command", Command.Type),
		edge.To("dns_record", DNSRecord.Type),
		edge.To("file_download", FileDownload.Type),
		edge.To("file_delete", FileDelete.Type),
		edge.To("file_extract", FileExtract.Type),
		edge.To("tag", Tag.Type),
	}
}
