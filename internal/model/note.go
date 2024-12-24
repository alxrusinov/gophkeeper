package model

type Note struct {
	ID     string `json:"id,omitempty" bson:"id,omitempty"`
	UserID string `json:"user_id,omitempty" bson:"user_id"`
	Title  string `json:"title" bson:"title"`
	Data   string `json:"data" bson:"data"`
	Meta   string `json:"meta" bson:"meta"`
}
