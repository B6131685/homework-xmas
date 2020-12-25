package schema

import (

    "github.com/facebook/ent"
    "github.com/facebook/ent/schema/field"
)
// Checkout holds the schema definition for the Checkout entity.
type Checkout struct {
	ent.Schema
}

// Fields of the Checkout.
func (Checkout) Fields() []ent.Field {
	return []ent.Field{
        field.Time("Daytime"),
    }
}

// Edges of the Checkout.
func (Checkout) Edges() []ent.Edge {
	return nil
}
