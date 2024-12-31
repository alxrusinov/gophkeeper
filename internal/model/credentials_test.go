package model

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCredentialsDocumentFromCredentials(t *testing.T) {
	sucessID := primitive.NewObjectID()
	errID := "123"

	successSource := &Credentials{
		ID:     sucessID.Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Data: Login{
			Username: "user",
			Password: "pass",
		},
		Title: "title",
		Meta:  "meta",
	}

	errSource := &Credentials{
		ID:     errID,
		UserID: primitive.NewObjectID().Hex(),
		Data: Login{
			Username: "user",
			Password: "pass",
		},
		Title: "title",
		Meta:  "meta",
	}

	successResult := &CredentialsDocument{
		ID:     sucessID,
		UserID: successSource.UserID,
		Title:  successSource.Title,
		Data:   successSource.Data,
		Meta:   successSource.Meta,
	}
	type args struct {
		cred Credentials
	}
	tests := []struct {
		name    string
		args    args
		want    *CredentialsDocument
		wantErr bool
	}{
		{
			name: "1# success",
			args: args{
				cred: *successSource,
			},
			want:    successResult,
			wantErr: false,
		},
		{
			name: "2# error",
			args: args{
				cred: *errSource,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CredentialsDocumentFromCredentials(tt.args.cred)
			if (err != nil) != tt.wantErr {
				t.Errorf("CredentialsDocumentFromCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CredentialsDocumentFromCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsFromCredentialsDocument(t *testing.T) {
	sucessID := primitive.NewObjectID()

	successSource := &CredentialsDocument{
		ID:     sucessID,
		UserID: primitive.NewObjectID().Hex(),
		Data: Login{
			Username: "user",
			Password: "pass",
		},
		Title: "title",
		Meta:  "meta",
	}

	successResult := &Credentials{
		ID:     sucessID.Hex(),
		UserID: successSource.UserID,
		Data:   successSource.Data,
		Title:  successSource.Title,
		Meta:   successSource.Meta,
	}
	type args struct {
		credDoc CredentialsDocument
	}
	tests := []struct {
		name string
		args args
		want *Credentials
	}{
		{
			name: "1# success",
			args: args{
				credDoc: *successSource,
			},
			want: successResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CredentialsFromCredentialsDocument(tt.args.credDoc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CredentialsFromCredentialsDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}
