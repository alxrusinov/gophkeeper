package mongo

import (
	"context"
	"testing"

	"github.com/alxrusinov/gophkeeper/internal/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestMongo_DeleteBinary(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	ctx := context.Background()

	data := &model.SourceID{
		ID:     primitive.NewObjectID().Hex(),
		UserID: primitive.NewObjectID().Hex(),
	}

	tests := []struct {
		name        string
		prepareMock func(mt *mtest.T)
		data        *model.SourceID
		want        *model.SourceID
		wantErr     bool
	}{
		{
			name: "1# success",
			prepareMock: func(mt *mtest.T) {

				mt.AddMockResponses(mtest.CreateSuccessResponse())
			},
			data:    data,
			want:    data,
			wantErr: false,
		},
		{
			name: "2# error",
			prepareMock: func(mt *mtest.T) {

				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
			},
			data:    data,
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

			got, err := mongo.DeleteBinary(ctx, tt.data)

			if !tt.wantErr {
				assert.Equal(t, tt.want.UserID, got.UserID)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
