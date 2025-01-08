package mongo

import (
	"context"
	"errors"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

// AddBinary - adds new binaey data for user
func (m *Mongo) AddBinary(ctx context.Context, data *model.BinaryUpload) (*model.Binary, error) {

	session, err := m.client.StartSession()
	if err != nil {
		return nil, err
	}

	defer session.EndSession(context.Background())

	err = session.StartTransaction()

	if err != nil {
		return nil, err
	}

	bucket, err := gridfs.NewBucket(m.client.Database(DataBase))

	if err != nil {
		return nil, err
	}

	fileID, err := bucket.UploadFromStream(data.Title, data.Data)

	if err != nil {
		return nil, err
	}

	sendData := &model.Binary{
		UserID:   data.UserID,
		Title:    data.Title,
		Meta:     data.Meta,
		MimeType: data.MimeType,
		FileID:   fileID.Hex(),
	}

	insertedResult, err := m.client.Database(DataBase).Collection(BinaryCollection).InsertOne(ctx, sendData)

	if err != nil {
		return nil, err
	}

	err = session.CommitTransaction(context.Background())

	if err != nil {
		return nil, err
	}

	if id, ok := insertedResult.InsertedID.(primitive.ObjectID); ok {
		sendData.ID = id.Hex()
		sendData.UserID = ""
		return sendData, nil
	}
	return nil, errors.New("credentials was not saved")
}
