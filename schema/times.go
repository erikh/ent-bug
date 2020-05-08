package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// TimesMixin is a mixin for CI-friendly time specifiers for run timing, etc.
type TimesMixin struct {
	ent.Mixin
}

// Fields for the TimesMixin.
func (TimesMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("started_at").Nillable().Optional(),
		field.Time("finished_at").Nillable().Optional(),
	}
}
