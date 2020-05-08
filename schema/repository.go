package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Repository holds the schema definition for the Repository entity.
type Repository struct {
	ent.Schema
}

// Fields of the Repository.
func (Repository) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owners", Owner.Type),
		edge.From("shas", SHA.Type).Ref("repository"),
	}
}

// Owner contains the administrative assignments to the repository
type Owner struct {
	ent.Schema
}

// Fields of the RepositoryOwner.
func (Owner) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the RepositoryOwner.
func (Owner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repository", Repository.Type).Ref("owners").Unique(),
	}
}
