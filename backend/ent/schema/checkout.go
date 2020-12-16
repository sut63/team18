package schema

import "github.com/facebookincubator/ent"

// Checkout holds the schema definition for the Checkout entity.
type Checkout struct {
	ent.Schema
}

// Fields of the Checkout.
func (Checkout) Fields() []ent.Field {
	return nil
}

// Edges of the Checkout.
func (Checkout) Edges() []ent.Edge {
	return nil
}
