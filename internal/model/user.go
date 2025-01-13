package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User - structure of user
type User struct {
	ID       string `json:"id,omitempty" bson:"id"`
	Username string `json:"username" bson:"username"`
}

// UserDocument - structure of user document
type UserDocument struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
}

// UserDocumentFromUser - create new user document from user
func UserDocumentFromUser(user User) (*UserDocument, error) {
	id, err := primitive.ObjectIDFromHex(user.ID)

	if err != nil {
		return nil, err
	}

	return &UserDocument{
		ID:       id,
		Username: user.Username,
	}, nil
}

// UserDocumentFromUser - create new user document from user
func UserFromUserDocument(userDoc UserDocument) *User {

	return &User{
		ID:       userDoc.ID.Hex(),
		Username: userDoc.Username,
	}
}
