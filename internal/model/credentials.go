package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Credentials - structure of creds
type Credentials struct {
	ID     string `json:"id,omitempty" bson:"id,omitempty"`
	UserID string `json:"user_id,omitempty" bson:"user_id"`
	Title  string `json:"title" bson:"title"`
	Data   Login  `json:"data" bson:"data"`
	Meta   string `json:"meta" bson:"meta"`
}

// CredentialsDocument - structure of creds document
type CredentialsDocument struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID string             `bson:"user_id"`
	Title  string             `bson:"title"`
	Data   Login              `bson:"data"`
	Meta   string             `bson:"meta"`
}

// CredentialsDocumentFromCredentials - create
// new credentials document from credentials
func CredentialsDocumentFromCredentials(cred Credentials) (*CredentialsDocument, error) {
	id, err := primitive.ObjectIDFromHex(cred.ID)

	if err != nil {
		return nil, err
	}

	return &CredentialsDocument{
		ID:     id,
		UserID: cred.UserID,
		Title:  cred.Title,
		Data:   cred.Data,
		Meta:   cred.Meta,
	}, nil
}

// CredentialsFromCredentialsDocument - create
// new credentials from credentials document
func CredentialsFromCredentialsDocument(credDoc CredentialsDocument) *Credentials {

	return &Credentials{
		ID:     credDoc.ID.Hex(),
		UserID: credDoc.UserID,
		Title:  credDoc.Title,
		Data:   credDoc.Data,
		Meta:   credDoc.Meta,
	}
}
