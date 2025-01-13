package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// BankCard - structure of bank card
type BankCard struct {
	ID     string `json:"id,omitempty" bson:"id,omitempty"`
	UserID string `json:"user_id,omitempty" bson:"user_id"`
	Title  string `json:"title" bson:"title"`
	Data   int    `json:"data" bson:"data"`
	Meta   string `json:"meta" bson:"meta"`
}

// BankCardDocument - structure of bank card in repo
type BankCardDocument struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID string             `bson:"user_id"`
	Title  string             `bson:"title"`
	Data   int                `bson:"data"`
	Meta   string             `bson:"meta"`
}

// BankCardDocumentFromBankCard - create new bank card document from bank card
func BankCardDocumentFromBankCard(card BankCard) (*BankCardDocument, error) {
	id, err := primitive.ObjectIDFromHex(card.ID)

	if err != nil {
		return nil, err
	}

	return &BankCardDocument{
		ID:     id,
		UserID: card.UserID,
		Title:  card.Title,
		Data:   card.Data,
		Meta:   card.Meta,
	}, nil
}

// BankCardDocumentFromBankCard - create new bank card document from bank card
func BankCardFromBankCardDocument(cardDoc BankCardDocument) *BankCard {

	return &BankCard{
		ID:     cardDoc.ID.Hex(),
		UserID: cardDoc.UserID,
		Title:  cardDoc.Title,
		Data:   cardDoc.Data,
		Meta:   cardDoc.Meta,
	}
}
