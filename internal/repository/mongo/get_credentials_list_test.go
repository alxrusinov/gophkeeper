package mongo

import (
	"context"
	"fmt"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestMongo_GetCredentialsList(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	ctx := context.Background()

	userID := primitive.NewObjectID().Hex()

	result := []model.Credentials{{
		ID:     primitive.NewObjectID().Hex(),
		UserID: userID,
		Title:  "title",
		Data: model.Login{
			Username: "user",
			Password: "password",
		},
		Meta: "meta",
	}}

	tests := []struct {
		name        string
		prepareMock func(mt *mtest.T)
		data        string
		want        []model.Credentials
		wantErr     bool
	}{
		{
			name: "1# success",
			prepareMock: func(mt *mtest.T) {

				first := mtest.CreateCursorResponse(1, fmt.Sprintf("%s.%s", DataBase, CredentialsCollection), mtest.FirstBatch, bson.D{
					{
						Key: "_id", Value: userID,
					},
					{
						Key: "user_id", Value: result[0].UserID,
					},
					{
						Key: "title", Value: result[0].Title,
					},
					{
						Key: "data", Value: result[0].Data,
					},
					{
						Key: "meta", Value: result[0].Meta,
					},
				})

				killCursor := mtest.CreateCursorResponse(0, fmt.Sprintf("%s.%s", DataBase, CredentialsCollection), mtest.NextBatch)

				mt.AddMockResponses(first, killCursor)
			},
			data:    userID,
			want:    result,
			wantErr: false,
		},
		{
			name: "2# error",
			prepareMock: func(mt *mtest.T) {

				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			data:    userID,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			tt.prepareMock(mt)

			mongo := &Mongo{
				client: mt.Client,
			}

			got, err := mongo.GetCredentialsList(ctx, tt.data)

			if !tt.wantErr {
				assert.NotEmpty(t, got)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
