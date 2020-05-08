package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Submission holds the schema definition for the Submission entity.
type Submission struct {
	ent.Schema
}

// Mixin supplies times to the submissions.
func (Submission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimesMixin{},
		StatusMixin{},
	}
}

// Fields of the Submission.
func (Submission) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Submission.
func (Submission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("base_sha", SHA.Type).Unique().Required(),
		edge.To("head_sha", SHA.Type).Unique().Required(),
		edge.To("tasks", Task.Type),
	}
}
