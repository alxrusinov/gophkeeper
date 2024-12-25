package mongo

import (
	"context"
	"fmt"

	"github.com/alxrusinov/gophkeeper/internal/customerrors"
	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetNote - returns note for user by note id
func (m *Mongo) GetNote(ctx context.Context, userID, noteID string) (*model.Note, error) {
	result := new(model.NoteDocument)

	noteIDFilter, err := primitive.ObjectIDFromHex(noteID)

	if err != nil {
		return nil, &customerrors.NotFound{Err: fmt.Errorf("note with id %s does not exist", noteID)}
	}

	filter := bson.D{{Key: "user_id", Value: userID}, {Key: "_id", Value: noteIDFilter}}

	err = m.client.Database(DataBase).Collection(NoteCollection).FindOne(ctx, filter, options.FindOne().SetProjection(bson.M{"user_id": 0})).Decode(result)

	if err != nil {
		return nil, &customerrors.NotFound{Err: fmt.Errorf("note with id %s does not exist", noteID)}
	}

	note := model.NoteFromNoteDocument(*result)

	return note, nil
}
