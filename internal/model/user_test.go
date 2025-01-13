package model

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserDocumentFromUser(t *testing.T) {
	sucessID := primitive.NewObjectID()
	errID := "123"

	successSource := &User{
		ID:       sucessID.Hex(),
		Username: "user",
	}

	errSource := &User{
		ID:       errID,
		Username: "user",
	}

	successResult := &UserDocument{
		ID:       sucessID,
		Username: "user",
	}
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		args    args
		want    *UserDocument
		wantErr bool
	}{
		{
			name: "1# success",
			args: args{
				user: *successSource,
			},
			want:    successResult,
			wantErr: false,
		},
		{
			name: "2# error",
			args: args{
				user: *errSource,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserDocumentFromUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserDocumentFromUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserDocumentFromUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserFromUserDocument(t *testing.T) {
	sucessID := primitive.NewObjectID()

	successSource := &UserDocument{
		ID:       sucessID,
		Username: "user",
	}

	successResult := &User{
		ID:       sucessID.Hex(),
		Username: successSource.Username,
	}
	type args struct {
		userDoc UserDocument
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "1# success",
			args: args{
				userDoc: *successSource,
			},
			want: successResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserFromUserDocument(tt.args.userDoc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserFromUserDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}
