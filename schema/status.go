package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

var statusValues = []string{"success", "failure", "canceled", "unstarted"}

// StatusMixin encapsulates statuses for CI jobs.
type StatusMixin struct {
	ent.Mixin
}

// Fields is the status fields for the StatusMixin.
func (StatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("completed").Default(false),
		field.Enum("state").Values(statusValues...).Default("unstarted"),
	}
}
