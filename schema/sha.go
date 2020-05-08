package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// SHA holds the schema definition for the SHA entity.
type SHA struct {
	ent.Schema
}

// Fields of the Ref.
func (SHA) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the SHA.
func (SHA) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("branches", Branch.Type),
		edge.To("repository", Repository.Type).Unique(),
		edge.From("submissions", Submission.Type).Ref("head_sha"),
		edge.From("submissions_base", Submission.Type).Ref("base_sha"),
	}
}

// Branch is a branch associated with many shas.
type Branch struct {
	ent.Schema
}

// Fields is mostly the name here.
func (Branch) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Immutable(),
	}
}

// Edges denotes the edges between the branch and sha.
func (Branch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sha", SHA.Type).Ref("branches"),
	}
}
