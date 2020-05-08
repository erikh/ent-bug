package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Mixin includes the times grouping of CI times.
func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimesMixin{},
	}
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").Immutable(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("submission", Submission.Type).Ref("tasks").Unique().Required(),
		edge.To("runs", Run.Type),
	}
}
