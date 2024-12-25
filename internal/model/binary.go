package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Binary - structure of binary type of data
type Binary struct {
	ID       string `json:"id,omitempty" bson:"id,omitempty"`
	UserID   string `json:"user_id,omitempty" bson:"user_id"`
	Title    string `json:"title" bson:"title"`
	MimeType string `json:"mime_type" bson:"mime_type"`
	Data     []byte `json:"data" bson:"data"`
	Meta     string `json:"meta" bson:"meta"`
}

// BinaryDocument - structure of binary type of data document
type BinaryDocument struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserID   string             `bson:"user_id"`
	Title    string             `bson:"title"`
	MimeType string             `bson:"mime_type"`
	Data     []byte             `bson:"data"`
	Meta     string             `bson:"meta"`
}

// BinaryDocumentFromBinary - create new binary document from binary
func BinaryDocumentFromBinary(binary Binary) (*BinaryDocument, error) {
	id, err := primitive.ObjectIDFromHex(binary.ID)

	if err != nil {
		return nil, err
	}

	return &BinaryDocument{
		ID:       id,
		UserID:   binary.UserID,
		Title:    binary.Title,
		MimeType: binary.MimeType,
		Data:     binary.Data,
		Meta:     binary.Meta,
	}, nil
}

// BinaryFromBinaryDocument - create new binary from binary document
func BinaryFromBinaryDocument(binaryDoc BinaryDocument) *Binary {

	return &Binary{
		ID:       binaryDoc.ID.Hex(),
		UserID:   binaryDoc.UserID,
		Title:    binaryDoc.Title,
		MimeType: binaryDoc.MimeType,
		Data:     binaryDoc.Data,
		Meta:     binaryDoc.Meta,
	}
}
