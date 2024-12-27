package model

// SourceID - struct of id source
type SourceID struct {
	ID     string `json:"id" bson:"_id"`
	UserID string `json:"user_id,omitempty" bson:"user_id,omitempty"`
}
