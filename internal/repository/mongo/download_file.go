package mongo

import (
	"bytes"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

// DownloadFile - downloads file by ud
func (m *Mongo) DownloadFile(ctx context.Context, fileID string) (*bytes.Buffer, error) {
	bucket, err := gridfs.NewBucket(m.client.Database(DataBase))

	if err != nil {
		return nil, err
	}

	fileBuffer := bytes.NewBuffer(nil)

	filter, err := primitive.ObjectIDFromHex(fileID)

	if err != nil {
		return nil, err
	}

	_, err = bucket.DownloadToStream(filter, fileBuffer)

	if err != nil {
		return nil, err
	}

	return fileBuffer, nil

}
