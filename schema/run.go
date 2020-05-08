package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Run holds the schema definition for the Run entity.
type Run struct {
	ent.Schema
}

// Mixin includes times from a central source.
func (Run) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimesMixin{},
		StatusMixin{},
	}
}

// Fields of the Run.
func (Run) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Immutable(),
		field.Bytes("metadata").Immutable(),
	}
}

// Edges of the Run.
func (Run) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).Ref("runs").Unique().Required(),
	}
}
