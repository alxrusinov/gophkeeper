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

func TestMongo_CheckUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	ctx := context.Background()

	userID := primitive.NewObjectID().Hex()

	user := &model.User{
		Username: "user",
		ID:       userID,
	}

	tests := []struct {
		name        string
		prepareMock func(mt *mtest.T)
		userID      string
		want        bool
		wantErr     bool
	}{
		{
			name: "1# success",
			prepareMock: func(mt *mtest.T) {
				first := mtest.CreateCursorResponse(1, fmt.Sprintf("%s.%s", DataBase, UserCollection), mtest.FirstBatch, bson.D{
					{
						Key: "_id", Value: userID,
					},
					{
						Key: "username", Value: user.Username,
					},
				})

				killCursor := mtest.CreateCursorResponse(0, fmt.Sprintf("%s.%s", DataBase, UserCollection), mtest.NextBatch)

				mt.AddMockResponses(first, killCursor)
			},
			userID:  userID,
			want:    true,
			wantErr: false,
		},
		{
			name: "2# wrong id",
			prepareMock: func(mt *mtest.T) {
				first := mtest.CreateCursorResponse(1, fmt.Sprintf("%s.%s", DataBase, UserCollection), mtest.FirstBatch, bson.D{
					{
						Key: "_id", Value: userID,
					},
					{
						Key: "username", Value: user.Username,
					},
				})

				killCursor := mtest.CreateCursorResponse(0, fmt.Sprintf("%s.%s", DataBase, UserCollection), mtest.NextBatch)

				mt.AddMockResponses(first, killCursor)
			},
			userID:  "123",
			want:    false,
			wantErr: true,
		},
		{
			name: "2# not found",
			prepareMock: func(mt *mtest.T) {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			userID:  userID,
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			tt.prepareMock(mt)

			mongo := &Mongo{
				client: mt.Client,
			}

			got, err := mongo.CheckUser(ctx, tt.userID)

			switch tt.name {
			case tests[0].name:
				assert.True(t, got)
				assert.Nil(t, err)
			case tests[1].name:
				assert.False(t, got)
				assert.NotNil(t, err)
			case tests[2].name:
				assert.False(t, got)
				assert.NotNil(t, err)
			}
		})
	}
}
